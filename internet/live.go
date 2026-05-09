package internet

import (
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"
)

func Search(query string) []string {

	client := &http.Client{
		Timeout: 15 * time.Second,
	}

	searchURL :=
		"https://html.duckduckgo.com/html/?q=" +
			url.QueryEscape(query)

	req, err := http.NewRequest(
		"GET",
		searchURL,
		nil,
	)

	if err != nil {
		return []string{}
	}

	req.Header.Set(
		"User-Agent",
		"Mozilla/5.0",
	)

	resp, err := client.Do(req)

	if err != nil {
		return []string{}
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return []string{}
	}

	html := string(body)

	// extract snippets
	re := regexp.MustCompile(
		`result__snippet.*?>(.*?)<`,
	)

	matches := re.FindAllStringSubmatch(
		html,
		-1,
	)

	var results []string

	for _, m := range matches {

		if len(m) < 2 {
			continue
		}

		text := m[1]

		// remove tags
		text = regexp.MustCompile(
			`<[^>]+>`,
		).ReplaceAllString(
			text,
			"",
		)

		text = strings.TrimSpace(text)

		if len(text) < 50 {
			continue
		}

		// filter garbage
		lower := strings.ToLower(text)

		if strings.Contains(lower, "enable javascript") {
			continue
		}

		if strings.Contains(lower, "sign up") {
			continue
		}

		if strings.Contains(lower, "create account") {
			continue
		}

		if strings.Contains(lower, "blocked") {
			continue
		}

		results = append(results, text)

		if len(results) >= 10 {
			break
		}
	}

	return results
}
