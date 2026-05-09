package vector

import "strings"

func Similarity(
	query string,
	text string,
) int {

	score := 0

	queryWords :=
		strings.Fields(
			strings.ToLower(query),
		)

	lower :=
		strings.ToLower(text)

	for _, q := range queryWords {

		if strings.Contains(lower, q) {
			score++
		}
	}

	return score
}
