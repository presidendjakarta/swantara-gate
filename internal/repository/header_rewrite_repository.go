package repository

import (
	"database/sql"
	"fmt"

	"github.com/presidendjakarta/swantara-gate/internal/model"
)

// =========================================================
// REQUEST HEADER RULE REPOSITORY
// =========================================================

// RequestHeaderRuleRepository menangani operasi database untuk Request Header Rule
type RequestHeaderRuleRepository struct {
	DB *sql.DB
}

// NewRequestHeaderRuleRepository membuat instance baru
func NewRequestHeaderRuleRepository(db *sql.DB) *RequestHeaderRuleRepository {
	return &RequestHeaderRuleRepository{DB: db}
}

// Create membuat request header rule baru
func (r *RequestHeaderRuleRepository) Create(req *model.CreateRequestHeaderRuleRequest) (*model.RequestHeaderRule, error) {
	query := `
		INSERT INTO request_header_rules (virtual_directory_id, header_name, operation, value_source,
			header_value, source_header, variable_name, execution_order, is_active)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
		RETURNING id, created_at
	`
	var item model.RequestHeaderRule
	err := r.DB.QueryRow(query,
		req.VirtualDirectoryID, req.HeaderName, req.Operation, req.ValueSource,
		req.HeaderValue, req.SourceHeader, req.VariableName, req.ExecutionOrder, req.IsActive,
	).Scan(&item.ID, &item.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("gagal membuat request header rule: %w", err)
	}

	item.VirtualDirectoryID = req.VirtualDirectoryID
	item.HeaderName = req.HeaderName
	item.Operation = req.Operation
	item.ValueSource = req.ValueSource
	item.HeaderValue = req.HeaderValue
	item.SourceHeader = req.SourceHeader
	item.VariableName = req.VariableName
	item.ExecutionOrder = req.ExecutionOrder
	item.IsActive = req.IsActive
	return &item, nil
}

// GetByID mengambil request header rule berdasarkan ID
func (r *RequestHeaderRuleRepository) GetByID(id int64) (*model.RequestHeaderRule, error) {
	query := `
		SELECT rh.id, rh.virtual_directory_id, rh.header_name, rh.operation, rh.value_source,
		       rh.header_value, rh.source_header, rh.variable_name, rh.execution_order,
		       rh.is_active, rh.created_at, COALESCE(vd.source_path, '')
		FROM request_header_rules rh
		LEFT JOIN virtual_directories vd ON rh.virtual_directory_id = vd.id
		WHERE rh.id = ?
	`
	var item model.RequestHeaderRule
	err := r.DB.QueryRow(query, id).Scan(
		&item.ID, &item.VirtualDirectoryID, &item.HeaderName, &item.Operation, &item.ValueSource,
		&item.HeaderValue, &item.SourceHeader, &item.VariableName, &item.ExecutionOrder,
		&item.IsActive, &item.CreatedAt, &item.SourcePath,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("request header rule tidak ditemukan")
		}
		return nil, fmt.Errorf("gagal mengambil request header rule: %w", err)
	}
	return &item, nil
}

// GetAll mengambil semua request header rules dengan pagination
func (r *RequestHeaderRuleRepository) GetAll(page, limit int) ([]model.RequestHeaderRule, error) {
	offset := (page - 1) * limit
	query := `
		SELECT rh.id, rh.virtual_directory_id, rh.header_name, rh.operation, rh.value_source,
		       rh.header_value, rh.source_header, rh.variable_name, rh.execution_order,
		       rh.is_active, rh.created_at, COALESCE(vd.source_path, '')
		FROM request_header_rules rh
		LEFT JOIN virtual_directories vd ON rh.virtual_directory_id = vd.id
		ORDER BY rh.execution_order ASC, rh.id DESC
		LIMIT ? OFFSET ?
	`
	rows, err := r.DB.Query(query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("gagal mengambil request header rules: %w", err)
	}
	defer rows.Close()

	var items []model.RequestHeaderRule
	for rows.Next() {
		var item model.RequestHeaderRule
		err := rows.Scan(
			&item.ID, &item.VirtualDirectoryID, &item.HeaderName, &item.Operation, &item.ValueSource,
			&item.HeaderValue, &item.SourceHeader, &item.VariableName, &item.ExecutionOrder,
			&item.IsActive, &item.CreatedAt, &item.SourcePath,
		)
		if err != nil {
			return nil, fmt.Errorf("gagal scan request header rule: %w", err)
		}
		items = append(items, item)
	}
	return items, nil
}

