package internet

import (
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

func LiveSearch(query string) string {

	searchURL :=
		"https://html.duckduckgo.com/html/?q=" +
			url.QueryEscape(query)

	req, err := http.NewRequest(
		"GET",
		searchURL,
		nil,
	)

	if err != nil {
		return ""
	}

	/*
	   USER AGENT
	*/

	req.Header.Set(
		"User-Agent",
		"Mozilla/5.0",
	)

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		return ""
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return ""
	}

	html := string(body)

	/*
	   EXTRACT SEARCH RESULTS
	*/

	re := regexp.MustCompile(
		`(?s)<a[^>]*class="result__a"[^>]*>(.*?)</a>`,
	)

	matches := re.FindAllStringSubmatch(
		html,
		10,
	)

	if len(matches) == 0 {
		return ""
	}

	output := ""

	for _, match := range matches {

		text := match[1]

		/*
		   REMOVE HTML
		*/

		text = regexp.MustCompile(
			`<[^>]*>`,
		).ReplaceAllString(text, "")

		text = strings.TrimSpace(text)

		if text == "" {
			continue
		}

		output += "• " + text + "\n\n"
	}

	return strings.TrimSpace(output)
}
