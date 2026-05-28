package proxy

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"
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

// AccessLogger logs proxy requests to the database
type AccessLogger struct {
	db      *sql.DB
	enabled bool
	logCh   chan *AccessLogEntry
}

// NewAccessLogger creates a new access logger
func NewAccessLogger(db *sql.DB) *AccessLogger {
	al := &AccessLogger{
		db:      db,
		enabled: true,
		logCh:   make(chan *AccessLogEntry, 1000),
	}

	// Ensure the access_logs table exists
	al.ensureTable()

	// Start background writer
	go al.writer()

	return al
}

// ensureTable creates the access_logs table if it doesn't exist
func (al *AccessLogger) ensureTable() {
	_, err := al.db.Exec(`
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
	if err != nil {
		log.Printf("[Proxy:AccessLog] Error creating table: %v", err)
	}
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

	tx, err := al.db.Begin()
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
