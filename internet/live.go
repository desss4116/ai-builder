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

	re := regexp.MustCompile(
		`result__snippet.*?>(.*?)<`,
	)

	matches :=
		re.FindAllStringSubmatch(
			html,
			-1,
		)

	var results []string

	for _, m := range matches {

		if len(m) < 2 {
			continue
		}

		text := m[1]

		text = regexp.MustCompile(
			`<[^>]+>`,
		).ReplaceAllString(
			text,
			"",
		)

		text = strings.TrimSpace(text)

		if len(text) < 80 {
			continue
		}

		lower := strings.ToLower(text)

		bad := []string{
			"reddit",
			"blocked",
			"sign up",
			"cookie",
			"privacy policy",
			"enable javascript",
			"jump to content",
		}

		skip := false

		for _, b := range bad {

			if strings.Contains(lower, b) {
				skip = true
			}
		}

		if skip {
			continue
		}

		results = append(results, text)

		if len(results) >= 10 {
			break
		}
	}

	return results
}
