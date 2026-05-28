package proxy

import (
	"database/sql"
	"log"
	"net/http"
	"sort"
	"strings"
	"sync"
)

// RequestHeaderRule represents a header manipulation rule for requests
type RequestHeaderRule struct {
	ID             int64
	VDirID         int64
	HeaderName     string
	Operation      string // add, set, delete, copy, rename
	ValueSource    string // static, header, variable
	HeaderValue    string
	SourceHeader   string
	VariableName   string
	ExecutionOrder int
}

// ResponseHeaderRule represents a header manipulation rule for responses
type ResponseHeaderRule struct {
	ID             int64
	VDirID         int64
	HeaderName     string
	Operation      string // add, set, delete
	HeaderValue    string
	ExecutionOrder int
}

// HeaderProcessor loads and applies header manipulation rules
type HeaderProcessor struct {
	mu              sync.RWMutex
	db              *sql.DB
	requestRules    map[int64][]*RequestHeaderRule  // vdir_id -> rules
	responseRules   map[int64][]*ResponseHeaderRule // vdir_id -> rules
}

// NewHeaderProcessor creates a new HeaderProcessor
func NewHeaderProcessor(db *sql.DB) *HeaderProcessor {
	hp := &HeaderProcessor{
		db:            db,
		requestRules:  make(map[int64][]*RequestHeaderRule),
		responseRules: make(map[int64][]*ResponseHeaderRule),
	}
	hp.Reload()
	return hp
}

// Reload loads all header rules from the database
func (hp *HeaderProcessor) Reload() {
	reqRules := make(map[int64][]*RequestHeaderRule)
	resRules := make(map[int64][]*ResponseHeaderRule)

	// Load request header rules
	rows, err := hp.db.Query(`
		SELECT id, virtual_directory_id, header_name, operation, 
		       COALESCE(value_source, 'static'), COALESCE(header_value, ''),
		       COALESCE(source_header, ''), COALESCE(variable_name, ''), execution_order
		FROM request_header_rules
		WHERE is_active = 1
		ORDER BY execution_order ASC
	`)
	if err != nil {
		log.Printf("[Proxy] Error loading request header rules: %v", err)
	} else {
		defer rows.Close()
		for rows.Next() {
			var rule RequestHeaderRule
			if err := rows.Scan(&rule.ID, &rule.VDirID, &rule.HeaderName,
				&rule.Operation, &rule.ValueSource, &rule.HeaderValue,
				&rule.SourceHeader, &rule.VariableName, &rule.ExecutionOrder); err != nil {
				log.Printf("[Proxy] Error scanning request header rule: %v", err)
				continue
			}
			reqRules[rule.VDirID] = append(reqRules[rule.VDirID], &rule)
		}
	}

	// Load response header rules
	rows2, err := hp.db.Query(`
		SELECT id, virtual_directory_id, header_name, operation,
		       COALESCE(header_value, ''), execution_order
		FROM response_header_rules
		WHERE is_active = 1
		ORDER BY execution_order ASC
	`)
	if err != nil {
		log.Printf("[Proxy] Error loading response header rules: %v", err)
	} else {
		defer rows2.Close()
		for rows2.Next() {
			var rule ResponseHeaderRule
			if err := rows2.Scan(&rule.ID, &rule.VDirID, &rule.HeaderName,
				&rule.Operation, &rule.HeaderValue, &rule.ExecutionOrder); err != nil {
				log.Printf("[Proxy] Error scanning response header rule: %v", err)
				continue
			}
			resRules[rule.VDirID] = append(resRules[rule.VDirID], &rule)
		}
	}

	// Sort rules by execution order
	for _, rules := range reqRules {
		sort.Slice(rules, func(i, j int) bool {
			return rules[i].ExecutionOrder < rules[j].ExecutionOrder
		})
	}
	for _, rules := range resRules {
		sort.Slice(rules, func(i, j int) bool {
			return rules[i].ExecutionOrder < rules[j].ExecutionOrder
		})
	}

	hp.mu.Lock()
	hp.requestRules = reqRules
	hp.responseRules = resRules
	hp.mu.Unlock()

	log.Printf("[Proxy] Header rules loaded: %d request rule groups, %d response rule groups",
		len(reqRules), len(resRules))
}

// ApplyRequestHeaders applies request header rules to the outgoing request
func (hp *HeaderProcessor) ApplyRequestHeaders(vdirID int64, req *http.Request) {
	hp.mu.RLock()
	rules, ok := hp.requestRules[vdirID]
	hp.mu.RUnlock()

	if !ok || len(rules) == 0 {
		return
	}

	for _, rule := range rules {
		switch strings.ToLower(rule.Operation) {
		case "add":
			value := hp.resolveValue(rule, req.Header)
			req.Header.Add(rule.HeaderName, value)
		case "set":
			value := hp.resolveValue(rule, req.Header)
			req.Header.Set(rule.HeaderName, value)
		case "delete":
			req.Header.Del(rule.HeaderName)
		case "copy":
			if rule.SourceHeader != "" {
				if val := req.Header.Get(rule.SourceHeader); val != "" {
					req.Header.Set(rule.HeaderName, val)
				}
			}
		case "rename":
			if rule.SourceHeader != "" {
				if val := req.Header.Get(rule.SourceHeader); val != "" {
					req.Header.Del(rule.SourceHeader)
					req.Header.Set(rule.HeaderName, val)
				}
			}
		}
	}
}

// ApplyResponseHeaders applies response header rules to the proxy response
func (hp *HeaderProcessor) ApplyResponseHeaders(vdirID int64, header http.Header) {
	hp.mu.RLock()
	rules, ok := hp.responseRules[vdirID]
	hp.mu.RUnlock()

	if !ok || len(rules) == 0 {
		return
	}

	for _, rule := range rules {
		switch strings.ToLower(rule.Operation) {
		case "add":
			header.Add(rule.HeaderName, rule.HeaderValue)
		case "set":
			header.Set(rule.HeaderName, rule.HeaderValue)
		case "delete":
			header.Del(rule.HeaderName)
		}
	}
}

// resolveValue resolves the value for a request header rule based on value_source
func (hp *HeaderProcessor) resolveValue(rule *RequestHeaderRule, headers http.Header) string {
	switch strings.ToLower(rule.ValueSource) {
	case "header":
		if rule.SourceHeader != "" {
			return headers.Get(rule.SourceHeader)
		}
		return rule.HeaderValue
	case "variable":
		// Supported variables
		switch rule.VariableName {
		case "request_id":
			return generateRequestID()
		case "timestamp":
			return currentTimestamp()
		default:
			return rule.HeaderValue
		}
	default: // "static"
		return rule.HeaderValue
	}
}

// generateRequestID generates a simple request ID
func generateRequestID() string {
	return strings.Replace(currentTimestamp(), " ", "-", -1)
}

// currentTimestamp returns current time as string
func currentTimestamp() string {
	return strings.Replace(
		strings.Replace(
			strings.Replace(
				strings.Replace(
					http.TimeFormat, "Mon, ", "", 1),
				" GMT", "", 1),
			":", "", -1),
		" ", "", -1)
}
