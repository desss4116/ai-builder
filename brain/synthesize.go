package brain

import (
	"sort"
	"strings"
)

type Chunk struct {
	Text  string
	Score int
}

func Score(query string, text string) int {
	query = strings.ToLower(query)
	text = strings.ToLower(text)

	score := 0

	words := strings.Fields(query)

	for _, w := range words {
		if strings.Contains(text, w) {
			score += 10
		}
	}

	score += len(text) / 300

	return score
}

func Synthesize(query string, texts []string) string {

	var chunks []Chunk

	for _, t := range texts {

		if len(t) < 120 {
			continue
		}

		s := Score(query, t)

		chunks = append(chunks, Chunk{
			Text:  t,
			Score: s,
		})
	}

	sort.Slice(chunks, func(i, j int) bool {
		return chunks[i].Score > chunks[j].Score
	})

	final := ""

	limit := 0

	for _, c := range chunks {

		if limit >= 3 {
			break
		}

		if strings.Contains(final, c.Text) {
			continue
		}

		final += c.Text + "\n\n"

		limit++
	}

	if final == "" {
		return "Ничего не найдено."
	}

	return final
}
