package brain

import "strings"

func DetectIntent(query string) string {

	q := strings.ToLower(query)

	if strings.Contains(q, "создай") &&
		strings.Contains(q, "сайт") {
		return "website"
	}

	if strings.Contains(q, "топ") {
		return "top"
	}

	if strings.Contains(q, "что такое") {
		return "education"
	}

	if strings.Contains(q, "значение слова") {
		return "dictionary"
	}

	return "general"
}
