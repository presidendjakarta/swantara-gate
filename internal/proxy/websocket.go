package proxy

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"strings"
	"time"
)

// WebSocketProxy handles WebSocket upgrade and bidirectional proxying
type WebSocketProxy struct{}

// NewWebSocketProxy creates a new WebSocketProxy
func NewWebSocketProxy() *WebSocketProxy {
	return &WebSocketProxy{}
}

// IsWebSocketUpgrade checks if the request is a WebSocket upgrade
func (ws *WebSocketProxy) IsWebSocketUpgrade(r *http.Request) bool {
	return strings.EqualFold(r.Header.Get("Upgrade"), "websocket") &&
		strings.Contains(strings.ToLower(r.Header.Get("Connection")), "upgrade")
}

// Proxy handles the WebSocket connection between client and upstream
func (ws *WebSocketProxy) Proxy(w http.ResponseWriter, r *http.Request, upstream *UpstreamConfig, route *VDirConfig) error {
	// Build target address
	targetAddr := fmt.Sprintf("%s:%d", upstream.TargetHost, upstream.TargetPort)

	// Determine the upstream path
	path := r.URL.Path
	if route.StripPrefix && route.MatchType == "prefix" {
		path = strings.TrimPrefix(path, route.SourcePath)
		if path == "" || path[0] != '/' {
			path = "/" + path
		}
	}
	targetPath := strings.TrimSuffix(route.TargetPath, "/")
	if targetPath != "" && targetPath != "/" {
		path = targetPath + path
	}

	// Dial upstream
	timeout := time.Duration(route.ProxyTimeoutSeconds) * time.Second
	if timeout <= 0 {
		timeout = 30 * time.Second
	}

	upstreamConn, err := net.DialTimeout("tcp", targetAddr, timeout)
	if err != nil {
		return fmt.Errorf("failed to connect to upstream WebSocket: %w", err)
	}
	defer upstreamConn.Close()

	// Build the upgrade request to send to upstream
	reqURL := path
	if r.URL.RawQuery != "" {
		reqURL += "?" + r.URL.RawQuery
	}

	// Write the HTTP upgrade request to upstream
	upgradeReq := fmt.Sprintf("%s %s HTTP/1.1\r\n", r.Method, reqURL)
	upgradeReq += fmt.Sprintf("Host: %s\r\n", targetAddr)
	for key, vals := range r.Header {
		for _, val := range vals {
			upgradeReq += fmt.Sprintf("%s: %s\r\n", key, val)
		}
	}
	upgradeReq += "\r\n"

	_, err = upstreamConn.Write([]byte(upgradeReq))
	if err != nil {
		return fmt.Errorf("failed to send upgrade request: %w", err)
	}

	// Hijack the client connection
	hijacker, ok := w.(http.Hijacker)
	if !ok {
		return fmt.Errorf("response writer does not support hijacking")
	}

	clientConn, _, err := hijacker.Hijack()
	if err != nil {
		return fmt.Errorf("failed to hijack client connection: %w", err)
	}
	defer clientConn.Close()

	// Bidirectional copy
	errCh := make(chan error, 2)

	// Upstream -> Client
	go func() {
		_, err := io.Copy(clientConn, upstreamConn)
		errCh <- err
	}()

	// Client -> Upstream
	go func() {
		_, err := io.Copy(upstreamConn, clientConn)
		errCh <- err
	}()

	// Wait for either direction to finish
	err = <-errCh
	if err != nil {
		log.Printf("[WebSocket] Connection closed: %v", err)
	}

	return nil
}
