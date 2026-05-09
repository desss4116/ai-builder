package internet

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"
)

type WikiResponse struct {
	Extract string `json:"extract"`
}

func LiveSearch(query string) []string {

	results := duck(query)

	if len(results) > 0 {
		return results
	}

	results = wiki(query)

	if len(results) > 0 {
		return results
	}

	results = google(query)

	return results
}

func duck(query string) []string {

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	searchURL :=
		"https://html.duckduckgo.com/html/?q=" +
			url.QueryEscape(query)

	req, _ := http.NewRequest(
		"GET",
		searchURL,
		nil,
	)

	req.Header.Set(
		"User-Agent",
		"Mozilla/5.0",
	)

	resp, err := client.Do(req)

	if err != nil {
		return []string{}
	}

	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

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

		text :=
			strings.TrimSpace(m[1])

		if len(text) < 40 {
			continue
		}

		results = append(
			results,
			text,
		)

		if len(results) >= 5 {
			break
		}
	}

	return results
}

func wiki(query string) []string {

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	api :=
		"https://ru.wikipedia.org/api/rest_v1/page/summary/" +
			url.PathEscape(query)

	req, _ := http.NewRequest(
		"GET",
		api,
		nil,
	)

	req.Header.Set(
		"User-Agent",
		"Mozilla/5.0",
	)

	resp, err := client.Do(req)

	if err != nil {
		return []string{}
	}

	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var data WikiResponse

	err = json.Unmarshal(
		body,
		&data,
	)

	if err != nil {
		return []string{}
	}

	if len(data.Extract) < 40 {
		return []string{}
	}

	return []string{
		data.Extract,
	}
}

func google(query string) []string {

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	searchURL :=
		"https://www.google.com/search?q=" +
			url.QueryEscape(query)

	req, _ := http.NewRequest(
		"GET",
		searchURL,
		nil,
	)

	req.Header.Set(
		"User-Agent",
		"Mozilla/5.0",
	)

	resp, err := client.Do(req)

	if err != nil {
		return []string{}
	}

	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	text := string(body)

	re := regexp.MustCompile(
		`<span.*?>(.*?)</span>`,
	)

	matches :=
		re.FindAllStringSubmatch(
			text,
			-1,
		)

	var results []string

	for _, m := range matches {

		if len(m) < 2 {
			continue
		}

		t :=
			strings.TrimSpace(m[1])

		if len(t) < 80 {
			continue
		}

		results = append(
			results,
			t,
		)

		if len(results) >= 3 {
			break
		}
	}

	return results
}
