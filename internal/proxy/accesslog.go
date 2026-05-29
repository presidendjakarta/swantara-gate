package proxy

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	_ "modernc.org/sqlite"
)

// AccessLogEntry represents a single access log record
type AccessLogEntry struct {
	Timestamp    time.Time
	Method       string
	Host         string
	Path         string
	StatusCode   int
	Duration     time.Duration
	ClientIP     string
	UserAgent    string
	BytesSent    int
	VHostID      int64
	VDirID       int64
	UpstreamHost string
	UpstreamPort int
}

// AccessLogger logs proxy requests to rotating SQLite files (one per day)
type AccessLogger struct {
	logDir    string
	enabled   bool
	logCh     chan *AccessLogEntry
	currentDB *sql.DB
	currentDate string
	mu        sync.Mutex
}

// NewAccessLogger creates a new access logger with daily rotation
func NewAccessLogger(db *sql.DB, logDir string) *AccessLogger {
	if logDir == "" {
		logDir = "logs"
	}

	// Create log directory if not exists
	os.MkdirAll(logDir, 0755)

	al := &AccessLogger{
		logDir:  logDir,
		enabled: true,
		logCh:   make(chan *AccessLogEntry, 1000),
	}

	// Initialize today's log file
	al.rotateIfNeeded()

	// Start background writer
	go al.writer()

	// Start daily rotation checker
	go al.rotationChecker()

	return al
}

// rotateIfNeeded creates a new log file if the date has changed
func (al *AccessLogger) rotateIfNeeded() {
	al.mu.Lock()
	defer al.mu.Unlock()

	today := time.Now().Format("2006-01-02")
	if al.currentDate == today && al.currentDB != nil {
		return // Already initialized for today
	}

	// Close yesterday's DB
	if al.currentDB != nil {
		al.currentDB.Close()
	}

	// Open today's log file
	logFile := filepath.Join(al.logDir, fmt.Sprintf("access-%s.db", today))
	db, err := sql.Open("sqlite", logFile)
	if err != nil {
		log.Printf("[Proxy:AccessLog] Error opening log file %s: %v", logFile, err)
		return
	}

	// Enable WAL mode for better concurrent performance
	db.Exec("PRAGMA journal_mode=WAL")
	db.Exec("PRAGMA synchronous=NORMAL")

	// Create table
	db.Exec(`
		CREATE TABLE IF NOT EXISTS access_logs (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			timestamp DATETIME DEFAULT CURRENT_TIMESTAMP,
			method TEXT,
			host TEXT,
			path TEXT,
			status_code INTEGER,
			duration_ms INTEGER,
			client_ip TEXT,
			user_agent TEXT,
			bytes_sent INTEGER,
			vhost_id INTEGER,
			vdir_id INTEGER,
			upstream_host TEXT,
			upstream_port INTEGER
		)
	`)

	al.currentDB = db
	al.currentDate = today

	log.Printf("[Proxy:AccessLog] Rotated to new log file: %s", logFile)

	// Cleanup old logs (keep only last 24 hours)
	go al.cleanupOldLogs()
}

// rotationChecker runs daily at midnight to rotate log file
func (al *AccessLogger) rotationChecker() {
	for {
		now := time.Now()
		// Calculate time until midnight
		tomorrow := now.Add(24 * time.Hour)
		midnight := time.Date(tomorrow.Year(), tomorrow.Month(), tomorrow.Day(), 0, 0, 0, 0, now.Location())
		timeUntilMidnight := midnight.Sub(now)

		time.Sleep(timeUntilMidnight)
		al.rotateIfNeeded()
	}
}

// cleanupOldLogs removes log files older than 24 hours
func (al *AccessLogger) cleanupOldLogs() {
	cutoff := time.Now().Add(-24 * time.Hour)
	cutoffDate := cutoff.Format("2006-01-02")

	entries, err := os.ReadDir(al.logDir)
	if err != nil {
		log.Printf("[Proxy:AccessLog] Error reading log directory: %v", err)
		return
	}

	for _, entry := range entries {
		if !strings.HasSuffix(entry.Name(), ".db") || !strings.HasPrefix(entry.Name(), "access-") {
			continue
		}

		// Extract date from filename (access-2006-01-02.db)
		name := entry.Name()
		dateStr := strings.TrimPrefix(name, "access-")
		dateStr = strings.TrimSuffix(dateStr, ".db")

		if dateStr < cutoffDate {
			filePath := filepath.Join(al.logDir, name)
			if err := os.Remove(filePath); err != nil {
				log.Printf("[Proxy:AccessLog] Error removing old log %s: %v", name, err)
			} else {
				log.Printf("[Proxy:AccessLog] Removed old log file: %s", name)
			}
		}
	}
}

