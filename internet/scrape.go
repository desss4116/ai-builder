package internet

import (
	"regexp"
)

func ScrapeContent(html string) string {

	re := regexp.MustCompile(`<[^>]*>`)

	clean := re.ReplaceAllString(html, "")

	if len(clean) > 3000 {
		clean = clean[:3000]
	}

	return clean
}
