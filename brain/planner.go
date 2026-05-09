package brain

import "strings"

func DetectIntent(query string) string {

	q := strings.ToLower(query)

	if strings.Contains(q, "фильм") ||
		strings.Contains(q, "movie") {
		return "movies"
	}

	if strings.Contains(q, "игр") ||
		strings.Contains(q, "game") {
		return "games"
	}

	if strings.Contains(q, "кто") {
		return "person"
	}

	if strings.Contains(q, "что") {
		return "definition"
	}

	if strings.Contains(q, "посоветуй") {
		return "recommendation"
	}

	return "general"
}
