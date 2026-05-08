package internet

import (
	"regexp"
	"strings"
)

func ScrapeContent(html string) string {

	re := regexp.MustCompile(`<[^>]*>`)

	clean := re.ReplaceAllString(html, "")

	clean = strings.ReplaceAll(clean, "\n", " ")

	clean = strings.ReplaceAll(clean, "\t", " ")

	if len(clean) > 4000 {
		clean = clean[:4000]
	}

	return clean
}
