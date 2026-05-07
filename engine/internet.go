package engine

import (
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type WebsiteAnalysis struct {
	IsDark bool
	HasPricing bool
	HasTestimonials bool
	HasGallery bool
}

func SearchInternet(query string) string {

	searchURL :=
		"https://html.duckduckgo.com/html/?q=" +
			url.QueryEscape(query)

	client := &http.Client{
		Timeout:10*time.Second,
	}

	req,_ := http.NewRequest(
		"GET",
		searchURL,
		nil,
	)

	req.Header.Set(
		"User-Agent",
		"Mozilla/5.0",
	)

	resp,err := client.Do(req)

	if err != nil {
		return ""
	}

	defer resp.Body.Close()

	body,_ := io.ReadAll(resp.Body)

	return string(body)
}

func AnalyzeWebsite(
	html string,
) WebsiteAnalysis {

	a := WebsiteAnalysis{}

	html =
		strings.ToLower(html)

	if strings.Contains(
		html,
		"testimonial",
	){
		a.HasTestimonials = true
	}

	if strings.Contains(
		html,
		"pricing",
	){
		a.HasPricing = true
	}

	if strings.Contains(
		html,
		"gallery",
	){
		a.HasGallery = true
	}

	if strings.Contains(
		html,
		"#000",
	){
		a.IsDark = true
	}

	return a
}
