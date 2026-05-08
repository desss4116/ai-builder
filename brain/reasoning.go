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

	data = strings.TrimSpace(data)

	if data == "" {

		return `
❌ Информация не найдена.
`
	}

	return `
🌐 Ответ:

` + data
}
