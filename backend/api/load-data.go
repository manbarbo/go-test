package api

import (
	"encoding/json"
	"fmt"
	"go-test/models"
	"io"
	"net/http"
	"net/url"
	"os"
	"sync"
)

// Loads data from the api url
func LoadData(insertFn func(*models.StockInformation) error, nextPage ...string) error {
	apiURL := os.Getenv("API_URL")
	token := os.Getenv("API_TOKEN")

	// Validate if exist the next page and if exist add it to the url
	if len(nextPage) > 0 && nextPage[0] != "" {
		parsedURL, err := url.Parse(apiURL)
		if err != nil {
			return fmt.Errorf("error parsing base URL: %w", err)
		}
		query := parsedURL.Query()
		query.Set("next_page", nextPage[0])
		parsedURL.RawQuery = query.Encode()
		apiURL = parsedURL.String()
	}

	// Create the new request
	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return err
	}

	// Add authorization header
	req.Header.Set("Authorization", "Bearer "+token)

	// Sent the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Read the response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var data models.StockInformationResponse
	if err := json.Unmarshal(body, &data); err != nil {
		return err
	}

	// Insert the elements
	var wg sync.WaitGroup
	errChan := make(chan error, len(data.Items))
	for _, item := range data.Items {
		itemCopy := item
		wg.Add(1)
		go func(stock models.StockInformation) {
			defer wg.Done()
			if err := insertFn(&stock); err != nil {
				errChan <- fmt.Errorf("insert error %s: %w", stock.Ticker, err)
			}
		}(itemCopy)
	}

	wg.Wait()
	close(errChan)

	if len(errChan) > 0 {
		for err := range errChan {
			fmt.Println("Error:", err)
		}
		return fmt.Errorf("errors occurred during insertion")
	}

	// if the next page is not "", call the api with the next page
	if data.NextPage != "" {
		return LoadData(insertFn, data.NextPage)
	}

	// return the data items
	return nil
}
