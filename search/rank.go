package search

import (
	"sort"
	"strings"
)

type Ranked struct {
	Text  string
	Score int
}

func Rank(query string, paragraphs []string) []string {

	query = strings.ToLower(query)

	var ranked []Ranked

	for _, p := range paragraphs {

		score := 0

		words := strings.Fields(query)

		for _, w := range words {

			if strings.Contains(
				strings.ToLower(p),
				w,
			) {
				score++
			}
		}

		if score > 0 {

			ranked = append(ranked, Ranked{
				Text:  p,
				Score: score,
			})
		}
	}

	sort.Slice(ranked, func(i, j int) bool {
		return ranked[i].Score > ranked[j].Score
	})

	var result []string

	for i, r := range ranked {

		if i >= 5 {
			break
		}

		result = append(result, r.Text)
	}

	return result
}
