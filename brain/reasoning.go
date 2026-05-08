package brain

import "strings"

func DetectIntent(query string) string {

	q := strings.ToLower(query)

	if strings.Contains(q, "создай") &&
		strings.Contains(q, "сайт") {

		return "website"
	}

	return "general"
}

func Think(query string, data string) string {

	if data != "" {

		return `
🌐 Ответ:

` + data
	}

	return `
❌ Информация пока не найдена.
`
}