// GetByDirectoryID mengambil rules berdasarkan virtual directory ID
func (r *RequestHeaderRuleRepository) GetByDirectoryID(dirID int64) ([]model.RequestHeaderRule, error) {
	query := `
		SELECT rh.id, rh.virtual_directory_id, rh.header_name, rh.operation, rh.value_source,
		       rh.header_value, rh.source_header, rh.variable_name, rh.execution_order,
		       rh.is_active, rh.created_at, COALESCE(vd.source_path, '')
		FROM request_header_rules rh
		LEFT JOIN virtual_directories vd ON rh.virtual_directory_id = vd.id
		WHERE rh.virtual_directory_id = ?
		ORDER BY rh.execution_order ASC
	`
	rows, err := r.DB.Query(query, dirID)
	if err != nil {
		return nil, fmt.Errorf("gagal mengambil request header rules: %w", err)
	}
	defer rows.Close()

	var items []model.RequestHeaderRule
	for rows.Next() {
		var item model.RequestHeaderRule
		err := rows.Scan(
			&item.ID, &item.VirtualDirectoryID, &item.HeaderName, &item.Operation, &item.ValueSource,
			&item.HeaderValue, &item.SourceHeader, &item.VariableName, &item.ExecutionOrder,
			&item.IsActive, &item.CreatedAt, &item.SourcePath,
		)
		if err != nil {
			return nil, fmt.Errorf("gagal scan request header rule: %w", err)
		}
		items = append(items, item)
	}
	return items, nil
}

// Update memperbarui request header rule
func (r *RequestHeaderRuleRepository) Update(id int64, req *model.UpdateRequestHeaderRuleRequest) error {
	query := `
		UPDATE request_header_rules
		SET header_name = ?, operation = ?, value_source = ?, header_value = ?,
		    source_header = ?, variable_name = ?, execution_order = ?, is_active = ?
		WHERE id = ?
	`
	_, err := r.DB.Exec(query,
		req.HeaderName, req.Operation, req.ValueSource, req.HeaderValue,
		req.SourceHeader, req.VariableName, req.ExecutionOrder, req.IsActive, id,
	)
	if err != nil {
		return fmt.Errorf("gagal update request header rule: %w", err)
	}
	return nil
}

// Delete menghapus request header rule
func (r *RequestHeaderRuleRepository) Delete(id int64) error {
	_, err := r.DB.Exec("DELETE FROM request_header_rules WHERE id = ?", id)
	if err != nil {
		return fmt.Errorf("gagal menghapus request header rule: %w", err)
	}
	return nil
}

// Count menghitung total request header rules
func (r *RequestHeaderRuleRepository) Count() (int64, error) {
	var count int64
	err := r.DB.QueryRow("SELECT COUNT(*) FROM request_header_rules").Scan(&count)
	return count, err
}

// =========================================================
// RESPONSE HEADER RULE REPOSITORY
// =========================================================

// ResponseHeaderRuleRepository menangani operasi database untuk Response Header Rule
type ResponseHeaderRuleRepository struct {
	DB *sql.DB
}

// NewResponseHeaderRuleRepository membuat instance baru
func NewResponseHeaderRuleRepository(db *sql.DB) *ResponseHeaderRuleRepository {
	return &ResponseHeaderRuleRepository{DB: db}
}

