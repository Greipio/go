package greip

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// ? Helper function to perform an HTTP GET request
func (g *Greip) getRequest(endpoint string, responseType interface{}, payload ...map[string]interface{}) error {
	baseURL := g.BaseURL
	urlEndpoint := fmt.Sprintf("%s%s", baseURL, endpoint)

	// Prepare headers
	req, err := http.NewRequest("GET", urlEndpoint, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", g.token))
	req.Header.Set("Content-Type", "application/json")

	// If test mode is enabled, add the 'mode' to the payload
	if g.test && len(payload) > 0 {
		if payload[0] == nil {
			payload[0] = make(map[string]interface{})
		}
		payload[0]["mode"] = "test"
	}

	// Construct query parameters from the payload
	query := req.URL.Query()
	for key, value := range payload[0] {
		query.Add(key, fmt.Sprintf("%v", value))
	}
	req.URL.RawQuery = query.Encode()

	// Execute the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check for non-2xx status codes
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("HTTP request failed with status code: %d", resp.StatusCode)
	}

	// Decode the JSON response
	var jsonResponse map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&jsonResponse); err != nil {
		return err
	}

	// Handle API-specific error in the response
	if status, ok := jsonResponse["status"].(string); ok && strings.ToLower(status) == "error" {
		description := jsonResponse["description"].(string)
		return fmt.Errorf("API error: %s", description)
	}

	// Extract the data and unmarshal it directly into responseType
	if data, ok := jsonResponse["data"]; ok {
		dataBytes, err := json.Marshal(data)
		if err != nil {
			return err
		}
		if err := json.Unmarshal(dataBytes, responseType); err != nil {
			return err
		}
	} else {
		return errors.New("invalid response format: missing data field")
	}

	return nil
}

// ? Helper function to perform an HTTP POST request
func (g *Greip) postRequest(endpoint string, responseType interface{}, payload map[string]interface{}) error {
	baseURL := g.BaseURL
	urlEndpoint := fmt.Sprintf("%s%s", baseURL, endpoint)

	// Prepare headers
	req, err := http.NewRequest("POST", urlEndpoint, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", g.token))
	req.Header.Set("Content-Type", "application/json")

	// If test mode is enabled, add the 'mode' to the payload
	if g.test {
		payload["mode"] = "test"
	}

	// Encode the payload
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	// Attach the payload to the request
	req.Body = io.NopCloser(bytes.NewReader(payloadBytes))

	// Execute the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check for non-2xx status codes
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("HTTP request failed with status code: %d", resp.StatusCode)
	}

	// Decode the JSON response
	var jsonResponse map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&jsonResponse); err != nil {
		return err
	}

	// Handle API-specific error in the response
	if status, ok := jsonResponse["status"].(string); ok && strings.ToLower(status) == "error" {
		description := jsonResponse["description"].(string)
		return fmt.Errorf("API error: %s", description)
	}

	// Extract the data and unmarshal it directly into responseType
	if data, ok := jsonResponse["data"]; ok {
		dataBytes, err := json.Marshal(data)
		if err != nil {
			return err
		}
		if err := json.Unmarshal(dataBytes, responseType); err != nil {
			return err
		}
	} else {
		return errors.New("invalid response format: missing data field")
	}

	return nil
}

// ? Helper function to validate params against the available list of parameters
func validateParams(params []string, availableParams []string) error {
	for _, param := range params {
		if !contains(availableParams, param) {
			return fmt.Errorf("invalid parameter: %s", param)
		}
	}
	return nil
}

// ? Helper function to validate the language (assuming only "EN" and "AR" are allowed)
func validateLang(lang string) error {
	allowedLangs := []string{"EN", "AR", "DE", "FR", "ES", "JA", "ZH", "RU"}
	if !contains(allowedLangs, strings.ToUpper(lang)) {
		return fmt.Errorf("invalid language: %s", lang)
	}
	return nil
}

// ? Helper function to check if a slice contains a specific value
func contains(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}
