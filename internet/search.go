package internet

import (
	"encoding/json"
	"net/http"
	"net/url"
	"time"
)

type SearchResult struct {
	Title string
	URL   string
}

type DuckResponse struct {
	RelatedTopics []struct {
		Text     string `json:"Text"`
		FirstURL string `json:"FirstURL"`
	} `json:"RelatedTopics"`
}

func Search(query string) []SearchResult {

	endpoint :=
		"https://api.duckduckgo.com/?q=" +
			url.QueryEscape(query) +
			"&format=json&no_html=1&skip_disambig=1"

	client := &http.Client{
		Timeout: 15 * time.Second,
	}

	resp, err := client.Get(endpoint)

	if err != nil {
		return nil
	}

	defer resp.Body.Close()

	var data DuckResponse

	err = json.NewDecoder(resp.Body).Decode(&data)

	if err != nil {
		return nil
	}

	var results []SearchResult

	for _, item := range data.RelatedTopics {

		if item.FirstURL == "" {
			continue
		}

		results = append(results, SearchResult{
			Title: item.Text,
			URL:   item.FirstURL,
		})

		if len(results) >= 5 {
			break
		}
	}

	return results
}