// Create membuat response header rule baru
func (r *ResponseHeaderRuleRepository) Create(req *model.CreateResponseHeaderRuleRequest) (*model.ResponseHeaderRule, error) {
	query := `
		INSERT INTO response_header_rules (virtual_directory_id, header_name, operation, header_value, execution_order, is_active)
		VALUES (?, ?, ?, ?, ?, ?)
		RETURNING id, created_at
	`
	var item model.ResponseHeaderRule
	err := r.DB.QueryRow(query,
		req.VirtualDirectoryID, req.HeaderName, req.Operation, req.HeaderValue, req.ExecutionOrder, req.IsActive,
	).Scan(&item.ID, &item.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("gagal membuat response header rule: %w", err)
	}

	item.VirtualDirectoryID = req.VirtualDirectoryID
	item.HeaderName = req.HeaderName
	item.Operation = req.Operation
	item.HeaderValue = req.HeaderValue
	item.ExecutionOrder = req.ExecutionOrder
	item.IsActive = req.IsActive
	return &item, nil
}

// GetByID mengambil response header rule berdasarkan ID
func (r *ResponseHeaderRuleRepository) GetByID(id int64) (*model.ResponseHeaderRule, error) {
	query := `
		SELECT rh.id, rh.virtual_directory_id, rh.header_name, rh.operation,
		       rh.header_value, rh.execution_order, rh.is_active, rh.created_at, COALESCE(vd.source_path, '')
		FROM response_header_rules rh
		LEFT JOIN virtual_directories vd ON rh.virtual_directory_id = vd.id
		WHERE rh.id = ?
	`
	var item model.ResponseHeaderRule
	err := r.DB.QueryRow(query, id).Scan(
		&item.ID, &item.VirtualDirectoryID, &item.HeaderName, &item.Operation,
		&item.HeaderValue, &item.ExecutionOrder, &item.IsActive, &item.CreatedAt, &item.SourcePath,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("response header rule tidak ditemukan")
		}
		return nil, fmt.Errorf("gagal mengambil response header rule: %w", err)
	}
	return &item, nil
}

// GetAll mengambil semua response header rules dengan pagination
func (r *ResponseHeaderRuleRepository) GetAll(page, limit int) ([]model.ResponseHeaderRule, error) {
	offset := (page - 1) * limit
	query := `
		SELECT rh.id, rh.virtual_directory_id, rh.header_name, rh.operation,
		       rh.header_value, rh.execution_order, rh.is_active, rh.created_at, COALESCE(vd.source_path, '')
		FROM response_header_rules rh
		LEFT JOIN virtual_directories vd ON rh.virtual_directory_id = vd.id
		ORDER BY rh.execution_order ASC, rh.id DESC
		LIMIT ? OFFSET ?
	`
	rows, err := r.DB.Query(query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("gagal mengambil response header rules: %w", err)
	}
	defer rows.Close()

	var items []model.ResponseHeaderRule
	for rows.Next() {
		var item model.ResponseHeaderRule
		err := rows.Scan(
			&item.ID, &item.VirtualDirectoryID, &item.HeaderName, &item.Operation,
			&item.HeaderValue, &item.ExecutionOrder, &item.IsActive, &item.CreatedAt, &item.SourcePath,
		)
		if err != nil {
			return nil, fmt.Errorf("gagal scan response header rule: %w", err)
		}
		items = append(items, item)
	}
	return items, nil
}

