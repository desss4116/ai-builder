package internet

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type DuckTopic struct {
	Text   string      `json:"Text"`
	Topics []DuckTopic `json:"Topics"`
}

type DuckResponse struct {
	AbstractText string      `json:"AbstractText"`
	Related      []DuckTopic `json:"RelatedTopics"`
}

func extractTopics(
	topics []DuckTopic,
	output *[]string,
) {

	for _, topic := range topics {

		/*
		   DIRECT TEXT
		*/

		if strings.TrimSpace(topic.Text) != "" {

			*output = append(
				*output,
				"• "+topic.Text,
			)
		}

		/*
		   NESTED TOPICS
		*/

		if len(topic.Topics) > 0 {

			extractTopics(
				topic.Topics,
				output,
			)
		}
	}
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
	   ABSTRACT
	*/

	if strings.TrimSpace(
		result.AbstractText,
	) != "" {

		return result.AbstractText
	}

	/*
	   RELATED TOPICS
	*/

	var outputs []string

	extractTopics(
		result.Related,
		&outputs,
	)

	if len(outputs) > 0 {

		if len(outputs) > 5 {
			outputs = outputs[:5]
		}

		return strings.Join(
			outputs,
			"\n\n",
		)
	}

	return ""
}
