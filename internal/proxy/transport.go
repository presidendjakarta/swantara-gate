package proxy

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

// ProxyResponse holds the response from an upstream
type ProxyResponse struct {
	StatusCode int
	Header     http.Header
	Body       []byte
}

// Transport handles forwarding requests to upstream servers with retry logic
type Transport struct {
	client *http.Client
}

// NewTransport creates a new Transport with default settings
func NewTransport() *Transport {
	return &Transport{
		client: &http.Client{
			// No default timeout; we set it per-request based on route config
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			},
		},
	}
}

// Forward sends the request to the selected upstream server
func (t *Transport) Forward(r *http.Request, upstream *UpstreamConfig, route *VDirConfig) (*ProxyResponse, error) {
	var lastErr error

	maxAttempts := route.RetryCount + 1
	if maxAttempts < 1 {
		maxAttempts = 1
	}

	for attempt := 0; attempt < maxAttempts; attempt++ {
		if attempt > 0 {
			// Delay between retries
			delay := time.Duration(route.RetryDelayMs) * time.Millisecond
			if delay <= 0 {
				delay = 100 * time.Millisecond
			}
			time.Sleep(delay)
		}

		resp, err := t.doForward(r, upstream, route)
		if err != nil {
			lastErr = err
			continue
		}

		// If server error and we have retries left, retry
		if resp.StatusCode >= 502 && resp.StatusCode <= 504 && attempt < maxAttempts-1 {
			lastErr = fmt.Errorf("upstream returned %d", resp.StatusCode)
			continue
		}

		return resp, nil
	}

	return nil, fmt.Errorf("all %d attempts failed: %w", maxAttempts, lastErr)
}

// doForward performs a single proxy request to the upstream
func (t *Transport) doForward(r *http.Request, upstream *UpstreamConfig, route *VDirConfig) (*ProxyResponse, error) {
	// Build target URL
	targetURL := buildTargetURL(r, upstream, route)

	// Read the request body (so we can retry)
	var bodyBytes []byte
	if r.Body != nil {
		var err error
		bodyBytes, err = io.ReadAll(r.Body)
		if err != nil {
			return nil, fmt.Errorf("failed to read request body: %w", err)
		}
		r.Body.Close()
		// Restore body for potential retries
		r.Body = io.NopCloser(bytes.NewReader(bodyBytes))
	}

	// Create the upstream request
	timeout := time.Duration(route.ProxyTimeoutSeconds) * time.Second
	if timeout <= 0 {
		timeout = 30 * time.Second
	}
	ctx, cancel := context.WithTimeout(r.Context(), timeout)
	defer cancel()

	var body io.Reader
	if len(bodyBytes) > 0 {
		body = bytes.NewReader(bodyBytes)
	}

	upReq, err := http.NewRequestWithContext(ctx, r.Method, targetURL, body)
	if err != nil {
		return nil, fmt.Errorf("failed to create upstream request: %w", err)
	}

	// Copy headers from original request
	copyHeaders(upReq.Header, r.Header)

	// Set or preserve Host header
	if route.PreserveHostHeader {
		upReq.Host = r.Host
	} else {
		upReq.Host = fmt.Sprintf("%s:%d", upstream.TargetHost, upstream.TargetPort)
	}

	// Add proxy headers
	upReq.Header.Set("X-Forwarded-For", extractClientIP(r))
	upReq.Header.Set("X-Forwarded-Host", r.Host)
	upReq.Header.Set("X-Forwarded-Proto", getScheme(r))
	upReq.Header.Set("X-Real-IP", extractClientIP(r))

	// Remove hop-by-hop headers
	removeHopByHopHeaders(upReq.Header)

	// Execute request
	resp, err := t.client.Do(upReq)
	if err != nil {
		return nil, fmt.Errorf("upstream request failed: %w", err)
	}
	defer resp.Body.Close()

	// Read response body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read upstream response: %w", err)
	}

	// Copy response headers
	respHeader := make(http.Header)
	copyHeaders(respHeader, resp.Header)
	removeHopByHopHeaders(respHeader)

	return &ProxyResponse{
		StatusCode: resp.StatusCode,
		Header:     respHeader,
		Body:       respBody,
	}, nil
}

// buildTargetURL constructs the full URL to the upstream server
func buildTargetURL(r *http.Request, upstream *UpstreamConfig, route *VDirConfig) string {
	scheme := upstream.Protocol
	if scheme == "" {
		scheme = "http"
	}

	// Determine the path
	path := r.URL.Path
	if route.StripPrefix && route.MatchType == "prefix" {
		path = strings.TrimPrefix(path, route.SourcePath)
		if path == "" || path[0] != '/' {
			path = "/" + path
		}
	}

	// Prepend target path
	targetPath := strings.TrimSuffix(route.TargetPath, "/")
	if targetPath != "" && targetPath != "/" {
		path = targetPath + path
	}

	// Build full URL
	host := fmt.Sprintf("%s:%d", upstream.TargetHost, upstream.TargetPort)
	targetURL := fmt.Sprintf("%s://%s%s", scheme, host, path)

	// Append query string
	if r.URL.RawQuery != "" {
		targetURL += "?" + r.URL.RawQuery
	}

	return targetURL
}

// copyHeaders copies all headers from src to dst
func copyHeaders(dst, src http.Header) {
	for k, vv := range src {
		for _, v := range vv {
			dst.Add(k, v)
		}
	}
}

// removeHopByHopHeaders removes hop-by-hop headers that shouldn't be forwarded
func removeHopByHopHeaders(h http.Header) {
	hopByHopHeaders := []string{
		"Connection",
		"Keep-Alive",
		"Proxy-Authenticate",
		"Proxy-Authorization",
		"Te",
		"Trailers",
		"Transfer-Encoding",
		"Upgrade",
	}
	for _, hdr := range hopByHopHeaders {
		h.Del(hdr)
	}
}

// getScheme returns the scheme used in the original request
func getScheme(r *http.Request) string {
	if r.TLS != nil {
		return "https"
	}
	if proto := r.Header.Get("X-Forwarded-Proto"); proto != "" {
		return proto
	}
	return "http"
}
