package service

import (
	"fmt"

	"github.com/presidendjakarta/swantara-gate/internal/model"
	"github.com/presidendjakarta/swantara-gate/internal/repository"
)

// Daftar operasi yang valid untuk header rules
var validHeaderOperations = map[string]bool{
	"set":    true,
	"add":    true,
	"remove": true,
	"rename": true,
	"copy":   true,
}

// Daftar value source yang valid untuk request header
var validValueSources = map[string]bool{
	"static":   true,
	"header":   true,
	"variable": true,
}

// Daftar operasi yang valid untuk query rewrite
var validQueryOperations = map[string]bool{
	"set":    true,
	"add":    true,
	"remove": true,
	"rename": true,
}

// =========================================================
// REQUEST HEADER RULE SERVICE
// =========================================================

// RequestHeaderRuleService menangani business logic untuk Request Header Rule
type RequestHeaderRuleService struct {
	Repo *repository.RequestHeaderRuleRepository
}

// NewRequestHeaderRuleService membuat instance baru
func NewRequestHeaderRuleService(repo *repository.RequestHeaderRuleRepository) *RequestHeaderRuleService {
	return &RequestHeaderRuleService{Repo: repo}
}

// Create membuat request header rule baru dengan validasi
func (s *RequestHeaderRuleService) Create(req *model.CreateRequestHeaderRuleRequest) (*model.RequestHeaderRule, error) {
	if req.VirtualDirectoryID <= 0 {
		return nil, fmt.Errorf("virtual_directory_id wajib diisi")
	}
	if req.HeaderName == "" {
		return nil, fmt.Errorf("header_name wajib diisi")
	}
	if req.Operation == "" {
		return nil, fmt.Errorf("operation wajib diisi")
	}
	if !validHeaderOperations[req.Operation] {
		return nil, fmt.Errorf("operation tidak valid, gunakan: set, add, remove, rename, copy")
	}
	if req.ValueSource == "" {
		req.ValueSource = "static"
	}
	if !validValueSources[req.ValueSource] {
		return nil, fmt.Errorf("value_source tidak valid, gunakan: static, header, variable")
	}
	if req.ExecutionOrder <= 0 {
		req.ExecutionOrder = 1
	}
	return s.Repo.Create(req)
}

// GetByID mengambil request header rule berdasarkan ID
func (s *RequestHeaderRuleService) GetByID(id int64) (*model.RequestHeaderRule, error) {
	return s.Repo.GetByID(id)
}

// GetAll mengambil semua request header rules dengan pagination
func (s *RequestHeaderRuleService) GetAll(page, limit int) ([]model.RequestHeaderRule, int64, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}
	total, err := s.Repo.Count()
	if err != nil {
		return nil, 0, err
	}
	items, err := s.Repo.GetAll(page, limit)
	if err != nil {
		return nil, 0, err
	}
	return items, total, nil
}

// GetByDirectoryID mengambil rules berdasarkan directory
func (s *RequestHeaderRuleService) GetByDirectoryID(dirID int64) ([]model.RequestHeaderRule, error) {
	return s.Repo.GetByDirectoryID(dirID)
}

// Update memperbarui request header rule
func (s *RequestHeaderRuleService) Update(id int64, req *model.UpdateRequestHeaderRuleRequest) error {
	_, err := s.Repo.GetByID(id)
	if err != nil {
		return fmt.Errorf("request header rule tidak ditemukan")
	}
	if req.Operation != "" && !validHeaderOperations[req.Operation] {
		return fmt.Errorf("operation tidak valid, gunakan: set, add, remove, rename, copy")
	}
	if req.ValueSource != "" && !validValueSources[req.ValueSource] {
		return fmt.Errorf("value_source tidak valid, gunakan: static, header, variable")
	}
	return s.Repo.Update(id, req)
}

// Delete menghapus request header rule
func (s *RequestHeaderRuleService) Delete(id int64) error {
	_, err := s.Repo.GetByID(id)
	if err != nil {
		return fmt.Errorf("request header rule tidak ditemukan")
	}
	return s.Repo.Delete(id)
}

// =========================================================
// RESPONSE HEADER RULE SERVICE
// =========================================================

// ResponseHeaderRuleService menangani business logic untuk Response Header Rule
type ResponseHeaderRuleService struct {
	Repo *repository.ResponseHeaderRuleRepository
}

// NewResponseHeaderRuleService membuat instance baru
func NewResponseHeaderRuleService(repo *repository.ResponseHeaderRuleRepository) *ResponseHeaderRuleService {
	return &ResponseHeaderRuleService{Repo: repo}
}

