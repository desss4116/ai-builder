package internet

import (
	"regexp"
)

func ExtractTitles(html string) []string {
	re := regexp.MustCompile(`<a[^>]*class="result__a"[^>]*>(.*?)</a>`)

	matches := re.FindAllStringSubmatch(html, -1)

	var results []string

	for _, match := range matches {
		if len(match) > 1 {
			results = append(results, match[1])
		}
	}

	return results
}
