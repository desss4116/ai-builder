package internet

import (
	"regexp"
)

func ExtractTitles(html string) []string {

	re := regexp.MustCompile(
		`<a.*?>(.*?)</a>`,
	)

	matches := re.FindAllStringSubmatch(
		html,
		20,
	)

	var results []string

	for _, m := range matches {

		if len(m) > 1 {

			text := StripTags(m[1])

			if len(text) > 15 {
				results = append(results, text)
			}
		}
	}

	return results
}

func StripTags(s string) string {

	re := regexp.MustCompile(`<.*?>`)

	return re.ReplaceAllString(s, "")
}
