package internet

import (
	"io"
	"net/http"
	"net/url"
)

func SearchInternet(query string) (string, error) {

	search :=
		"https://html.duckduckgo.com/html/?q=" +
			url.QueryEscape(query)

	req, _ := http.NewRequest(
		"GET",
		search,
		nil,
	)

	req.Header.Set(
		"User-Agent",
		"Mozilla/5.0",
	)

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	return string(body), nil
}
