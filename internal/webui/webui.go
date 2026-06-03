package webui

import (
	"embed"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

//go:embed static/*
var webFiles embed.FS

// AdminUI serves the admin panel web interface
type AdminUI struct {
	webRoot     string
	devMode     bool
	adminAPIURL string
}

// NewAdminUI creates a new admin UI handler
func NewAdminUI(webRoot string, devMode bool, adminAPIURL string) *AdminUI {
	return &AdminUI{
		webRoot:     webRoot,
		devMode:     devMode,
		adminAPIURL: adminAPIURL,
	}
}

// Handler returns an http.Handler for serving the admin panel
func (a *AdminUI) Handler() http.Handler {
	if a.devMode {
		// Development mode: serve from disk
		return a.serveFromDisk()
	}
	// Production mode: serve from embedded files
	return a.serveEmbedded()
}

// serveFromDisk serves static files from the filesystem (development)
func (a *AdminUI) serveFromDisk() http.Handler {
	webDir := filepath.Join(a.webRoot, "static")
	
	// Create directory if not exists
	os.MkdirAll(webDir, 0755)
	
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if file exists
		filePath := filepath.Join(webDir, r.URL.Path)
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			// Serve index.html for client-side routing
			http.ServeFile(w, r, filepath.Join(webDir, "index.html"))
			return
		}
		
		// Serve static file
		http.FileServer(http.Dir(webDir)).ServeHTTP(w, r)
	})
}

// serveEmbedded serves static files from embedded filesystem (production)
func (a *AdminUI) serveEmbedded() http.Handler {
	// Strip the "static/" prefix for serving
	subFS, err := fs.Sub(webFiles, "static")
	if err != nil {
		panic("failed to access embedded web files: " + err.Error())
	}
	
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Try to serve the file
		filePath := strings.TrimPrefix(r.URL.Path, "/")
		if filePath == "" {
			filePath = "index.html"
		}
		
		// Check if file exists in embedded FS
		if _, err := fs.Stat(subFS, filePath); err != nil {
			// Serve index.html for client-side routing
			filePath = "index.html"
		}
		
		// Serve the file
		http.ServeFileFS(w, r, subFS, filePath)
	})
}
