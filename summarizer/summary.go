package summarizer

import "strings"

func Summarize(content string) string {

	content = strings.TrimSpace(content)

	if content == "" {

		return `
❌ Ничего не найдено.
`
	}

	/*
	   REMOVE GOOGLE GARBAGE
	*/

	content = strings.ReplaceAll(
		content,
		"Все",
		"",
	)

	content = strings.ReplaceAll(
		content,
		"Картинки",
		"",
	)

	content = strings.ReplaceAll(
		content,
		"Видео",
		"",
	)

	content = strings.ReplaceAll(
		content,
		"Новости",
		"",
	)

	/*
	   LIMIT
	*/

	if len(content) > 2000 {
		content = content[:2000]
	}

	return content
}
