package internet

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type Topic struct {
	Text string `json:"Text"`
}

type DuckResponse struct {
	Abstract     string  `json:"Abstract"`
	AbstractText string  `json:"AbstractText"`
	Heading      string  `json:"Heading"`
	Related      []Topic `json:"RelatedTopics"`
}

func LiveSearch(query string) string {

	endpoint :=
		"https://api.duckduckgo.com/?q=" +
			url.QueryEscape(query) +
			"&format=json&pretty=1"

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

	err = json.Unmarshal(body, &result)

	if err != nil {
		return ""
	}

	/*
	   MAIN ABSTRACT
	*/

	if result.AbstractText != "" {
		return result.AbstractText
	}

	if result.Abstract != "" {
		return result.Abstract
	}

	/*
	   RELATED TOPICS
	*/

	output := ""

	count := 0

	for _, topic := range result.Related {

		text := strings.TrimSpace(topic.Text)

		if text == "" {
			continue
		}

		output += "• " + text + "\n\n"

		count++

		if count >= 5 {
			break
		}
	}

	if output != "" {
		return output
	}

	return ""
}