// Create membuat response header rule baru dengan validasi
func (s *ResponseHeaderRuleService) Create(req *model.CreateResponseHeaderRuleRequest) (*model.ResponseHeaderRule, error) {
	if req.VirtualDirectoryID <= 0 {
		return nil, fmt.Errorf("virtual_directory_id wajib diisi")
	}
	if req.HeaderName == "" {
		return nil, fmt.Errorf("header_name wajib diisi")
	}
	if req.Operation == "" {
		return nil, fmt.Errorf("operation wajib diisi")
	}
	if !validHeaderOperations[req.Operation] {
		return nil, fmt.Errorf("operation tidak valid, gunakan: set, add, remove, rename, copy")
	}
	if req.ExecutionOrder <= 0 {
		req.ExecutionOrder = 1
	}
	return s.Repo.Create(req)
}

// GetByID mengambil response header rule berdasarkan ID
func (s *ResponseHeaderRuleService) GetByID(id int64) (*model.ResponseHeaderRule, error) {
	return s.Repo.GetByID(id)
}

// GetAll mengambil semua response header rules dengan pagination
func (s *ResponseHeaderRuleService) GetAll(page, limit int) ([]model.ResponseHeaderRule, int64, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}
	total, err := s.Repo.Count()
	if err != nil {
		return nil, 0, err
	}
	items, err := s.Repo.GetAll(page, limit)
	if err != nil {
		return nil, 0, err
	}
	return items, total, nil
}

// GetByDirectoryID mengambil rules berdasarkan directory
func (s *ResponseHeaderRuleService) GetByDirectoryID(dirID int64) ([]model.ResponseHeaderRule, error) {
	return s.Repo.GetByDirectoryID(dirID)
}

// Update memperbarui response header rule
func (s *ResponseHeaderRuleService) Update(id int64, req *model.UpdateResponseHeaderRuleRequest) error {
	_, err := s.Repo.GetByID(id)
	if err != nil {
		return fmt.Errorf("response header rule tidak ditemukan")
	}
	if req.Operation != "" && !validHeaderOperations[req.Operation] {
		return fmt.Errorf("operation tidak valid, gunakan: set, add, remove, rename, copy")
	}
	return s.Repo.Update(id, req)
}

// Delete menghapus response header rule
func (s *ResponseHeaderRuleService) Delete(id int64) error {
	_, err := s.Repo.GetByID(id)
	if err != nil {
		return fmt.Errorf("response header rule tidak ditemukan")
	}
	return s.Repo.Delete(id)
}

// =========================================================
// QUERY REWRITE SERVICE
// =========================================================

// QueryRewriteService menangani business logic untuk Query Rewrite
type QueryRewriteService struct {
	Repo *repository.QueryRewriteRepository
}

// NewQueryRewriteService membuat instance baru
func NewQueryRewriteService(repo *repository.QueryRewriteRepository) *QueryRewriteService {
	return &QueryRewriteService{Repo: repo}
}

// Create membuat query rewrite baru dengan validasi
func (s *QueryRewriteService) Create(req *model.CreateQueryRewriteRequest) (*model.QueryRewrite, error) {
	if req.VirtualDirectoryID <= 0 {
		return nil, fmt.Errorf("virtual_directory_id wajib diisi")
	}
	if req.ParamName == "" {
		return nil, fmt.Errorf("param_name wajib diisi")
	}
	if req.Operation == "" {
		req.Operation = "set"
	}
	if !validQueryOperations[req.Operation] {
		return nil, fmt.Errorf("operation tidak valid, gunakan: set, add, remove, rename")
	}
	return s.Repo.Create(req)
}

// GetByID mengambil query rewrite berdasarkan ID
func (s *QueryRewriteService) GetByID(id int64) (*model.QueryRewrite, error) {
	return s.Repo.GetByID(id)
}

// GetAll mengambil semua query rewrites dengan pagination
func (s *QueryRewriteService) GetAll(page, limit int) ([]model.QueryRewrite, int64, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}
	total, err := s.Repo.Count()
	if err != nil {
		return nil, 0, err
	}
	items, err := s.Repo.GetAll(page, limit)
	if err != nil {
		return nil, 0, err
	}
	return items, total, nil
}

// GetByDirectoryID mengambil rewrites berdasarkan directory
func (s *QueryRewriteService) GetByDirectoryID(dirID int64) ([]model.QueryRewrite, error) {
	return s.Repo.GetByDirectoryID(dirID)
}

// Update memperbarui query rewrite
func (s *QueryRewriteService) Update(id int64, req *model.UpdateQueryRewriteRequest) error {
	_, err := s.Repo.GetByID(id)
	if err != nil {
		return fmt.Errorf("query rewrite tidak ditemukan")
	}
	if req.Operation != "" && !validQueryOperations[req.Operation] {
		return fmt.Errorf("operation tidak valid, gunakan: set, add, remove, rename")
	}
	return s.Repo.Update(id, req)
}

// Delete menghapus query rewrite
func (s *QueryRewriteService) Delete(id int64) error {
	_, err := s.Repo.GetByID(id)
	if err != nil {
		return fmt.Errorf("query rewrite tidak ditemukan")
	}
	return s.Repo.Delete(id)
}
