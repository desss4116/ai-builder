package brain

import "strings"

func DetectSearchIntent(
	query string,
) string {

	q := strings.ToLower(query)

	/*
	   MOVIES
	*/

	if strings.Contains(q, "кто") ||
		strings.Contains(q, "что такое") ||
		strings.Contains(q, "объясни") ||
		strings.Contains(q, "биография") {

		return "explanation"
	}

	/*
	   WATCH
	*/

	if strings.Contains(q, "смотреть") ||
		strings.Contains(q, "фильм") ||
		strings.Contains(q, "онлайн") {

		return "watch"
	}

	/*
	   TOPS
	*/

	if strings.Contains(q, "топ") ||
		strings.Contains(q, "лучшие") {

		return "top"
	}

	return "general"
}
