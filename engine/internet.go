package engine

import (
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func SearchInternet(query string) string {

	searchURL :=
		"https://html.duckduckgo.com/html/?q=" +
			url.QueryEscape(query)

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	req, _ := http.NewRequest(
		"GET",
		searchURL,
		nil,
	)

	req.Header.Set(
		"User-Agent",
		"Mozilla/5.0",
	)

	resp, err := client.Do(req)

	if err != nil {
		return ""
	}

	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	return string(body)
}

func ExtractUsefulContent(
	html string,
) string {

	html = strings.ReplaceAll(
		html,
		"\n",
		" ",
	)

	if len(html) > 5000 {

		html = html[:5000]
	}

	return html
}
