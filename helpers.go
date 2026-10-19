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

// ? Helper function to perform an HTTP GET request (v2 path + query string)
func (g *Greip) getRequest(path string, responseType interface{}, queryParams map[string]interface{}) error {
	baseURL := strings.TrimSuffix(g.BaseURL, "/")
	urlEndpoint := fmt.Sprintf("%s/%s", baseURL, strings.TrimPrefix(path, "/"))

	req, err := http.NewRequest("GET", urlEndpoint, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", g.token))
	req.Header.Set("Content-Type", "application/json")

	if queryParams == nil {
		queryParams = make(map[string]interface{})
	}
	if g.test {
		queryParams["mode"] = "test"
	}
	if _, ok := queryParams["source"]; !ok {
		queryParams["source"] = "Go-Package"
	}

	query := req.URL.Query()
	for key, value := range queryParams {
		if value == nil {
			continue
		}
		str := fmt.Sprintf("%v", value)
		if str == "" {
			continue
		}
		query.Add(key, str)
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

	if success, ok := jsonResponse["success"].(bool); ok && !success {
		description, _ := jsonResponse["description"].(string)
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

// ? Helper function to perform an HTTP POST request (v2 path + JSON body)
func (g *Greip) postRequest(path string, responseType interface{}, payload map[string]interface{}) error {
	baseURL := strings.TrimSuffix(g.BaseURL, "/")
	urlEndpoint := fmt.Sprintf("%s/%s", baseURL, strings.TrimPrefix(path, "/"))

	// Prepare headers
	req, err := http.NewRequest("POST", urlEndpoint, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", g.token))
	req.Header.Set("Content-Type", "application/json")

	if g.test {
		payload["mode"] = "test"
	}
	if _, ok := payload["source"]; !ok {
		payload["source"] = "Go-Package"
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

	if success, ok := jsonResponse["success"].(bool); ok && !success {
		description, _ := jsonResponse["description"].(string)
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
