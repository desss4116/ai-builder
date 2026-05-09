package internet

import (
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func Search(query string) []string {

	url :=
		"https://html.duckduckgo.com/html/?q=" +
			query

	client := &http.Client{
		Timeout: 15 * time.Second,
	}

	req, _ := http.NewRequest(
		"GET",
		url,
		nil,
	)

	req.Header.Set(
		"User-Agent",
		"Mozilla/5.0",
	)

	resp, err := client.Do(req)

	if err != nil {
		return nil
	}

	defer resp.Body.Close()

	doc, err :=
		goquery.NewDocumentFromReader(
			resp.Body,
		)

	if err != nil {
		return nil
	}

	var results []string

	doc.Find(".result").Each(func(
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

		if len(text) > 200 {

			results = append(
				results,
				text,
			)
		}
	})

	return results
}
