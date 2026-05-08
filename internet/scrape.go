package internet

import (
	"io"
	"net/http"
	"regexp"
	"strings"
)

func ScrapePage(url string) string {

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
	   REMOVE TAGS
	*/

	re := regexp.MustCompile(`<[^>]*>`)

	text :=
		re.ReplaceAllString(
			html,
			" ",
		)

	text = strings.TrimSpace(text)

	/*
	   CLEAN SPACES
	*/

	text = regexp.MustCompile(`\s+`).
		ReplaceAllString(text, " ")

	if len(text) > 3000 {
		text = text[:3000]
	}

	return text
}
