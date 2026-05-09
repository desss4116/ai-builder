package internet

import (
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func Search(query string) []string {

	var results []string

	engines := []string{
		"https://html.duckduckgo.com/html/?q=",
		"https://lite.duckduckgo.com/lite/?q=",
	}

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	for _, engine := range engines {

		searchURL :=
			engine + url.QueryEscape(query)

		req, err := http.NewRequest(
			"GET",
			searchURL,
			nil,
		)

		if err != nil {
			continue
		}

		req.Header.Set(
			"User-Agent",
			"Mozilla/5.0",
		)

		req.Header.Set(
			"Accept-Language",
			"ru,en;q=0.9",
		)

		resp, err := client.Do(req)

		if err != nil {
			continue
		}

		body, err := io.ReadAll(resp.Body)

		resp.Body.Close()

		if err != nil {
			continue
		}

		html := string(body)

		if strings.Contains(
			strings.ToLower(html),
			"blocked",
		) {
			continue
		}

		doc, err :=
			goquery.NewDocumentFromReader(
				strings.NewReader(html),
			)

		if err != nil {
			continue
		}

		doc.Find("a").Each(func(
			i int,
			s *goquery.Selection,
		) {

			text :=
				strings.TrimSpace(
					s.Text(),
				)

			text =
				strings.ReplaceAll(
					text,
					"\n",
					" ",
				)

			if len(text) > 40 &&
				len(text) < 500 {

				results = append(
					results,
					text,
				)
			}
		})

		if len(results) > 0 {
			break
		}
	}

	clean := []string{}

	seen := map[string]bool{}

	for _, r := range results {

		r = strings.TrimSpace(r)

		if len(r) < 40 {
			continue
		}

		if seen[r] {
			continue
		}

		seen[r] = true

		clean = append(clean, r)

		if len(clean) >= 10 {
			break
		}
	}

	return clean
}
