package internet

import (
	"ai-builder/crawler"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"
)

func Search(query string) string {

	client := &http.Client{
		Timeout: 20 * time.Second,
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
		return ""
	}

	req.Header.Set(
		"User-Agent",
		"Mozilla/5.0",
	)

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

	re := regexp.MustCompile(
		`uddg=([^"]+)`,
	)

	matches := re.FindAllStringSubmatch(
		html,
		-1,
	)

	if len(matches) == 0 {
		return ""
	}

	var finalText string

	used := 0

	for _, m := range matches {

		if used >= 5 {
			break
		}

		link, err := url.QueryUnescape(m[1])

		if err != nil {
			continue
		}

		lower := strings.ToLower(link)

		// FILTER GARBAGE LINKS

		if strings.Contains(lower, "duckduckgo") {
			continue
		}

		if strings.Contains(lower, "wikisource") {
			continue
		}

		if strings.Contains(lower, "wiktionary") {
			continue
		}

		if strings.Contains(lower, "wikiquote") {
			continue
		}

		if strings.Contains(lower, "special:search") {
			continue
		}

		if strings.Contains(lower, "javascript") {
			continue
		}

		if strings.Contains(lower, ".js") {
			continue
		}

		if strings.Contains(lower, "invalid") {
			continue
		}

		text := crawler.Crawl(link)

		if len(text) < 500 {
			continue
		}

		// FILTER BAD CONTENT

		bad := strings.ToLower(text)

		if strings.Contains(
			bad,
			"недопустимое название",
		) {
			continue
		}

		if strings.Contains(
			bad,
			"create account",
		) {
			continue
		}

		if strings.Contains(
			bad,
			"jump to content",
		) {
			continue
		}

		if strings.Contains(
			bad,
			"not have an article",
		) {
			continue
		}

		if strings.Contains(
			bad,
			"body { display:none",
		) {
			continue
		}

		if strings.Contains(
			bad,
			"enable javascript",
		) {
			continue
		}

		if len(text) > 2000 {
			text = text[:2000]
		}

		finalText += text + "\n\n"

		used++
	}

	if finalText == "" {
		return ""
	}

	finalText = regexp.MustCompile(
		`\s+`,
	).ReplaceAllString(
		finalText,
		" ",
	)

	if len(finalText) > 6000 {
		finalText = finalText[:6000]
	}

	return finalText
}
