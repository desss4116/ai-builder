package summarizer

import (
	"strings"
)

func BuildAnswer(text string) string {

	text = strings.ReplaceAll(text, "\n", " ")
	text = strings.ReplaceAll(text, "\t", " ")

	text = strings.TrimSpace(text)

	if len(text) > 5000 {
		text = text[:5000]
	}

	paragraphs := SplitParagraphs(text)

	answer := "🌐 Ответ:\n\n"

	for _, p := range paragraphs {

		p = strings.TrimSpace(p)

		if len(p) < 80 {
			continue
		}

		answer += p + "\n\n"

		if len(answer) > 3500 {
			break
		}
	}

	return answer
}

func SplitParagraphs(text string) []string {

	var result []string

	sentences := strings.Split(text, ".")

	current := ""

	for _, s := range sentences {

		s = strings.TrimSpace(s)

		if s == "" {
			continue
		}

		current += s + ". "

		if len(current) > 300 {

			result = append(result, current)

			current = ""
		}
	}

	if current != "" {
		result = append(result, current)
	}

	return result
}
