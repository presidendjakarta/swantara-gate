package proxy

import (
	"database/sql"
	"log"
	"net/http"
	"sync"
	"time"
)

// MaintenanceConfig holds maintenance window settings
type MaintenanceConfig struct {
	VirtualHostID *int64
	Title         string
	StartAt       time.Time
	EndAt         time.Time
	ResponseCode  int
	Message       string
}

// MaintenanceChecker checks if routes are in maintenance mode
type MaintenanceChecker struct {
	mu      sync.RWMutex
	db      *sql.DB
	windows []*MaintenanceConfig
}

// NewMaintenanceChecker creates a new maintenance checker
func NewMaintenanceChecker(db *sql.DB) *MaintenanceChecker {
	m := &MaintenanceChecker{
		db: db,
	}
	m.Reload()
	return m
}

// Reload loads active maintenance windows from database
func (m *MaintenanceChecker) Reload() {
	var windows []*MaintenanceConfig

	rows, err := m.db.Query(`
		SELECT COALESCE(virtual_host_id, 0), COALESCE(title, ''), start_at, end_at,
		       maintenance_response_code, COALESCE(maintenance_message, 'Service under maintenance')
		FROM maintenance_windows WHERE is_active = 1
	`)
	if err != nil {
		log.Printf("[Proxy:Maintenance] Error loading windows: %v", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var cfg MaintenanceConfig
		var vhostID int64
		var startStr, endStr string
		if err := rows.Scan(&vhostID, &cfg.Title, &startStr, &endStr,
			&cfg.ResponseCode, &cfg.Message); err != nil {
			log.Printf("[Proxy:Maintenance] Error scanning window: %v", err)
			continue
		}
		if vhostID > 0 {
			cfg.VirtualHostID = &vhostID
		}

		// Parse times - try multiple formats
		cfg.StartAt = parseTime(startStr)
		cfg.EndAt = parseTime(endStr)

		if cfg.ResponseCode == 0 {
			cfg.ResponseCode = http.StatusServiceUnavailable
		}

		windows = append(windows, &cfg)
	}

	m.mu.Lock()
	m.windows = windows
	m.mu.Unlock()

	log.Printf("[Proxy:Maintenance] Loaded %d maintenance windows", len(windows))
}

// IsInMaintenance checks if a virtual host is currently in a maintenance window
func (m *MaintenanceChecker) IsInMaintenance(vhostID int64) (bool, *MaintenanceConfig) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	now := time.Now().UTC()
	for _, w := range m.windows {
		// Check if window applies to this vhost (nil means global)
		if w.VirtualHostID != nil && *w.VirtualHostID != vhostID {
			continue
		}

		// Check if current time is within the maintenance window
		if now.After(w.StartAt) && now.Before(w.EndAt) {
			return true, w
		}
	}

	return false, nil
}

// parseTime attempts to parse a time string in multiple formats
func parseTime(s string) time.Time {
	formats := []string{
		"2006-01-02 15:04:05",
		"2006-01-02T15:04:05",
		"2006-01-02T15:04:05Z",
		"2006-01-02 15:04:05+00:00",
		"2006-01-02 15:04:05-07:00",
		time.RFC3339,
	}
	for _, f := range formats {
		if t, err := time.Parse(f, s); err == nil {
			return t
		}
	}
	log.Printf("[Proxy:Maintenance] Failed to parse time: '%s'", s)
	return time.Time{}
}
