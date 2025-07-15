package api

import (
	"encoding/json"
	"fmt"
	"go-test/models"
	"io"
	"net/http"
	"net/url"
	"os"
)

// Loads data from the api url
func LoadData(nextPage ...string) ([]models.StockInformation, error) {
	apiURL := os.Getenv("API_URL")
	token := os.Getenv("API_TOKEN")

	// Validate if exist the next page and if exist add it to the url
	if len(nextPage) > 0 && nextPage[0] != "" {
		parsedURL, err := url.Parse(apiURL)
		if err != nil {
			return nil, fmt.Errorf("error parsing base URL: %w", err)
		}
		query := parsedURL.Query()
		query.Set("next_page", nextPage[0])
		parsedURL.RawQuery = query.Encode()
		apiURL = parsedURL.String()
	}

	// Create the new request
	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return nil, err
	}

	// Add authorization header
	req.Header.Set("Authorization", "Bearer "+token)

	// Sent the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var data models.StockInformationResponse
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	// if the next page is not "", call the api with the next page
	if data.NextPage != "" {
		nextItems, err := LoadData(data.NextPage)
		if err != nil {
			return nil, err
		}
		// join the data
		data.Items = append(data.Items, nextItems...)
	}

	// return the data items
	return data.Items, nil
}
