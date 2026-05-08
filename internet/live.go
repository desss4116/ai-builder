package internet

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

type DuckResponse struct {
	Abstract string `json:"Abstract"`
	Heading  string `json:"Heading"`
}

func LiveSearch(query string) string {

	endpoint :=
		"https://api.duckduckgo.com/?q=" +
			url.QueryEscape(query) +
			"&format=json&no_html=1"

	resp, err := http.Get(endpoint)

	if err != nil {
		return ""
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return ""
	}

	var result DuckResponse

	json.Unmarshal(body, &result)

	if result.Abstract != "" {
		return result.Abstract
	}

	if result.Heading != "" {
		return result.Heading
	}

	return ""
}
