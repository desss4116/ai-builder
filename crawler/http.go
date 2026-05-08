package crawler

import (
	"io"
	"net/http"
)

func FetchURL(url string) string {

	resp, err := http.Get(url)

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
