package engine

import (
	"io"
	"net/http"
)

func FetchWebsite(url string) string {

	resp, err := http.Get(url)

	if err != nil {
		return ""
	}

	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	return string(body)
}
