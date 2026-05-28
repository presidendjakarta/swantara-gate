package model

import "time"

// =========================================================
// REQUEST HEADER RULES
// =========================================================

// RequestHeaderRule merepresentasikan aturan manipulasi request header
type RequestHeaderRule struct {
	ID                 int64     `json:"id"`
	VirtualDirectoryID int64     `json:"virtual_directory_id"`
	HeaderName         string    `json:"header_name"`
	Operation          string    `json:"operation"`
	ValueSource        string    `json:"value_source"`
	HeaderValue        string    `json:"header_value"`
	SourceHeader       string    `json:"source_header"`
	VariableName       string    `json:"variable_name"`
	ExecutionOrder     int       `json:"execution_order"`
	IsActive           bool      `json:"is_active"`
	CreatedAt          time.Time `json:"created_at"`

	// Data join
	SourcePath string `json:"source_path,omitempty"`
}

// CreateRequestHeaderRuleRequest request untuk membuat request header rule baru
type CreateRequestHeaderRuleRequest struct {
	VirtualDirectoryID int64  `json:"virtual_directory_id" validate:"required"`
	HeaderName         string `json:"header_name" validate:"required"`
	Operation          string `json:"operation" validate:"required"`
	ValueSource        string `json:"value_source"`
	HeaderValue        string `json:"header_value"`
	SourceHeader       string `json:"source_header"`
	VariableName       string `json:"variable_name"`
	ExecutionOrder     int    `json:"execution_order"`
	IsActive           bool   `json:"is_active"`
}

// UpdateRequestHeaderRuleRequest request untuk update request header rule
type UpdateRequestHeaderRuleRequest struct {
	HeaderName     string `json:"header_name"`
	Operation      string `json:"operation"`
	ValueSource    string `json:"value_source"`
	HeaderValue    string `json:"header_value"`
	SourceHeader   string `json:"source_header"`
	VariableName   string `json:"variable_name"`
	ExecutionOrder int    `json:"execution_order"`
	IsActive       bool   `json:"is_active"`
}

// =========================================================
// RESPONSE HEADER RULES
// =========================================================

// ResponseHeaderRule merepresentasikan aturan manipulasi response header
type ResponseHeaderRule struct {
	ID                 int64     `json:"id"`
	VirtualDirectoryID int64     `json:"virtual_directory_id"`
	HeaderName         string    `json:"header_name"`
	Operation          string    `json:"operation"`
	HeaderValue        string    `json:"header_value"`
	ExecutionOrder     int       `json:"execution_order"`
	IsActive           bool      `json:"is_active"`
	CreatedAt          time.Time `json:"created_at"`

	// Data join
	SourcePath string `json:"source_path,omitempty"`
}

// CreateResponseHeaderRuleRequest request untuk membuat response header rule baru
type CreateResponseHeaderRuleRequest struct {
	VirtualDirectoryID int64  `json:"virtual_directory_id" validate:"required"`
	HeaderName         string `json:"header_name" validate:"required"`
	Operation          string `json:"operation" validate:"required"`
	HeaderValue        string `json:"header_value"`
	ExecutionOrder     int    `json:"execution_order"`
	IsActive           bool   `json:"is_active"`
}

// UpdateResponseHeaderRuleRequest request untuk update response header rule
type UpdateResponseHeaderRuleRequest struct {
	HeaderName     string `json:"header_name"`
	Operation      string `json:"operation"`
	HeaderValue    string `json:"header_value"`
	ExecutionOrder int    `json:"execution_order"`
	IsActive       bool   `json:"is_active"`
}

// =========================================================
// QUERY REWRITES
// =========================================================

// QueryRewrite merepresentasikan aturan rewrite query parameter
type QueryRewrite struct {
	ID                 int64     `json:"id"`
	VirtualDirectoryID int64     `json:"virtual_directory_id"`
	ParamName          string    `json:"param_name"`
	ParamValue         string    `json:"param_value"`
	Operation          string    `json:"operation"`
	CreatedAt          time.Time `json:"created_at"`

	// Data join
	SourcePath string `json:"source_path,omitempty"`
}

// CreateQueryRewriteRequest request untuk membuat query rewrite baru
type CreateQueryRewriteRequest struct {
	VirtualDirectoryID int64  `json:"virtual_directory_id" validate:"required"`
	ParamName          string `json:"param_name" validate:"required"`
	ParamValue         string `json:"param_value"`
	Operation          string `json:"operation"`
}

// UpdateQueryRewriteRequest request untuk update query rewrite
type UpdateQueryRewriteRequest struct {
	ParamName  string `json:"param_name"`
	ParamValue string `json:"param_value"`
	Operation  string `json:"operation"`
}
