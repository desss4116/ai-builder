package internet

import (
	"io"
	"net/http"
	"time"
)

func FetchPage(url string) string {

	client := &http.Client{
		Timeout: 15 * time.Second,
	}

	req, err := http.NewRequest("GET", url, nil)
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

	return string(body)
}
