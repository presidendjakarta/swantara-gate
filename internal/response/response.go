package response

import (
	"encoding/json"
	"net/http"
)

// Response structure untuk standardisasi response API
type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// JSON mengirim response JSON dengan status code yang sesuai
func JSON(w http.ResponseWriter, statusCode int, response Response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}

// Success mengirim response sukses dengan data
func Success(w http.ResponseWriter, message string, data interface{}) {
	JSON(w, http.StatusOK, Response{
		Success: true,
		Message: message,
		Data:    data,
	})
}

// Created mengirim response berhasil dibuat dengan data
func Created(w http.ResponseWriter, message string, data interface{}) {
	JSON(w, http.StatusCreated, Response{
		Success: true,
		Message: message,
		Data:    data,
	})
}

// Error mengirim response error
func Error(w http.ResponseWriter, statusCode int, message string) {
	JSON(w, statusCode, Response{
		Success: false,
		Message: message,
		Error:   message,
	})
}

// BadRequest mengirim response error 400 untuk request yang tidak valid
func BadRequest(w http.ResponseWriter, message string) {
	Error(w, http.StatusBadRequest, message)
}

// NotFound mengirim response error 404 untuk resource yang tidak ditemukan
func NotFound(w http.ResponseWriter, message string) {
	Error(w, http.StatusNotFound, message)
}

// InternalServerError mengirim response error 500 untuk error server
func InternalServerError(w http.ResponseWriter, message string) {
	Error(w, http.StatusInternalServerError, message)
}

// Conflict mengirim response error 409 untuk data yang sudah ada
func Conflict(w http.ResponseWriter, message string) {
	Error(w, http.StatusConflict, message)
}