// ensureTable kept for backward compatibility (no-op now)
func (al *AccessLogger) ensureTable() {
	// Table creation moved to rotateIfNeeded
}

// Log records an access log entry asynchronously
func (al *AccessLogger) Log(entry *AccessLogEntry) {
	if !al.enabled {
		return
	}
	select {
	case al.logCh <- entry:
	default:
		// Channel full, drop log entry (non-blocking)
	}
}

// LogRequest is a helper to build and log an entry from request/response context
func (al *AccessLogger) LogRequest(r *http.Request, statusCode int, bytesSent int, duration time.Duration, vhostID, vdirID int64, upstream *UpstreamConfig) {
	entry := &AccessLogEntry{
		Timestamp:  time.Now(),
		Method:     r.Method,
		Host:       r.Host,
		Path:       r.URL.Path,
		StatusCode: statusCode,
		Duration:   duration,
		ClientIP:   extractClientIP(r),
		UserAgent:  r.Header.Get("User-Agent"),
		BytesSent:  bytesSent,
		VHostID:    vhostID,
		VDirID:     vdirID,
	}
	if upstream != nil {
		entry.UpstreamHost = upstream.TargetHost
		entry.UpstreamPort = upstream.TargetPort
	}
	al.Log(entry)

	// Also log to stdout
	log.Printf("[Proxy] %s %s%s -> %d (%s) [%s]",
		entry.Method, entry.Host, entry.Path,
		entry.StatusCode, entry.Duration.Round(time.Millisecond),
		entry.ClientIP)
}

// writer is a background goroutine that writes log entries to the database
func (al *AccessLogger) writer() {
	// Batch insert every 100 entries or 1 second
	batch := make([]*AccessLogEntry, 0, 100)
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case entry, ok := <-al.logCh:
			if !ok {
				// Channel closed, flush remaining
				al.flush(batch)
				return
			}
			batch = append(batch, entry)
			if len(batch) >= 100 {
				al.flush(batch)
				batch = batch[:0]
			}
		case <-ticker.C:
			if len(batch) > 0 {
				al.flush(batch)
				batch = batch[:0]
			}
		}
	}
}

// flush writes a batch of log entries to the database
func (al *AccessLogger) flush(entries []*AccessLogEntry) {
	if len(entries) == 0 {
		return
	}

	// Ensure we have today's DB
	al.rotateIfNeeded()

	al.mu.Lock()
	db := al.currentDB
	al.mu.Unlock()

	if db == nil {
		log.Printf("[Proxy:AccessLog] No database available for writing logs")
		return
	}

	tx, err := db.Begin()
	if err != nil {
		log.Printf("[Proxy:AccessLog] Error starting transaction: %v", err)
		return
	}

	stmt, err := tx.Prepare(`
		INSERT INTO access_logs (timestamp, method, host, path, status_code, duration_ms, client_ip, user_agent, bytes_sent, vhost_id, vdir_id, upstream_host, upstream_port)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`)
	if err != nil {
		log.Printf("[Proxy:AccessLog] Error preparing statement: %v", err)
		tx.Rollback()
		return
	}
	defer stmt.Close()

	for _, e := range entries {
		_, err = stmt.Exec(
			e.Timestamp.Format("2006-01-02 15:04:05"),
			e.Method, e.Host, e.Path, e.StatusCode,
			e.Duration.Milliseconds(),
			e.ClientIP, e.UserAgent, e.BytesSent,
			e.VHostID, e.VDirID,
			e.UpstreamHost, e.UpstreamPort,
		)
		if err != nil {
			log.Printf("[Proxy:AccessLog] Error inserting log: %v", err)
		}
	}

	if err := tx.Commit(); err != nil {
		log.Printf("[Proxy:AccessLog] Error committing: %v", err)
	}
}

// Stop stops the access logger
func (al *AccessLogger) Stop() {
	close(al.logCh)
}

// formatDuration formats duration for logging
func formatDuration(d time.Duration) string {
	if d < time.Millisecond {
		return fmt.Sprintf("%dμs", d.Microseconds())
	}
	return fmt.Sprintf("%dms", d.Milliseconds())
}