// GetByDirectoryID mengambil rules berdasarkan virtual directory ID
func (r *ResponseHeaderRuleRepository) GetByDirectoryID(dirID int64) ([]model.ResponseHeaderRule, error) {
	query := `
		SELECT rh.id, rh.virtual_directory_id, rh.header_name, rh.operation,
		       rh.header_value, rh.execution_order, rh.is_active, rh.created_at, COALESCE(vd.source_path, '')
		FROM response_header_rules rh
		LEFT JOIN virtual_directories vd ON rh.virtual_directory_id = vd.id
		WHERE rh.virtual_directory_id = ?
		ORDER BY rh.execution_order ASC
	`
	rows, err := r.DB.Query(query, dirID)
	if err != nil {
		return nil, fmt.Errorf("gagal mengambil response header rules: %w", err)
	}
	defer rows.Close()

	var items []model.ResponseHeaderRule
	for rows.Next() {
		var item model.ResponseHeaderRule
		err := rows.Scan(
			&item.ID, &item.VirtualDirectoryID, &item.HeaderName, &item.Operation,
			&item.HeaderValue, &item.ExecutionOrder, &item.IsActive, &item.CreatedAt, &item.SourcePath,
		)
		if err != nil {
			return nil, fmt.Errorf("gagal scan response header rule: %w", err)
		}
		items = append(items, item)
	}
	return items, nil
}

// Update memperbarui response header rule
func (r *ResponseHeaderRuleRepository) Update(id int64, req *model.UpdateResponseHeaderRuleRequest) error {
	query := `
		UPDATE response_header_rules
		SET header_name = ?, operation = ?, header_value = ?, execution_order = ?, is_active = ?
		WHERE id = ?
	`
	_, err := r.DB.Exec(query, req.HeaderName, req.Operation, req.HeaderValue, req.ExecutionOrder, req.IsActive, id)
	if err != nil {
		return fmt.Errorf("gagal update response header rule: %w", err)
	}
	return nil
}

// Delete menghapus response header rule
func (r *ResponseHeaderRuleRepository) Delete(id int64) error {
	_, err := r.DB.Exec("DELETE FROM response_header_rules WHERE id = ?", id)
	if err != nil {
		return fmt.Errorf("gagal menghapus response header rule: %w", err)
	}
	return nil
}

// Count menghitung total response header rules
func (r *ResponseHeaderRuleRepository) Count() (int64, error) {
	var count int64
	err := r.DB.QueryRow("SELECT COUNT(*) FROM response_header_rules").Scan(&count)
	return count, err
}

// =========================================================
// QUERY REWRITE REPOSITORY
// =========================================================

// QueryRewriteRepository menangani operasi database untuk Query Rewrite
type QueryRewriteRepository struct {
	DB *sql.DB
}

// NewQueryRewriteRepository membuat instance baru
func NewQueryRewriteRepository(db *sql.DB) *QueryRewriteRepository {
	return &QueryRewriteRepository{DB: db}
}

// Create membuat query rewrite baru
func (r *QueryRewriteRepository) Create(req *model.CreateQueryRewriteRequest) (*model.QueryRewrite, error) {
	query := `
		INSERT INTO query_rewrites (virtual_directory_id, param_name, param_value, operation)
		VALUES (?, ?, ?, ?)
		RETURNING id, created_at
	`
	var item model.QueryRewrite
	err := r.DB.QueryRow(query,
		req.VirtualDirectoryID, req.ParamName, req.ParamValue, req.Operation,
	).Scan(&item.ID, &item.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("gagal membuat query rewrite: %w", err)
	}

	item.VirtualDirectoryID = req.VirtualDirectoryID
	item.ParamName = req.ParamName
	item.ParamValue = req.ParamValue
	item.Operation = req.Operation
	return &item, nil
}

// GetByID mengambil query rewrite berdasarkan ID
func (r *QueryRewriteRepository) GetByID(id int64) (*model.QueryRewrite, error) {
	query := `
		SELECT qr.id, qr.virtual_directory_id, qr.param_name, qr.param_value,
		       qr.operation, qr.created_at, COALESCE(vd.source_path, '')
		FROM query_rewrites qr
		LEFT JOIN virtual_directories vd ON qr.virtual_directory_id = vd.id
		WHERE qr.id = ?
	`
	var item model.QueryRewrite
	err := r.DB.QueryRow(query, id).Scan(
		&item.ID, &item.VirtualDirectoryID, &item.ParamName, &item.ParamValue,
		&item.Operation, &item.CreatedAt, &item.SourcePath,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("query rewrite tidak ditemukan")
		}
		return nil, fmt.Errorf("gagal mengambil query rewrite: %w", err)
	}
	return &item, nil
}

