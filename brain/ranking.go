package brain

import (
	"sort"
	"strings"
)

type RankedText struct {
	Text  string
	Score int
}

func RankTexts(
	query string,
	texts []string,
) []string {

	queryWords :=
		strings.Fields(
			strings.ToLower(query),
		)

	var ranked []RankedText

	for _, t := range texts {

		score := 0

		lower := strings.ToLower(t)

		// relevance

		for _, q := range queryWords {

			if strings.Contains(lower, q) {
				score += 15
			}
		}

		// knowledge quality

		good := []string{
			"это",
			"является",
			"фильм",
			"игра",
			"супергерой",
			"история",
			"родился",
			"marvel",
			"dc",
		}

		for _, g := range good {

			if strings.Contains(lower, g) {
				score += 10
			}
		}

		// garbage

		bad := []string{
			"youtube",
			"rutube",
			"tiktok",
			"playlist",
			"cookie",
			"privacy",
			"blocked",
			"sign up",
			"reddit",
			"jump to content",
		}

		for _, b := range bad {

			if strings.Contains(lower, b) {
				score -= 30
			}
		}

		if len(t) > 120 {
			score += 5
		}

		ranked = append(
			ranked,
			RankedText{
				Text:  t,
				Score: score,
			},
		)
	}

	sort.Slice(
		ranked,
		func(i, j int) bool {
			return ranked[i].Score >
				ranked[j].Score
		},
	)

	var final []string

	for _, r := range ranked {

		if r.Score <= 0 {
			continue
		}

		final = append(final, r.Text)

		if len(final) >= 5 {
			break
		}
	}

	return final
}
