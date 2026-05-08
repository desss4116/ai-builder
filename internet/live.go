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
		return "❌ Ошибка поиска."
	}

	req.Header.Set(
		"User-Agent",
		"Mozilla/5.0",
	)

	resp, err := client.Do(req)

	if err != nil {
		return "❌ Ошибка подключения."
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return "❌ Ошибка чтения."
	}

	html := string(body)

	// extract urls
	re :=
		regexp.MustCompile(
			`uddg=([^"]+)`,
		)

	matches := re.FindAllStringSubmatch(
		html,
		-1,
	)

	if len(matches) == 0 {
		return "❌ Ничего не найдено."
	}

	var finalText string

	used := 0

	for _, m := range matches {

		if used >= 3 {
			break
		}

		link, err := url.QueryUnescape(m[1])

		if err != nil {
			continue
		}

		// skip garbage
		lower := strings.ToLower(link)

		if strings.Contains(lower, "duckduckgo") {
			continue
		}

		if strings.Contains(lower, "javascript") {
			continue
		}

		if strings.Contains(lower, ".js") {
			continue
		}

		// crawl page
		text := crawler.Crawl(link)

		if len(text) < 300 {
			continue
		}

		finalText += text + "\n\n"

		used++
	}

	if finalText == "" {
		return "❌ Не удалось получить контент."
	}

	// cleanup
	finalText = strings.ReplaceAll(
		finalText,
		"\n",
		" ",
	)

	finalText = regexp.MustCompile(
		`\s+`,
	).ReplaceAllString(
		finalText,
		" ",
	)

	// limit
	if len(finalText) > 5000 {
		finalText = finalText[:5000]
	}

	return finalText
}
