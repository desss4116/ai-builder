package search

import (
	"strings"
)

func SemanticSearch(query string) string {

	q := strings.ToLower(query)

	score := 0

	if strings.Contains(q, "ai") {
		score += 10
	}

	if strings.Contains(q, "сайт") {
		score += 20
	}

	if strings.Contains(q, "фильм") {
		score += 15
	}

	return `
Semantic ranking completed.

Relevance score:
` + string(rune(score))
}
