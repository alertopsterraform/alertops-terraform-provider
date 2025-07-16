package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/hashicorp/go-retryablehttp"
)

type Client struct {
	apiKey     string
	baseURL    string
	httpClient *retryablehttp.Client
}

func NewClient(apiKey, baseURL string) *Client {
	retryClient := retryablehttp.NewClient()
	retryClient.RetryMax = 3
	retryClient.RetryWaitMin = 1 * time.Second
	retryClient.RetryWaitMax = 30 * time.Second

	return &Client{
		apiKey:     apiKey,
		baseURL:    strings.TrimRight(baseURL, "/"),
		httpClient: retryClient,
	}
}

func (c *Client) doRequest(ctx context.Context, method, path string, body interface{}) (*http.Response, error) {
	var bodyReader io.Reader
	var requestJSON []byte

	if body != nil {
		var err error
		requestJSON, err = json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal request body: %w", err)
		}
		bodyReader = bytes.NewReader(requestJSON)
	}

	url := fmt.Sprintf("%s%s", c.baseURL, path)
	req, err := retryablehttp.NewRequestWithContext(ctx, method, url, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("api-key", c.apiKey)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	
	// Debug: Show exactly what's being sent
	log.Printf("DEBUG: api-key header being sent: '%s'", c.apiKey)
	log.Printf("DEBUG: API key length: %d", len(c.apiKey))

	// Debug logging: print request details
	log.Printf("DEBUG: Making %s request to %s", method, url)
	log.Printf("DEBUG: Using API key: %s", c.apiKey)
	if requestJSON != nil {
		log.Printf("DEBUG: Request JSON: %s", string(requestJSON))
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	// Read response body for debugging
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("DEBUG: Failed to read response body: %v", err)
		return resp, nil
	}
	resp.Body.Close()

	// Debug logging: print response details
	log.Printf("DEBUG: Response status: %d", resp.StatusCode)
	log.Printf("DEBUG: Response body: %s", string(responseBody))

	// Recreate the response body reader for the calling function
	resp.Body = io.NopCloser(bytes.NewReader(responseBody))

	return resp, nil
}

func (c *Client) get(ctx context.Context, path string, result interface{}) error {
	resp, err := c.doRequest(ctx, "GET", path, nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		// Read the response body for error details
		bodyBytes, _ := io.ReadAll(resp.Body)
		fullURL := fmt.Sprintf("%s%s", c.baseURL, path)
		
		return fmt.Errorf("GET request failed:\n"+
			"━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━\n"+
			"URL: %s\n"+
			"METHOD: GET\n"+
			"HEADERS:\n"+
			"  api-key: %s\n"+
			"  Accept: application/json\n"+
			"━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━\n"+
			"RESPONSE STATUS: %d\n"+
			"RESPONSE BODY:\n%s\n"+
			"━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━", 
			fullURL, c.apiKey, resp.StatusCode, string(bodyBytes))
	}

	return json.NewDecoder(resp.Body).Decode(result)
}

func (c *Client) post(ctx context.Context, path string, body, result interface{}) error {
	resp, err := c.doRequest(ctx, "POST", path, body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		// Read the response body for error details
		bodyBytes, _ := io.ReadAll(resp.Body)
		
		// Include full request details in error for debugging
		requestJSON := "null"
		if body != nil {
			if jsonBytes, err := json.Marshal(body); err == nil {
				requestJSON = string(jsonBytes)
			}
		}
		
		fullURL := fmt.Sprintf("%s%s", c.baseURL, path)
		
		return fmt.Errorf("🔥🔥🔥 API CALL DEBUG - FIXED AUTHENTICATION 🔥🔥🔥\n"+
			"URL: %s\n"+
			"METHOD: POST\n"+
			"HEADERS:\n"+
			"  api-key: %s\n"+
			"  Content-Type: application/json\n"+
			"  Accept: application/json\n"+
			"REQUEST BODY:\n%s\n"+
			"RESPONSE STATUS: %d\n"+
			"RESPONSE BODY:\n%s", 
			fullURL, c.apiKey, requestJSON, resp.StatusCode, string(bodyBytes))
	}

	if result != nil {
		return json.NewDecoder(resp.Body).Decode(result)
	}

	return nil
}

func (c *Client) put(ctx context.Context, path string, body, result interface{}) error {
	resp, err := c.doRequest(ctx, "PUT", path, body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		// Read the response body for error details
		bodyBytes, _ := io.ReadAll(resp.Body)
		
		// Include full request details in error for debugging
		requestJSON := "null"
		if body != nil {
			if jsonBytes, err := json.Marshal(body); err == nil {
				requestJSON = string(jsonBytes)
			}
		}
		
		fullURL := fmt.Sprintf("%s%s", c.baseURL, path)
		
		return fmt.Errorf("PUT request failed:\n"+
			"━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━\n"+
			"URL: %s\n"+
			"METHOD: PUT\n"+
			"HEADERS:\n"+
			"  api-key: %s\n"+
			"  Content-Type: application/json\n"+
			"  Accept: application/json\n"+
			"REQUEST BODY:\n%s\n"+
			"━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━\n"+
			"RESPONSE STATUS: %d\n"+
			"RESPONSE BODY:\n%s\n"+
			"━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━", 
			fullURL, c.apiKey, requestJSON, resp.StatusCode, string(bodyBytes))
	}

	if result != nil {
		return json.NewDecoder(resp.Body).Decode(result)
	}

	return nil
}

func (c *Client) delete(ctx context.Context, path string) error {
	resp, err := c.doRequest(ctx, "DELETE", path, nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		// Read the response body for error details
		bodyBytes, _ := io.ReadAll(resp.Body)
		fullURL := fmt.Sprintf("%s%s", c.baseURL, path)
		
		return fmt.Errorf("DELETE request failed:\n"+
			"━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━\n"+
			"URL: %s\n"+
			"METHOD: DELETE\n"+
			"HEADERS:\n"+
			"  api-key: %s\n"+
			"  Accept: application/json\n"+
			"━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━\n"+
			"RESPONSE STATUS: %d\n"+
			"RESPONSE BODY:\n%s\n"+
			"━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━", 
			fullURL, c.apiKey, resp.StatusCode, string(bodyBytes))
	}

	return nil
} 