// GetAll mengambil semua query rewrites dengan pagination
func (r *QueryRewriteRepository) GetAll(page, limit int) ([]model.QueryRewrite, error) {
	offset := (page - 1) * limit
	query := `
		SELECT qr.id, qr.virtual_directory_id, qr.param_name, qr.param_value,
		       qr.operation, qr.created_at, COALESCE(vd.source_path, '')
		FROM query_rewrites qr
		LEFT JOIN virtual_directories vd ON qr.virtual_directory_id = vd.id
		ORDER BY qr.id DESC
		LIMIT ? OFFSET ?
	`
	rows, err := r.DB.Query(query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("gagal mengambil query rewrites: %w", err)
	}
	defer rows.Close()

	var items []model.QueryRewrite
	for rows.Next() {
		var item model.QueryRewrite
		err := rows.Scan(
			&item.ID, &item.VirtualDirectoryID, &item.ParamName, &item.ParamValue,
			&item.Operation, &item.CreatedAt, &item.SourcePath,
		)
		if err != nil {
			return nil, fmt.Errorf("gagal scan query rewrite: %w", err)
		}
		items = append(items, item)
	}
	return items, nil
}

// GetByDirectoryID mengambil rewrites berdasarkan virtual directory ID
func (r *QueryRewriteRepository) GetByDirectoryID(dirID int64) ([]model.QueryRewrite, error) {
	query := `
		SELECT qr.id, qr.virtual_directory_id, qr.param_name, qr.param_value,
		       qr.operation, qr.created_at, COALESCE(vd.source_path, '')
		FROM query_rewrites qr
		LEFT JOIN virtual_directories vd ON qr.virtual_directory_id = vd.id
		WHERE qr.virtual_directory_id = ?
		ORDER BY qr.id ASC
	`
	rows, err := r.DB.Query(query, dirID)
	if err != nil {
		return nil, fmt.Errorf("gagal mengambil query rewrites: %w", err)
	}
	defer rows.Close()

	var items []model.QueryRewrite
	for rows.Next() {
		var item model.QueryRewrite
		err := rows.Scan(
			&item.ID, &item.VirtualDirectoryID, &item.ParamName, &item.ParamValue,
			&item.Operation, &item.CreatedAt, &item.SourcePath,
		)
		if err != nil {
			return nil, fmt.Errorf("gagal scan query rewrite: %w", err)
		}
		items = append(items, item)
	}
	return items, nil
}

// Update memperbarui query rewrite
func (r *QueryRewriteRepository) Update(id int64, req *model.UpdateQueryRewriteRequest) error {
	query := `
		UPDATE query_rewrites
		SET param_name = ?, param_value = ?, operation = ?
		WHERE id = ?
	`
	_, err := r.DB.Exec(query, req.ParamName, req.ParamValue, req.Operation, id)
	if err != nil {
		return fmt.Errorf("gagal update query rewrite: %w", err)
	}
	return nil
}

// Delete menghapus query rewrite
func (r *QueryRewriteRepository) Delete(id int64) error {
	_, err := r.DB.Exec("DELETE FROM query_rewrites WHERE id = ?", id)
	if err != nil {
		return fmt.Errorf("gagal menghapus query rewrite: %w", err)
	}
	return nil
}

// Count menghitung total query rewrites
func (r *QueryRewriteRepository) Count() (int64, error) {
	var count int64
	err := r.DB.QueryRow("SELECT COUNT(*) FROM query_rewrites").Scan(&count)
	return count, err
}
