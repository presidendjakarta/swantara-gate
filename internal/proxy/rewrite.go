package proxy

import (
	"database/sql"
	"log"
	"net/http"
	"strings"
	"sync"
)

// QueryRewriteRule represents a query parameter rewrite rule
type QueryRewriteRule struct {
	ID         int64
	VDirID     int64
	ParamName  string
	ParamValue string
	Operation  string // set, delete, rename, append
}

// QueryRewriter loads and applies query parameter rewrite rules
type QueryRewriter struct {
	mu    sync.RWMutex
	db    *sql.DB
	rules map[int64][]*QueryRewriteRule // vdir_id -> rules
}

// NewQueryRewriter creates a new QueryRewriter
func NewQueryRewriter(db *sql.DB) *QueryRewriter {
	qr := &QueryRewriter{
		db:    db,
		rules: make(map[int64][]*QueryRewriteRule),
	}
	qr.Reload()
	return qr
}

// Reload loads all query rewrite rules from the database
func (qr *QueryRewriter) Reload() {
	rules := make(map[int64][]*QueryRewriteRule)

	rows, err := qr.db.Query(`
		SELECT id, virtual_directory_id, param_name, COALESCE(param_value, ''), COALESCE(operation, 'set')
		FROM query_rewrites
	`)
	if err != nil {
		log.Printf("[Proxy] Error loading query rewrite rules: %v", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var rule QueryRewriteRule
		if err := rows.Scan(&rule.ID, &rule.VDirID, &rule.ParamName,
			&rule.ParamValue, &rule.Operation); err != nil {
			log.Printf("[Proxy] Error scanning query rewrite rule: %v", err)
			continue
		}
		rules[rule.VDirID] = append(rules[rule.VDirID], &rule)
	}

	qr.mu.Lock()
	qr.rules = rules
	qr.mu.Unlock()

	log.Printf("[Proxy] Query rewrite rules loaded: %d route groups", len(rules))
}

// Apply applies query rewrite rules to the request URL
func (qr *QueryRewriter) Apply(vdirID int64, r *http.Request) {
	qr.mu.RLock()
	rules, ok := qr.rules[vdirID]
	qr.mu.RUnlock()

	if !ok || len(rules) == 0 {
		return
	}

	query := r.URL.Query()

	for _, rule := range rules {
		switch strings.ToLower(rule.Operation) {
		case "set":
			query.Set(rule.ParamName, rule.ParamValue)
		case "delete":
			query.Del(rule.ParamName)
		case "rename":
			// Rename: copy value from old param to new name, delete old
			if val := query.Get(rule.ParamName); val != "" {
				query.Del(rule.ParamName)
				query.Set(rule.ParamValue, val) // ParamValue holds new name
			}
		case "append":
			query.Add(rule.ParamName, rule.ParamValue)
		}
	}

	r.URL.RawQuery = query.Encode()
}
