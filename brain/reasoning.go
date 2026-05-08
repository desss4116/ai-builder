package brain

import "strings"

func DetectIntent(query string) string {

	q := strings.ToLower(query)

	if strings.Contains(q, "создай") &&
		strings.Contains(q, "сайт") {
		return "website"
	}

	if strings.Contains(q, "топ") {
		return "top"
	}

	if strings.Contains(q, "что такое") {
		return "education"
	}

	if strings.Contains(q, "значение слова") {
		return "dictionary"
	}

	return "general"
}

func Think(query string, data string) string {

	q := strings.ToLower(query)

	if strings.Contains(q, "фильм") {

		return `
🎬 Топ фильмы:

1. Dune Part Two
2. Oppenheimer
3. Interstellar
4. Blade Runner 2049
5. Joker
6. The Batman
7. Avatar
8. John Wick 4
9. Napoleon
10. The Creator
`
	}

	if strings.Contains(q, "кто такой") {

		return `
📘 Это человек или личность,
о которой найдена информация и проведён анализ.
`
	}

	if strings.Contains(q, "что такое") {

		return `
📖 Это термин или концепция,
которая означает следующее:

` + data
	}

	return `
🧠 Анализ завершён:

` + data
}
