package engine

import (
	"fmt"
	"net/http"
	"net/url"
	"io"
	"encoding/json"
	"time"
)

var mirrors = []string{"https://searx.be", "https://ononoki.org", "https://searx.work"}

func SearchInternet(query string) string {
	client := &http.Client{Timeout: 4 * time.Second}
	for _, m := range mirrors {
		u := fmt.Sprintf("%s/search?q=%s&format=json", m, url.QueryEscape(query))
		resp, err := client.Get(u)
		if err != nil || resp.StatusCode != 200 { continue }
		
		var res struct { Results []struct { Content string `json:"content"` } `json:"results"` }
		json.NewDecoder(resp.Body).Decode(&res)
		resp.Body.Close()

		if len(res.Results) > 0 {
			return res.Results[0].Content // Берем самый релевантный кусок
		}
	}
	return "Premium innovative solution for modern world"
}
