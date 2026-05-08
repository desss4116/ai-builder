package crawler

import (
	"io"
	"net/http"
	"regexp"
	"strings"
	"time"
)

func Crawl(url string) string {

	client := &http.Client{
		Timeout: 15 * time.Second,
	}

	req, err := http.NewRequest(
		"GET",
		url,
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

	// remove scripts
	reScript := regexp.MustCompile(`(?s)<script.*?</script>`)
	html = reScript.ReplaceAllString(html, "")

	// remove styles
	reStyle := regexp.MustCompile(`(?s)<style.*?</style>`)
	html = reStyle.ReplaceAllString(html, "")

	// remove tags
	reTags := regexp.MustCompile(`<[^>]+>`)
	text := reTags.ReplaceAllString(html, " ")

	// cleanup
	text = strings.ReplaceAll(text, "\n", " ")
	text = strings.ReplaceAll(text, "\r", " ")

	reSpace := regexp.MustCompile(`\s+`)
	text = reSpace.ReplaceAllString(text, " ")

	text = strings.TrimSpace(text)

	if len(text) > 5000 {
		text = text[:5000]
	}

	return text
}
