package internet

import (
	"ai-builder/crawler"
	"ai-builder/search"
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

	trustedDomains := []string{
		"wikipedia.org",
		"imdb.com",
		"kinopoisk",
		"fandom.com",
		"reddit.com",
		"github.com",
		"medium.com",
		"britannica.com",
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

		allowed := false

		for _, domain := range trustedDomains {

			if strings.Contains(lower, domain) {
				allowed = true
				break
			}
		}

		if !allowed {
			continue
		}

		if strings.Contains(lower, "duckduckgo") {
			continue
		}

		if strings.Contains(lower, "javascript") {
			continue
		}

		if strings.Contains(lower, ".js") {
			continue
		}

		text := crawler.Crawl(link)

		if search.IsBadContent(text) {
			continue
		}

		if len(text) > 2000 {
			text = text[:2000]
		}

		finalText += text + "\n\n"

		used++
	}

	if finalText == "" {
		return "❌ Не удалось получить информацию."
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
