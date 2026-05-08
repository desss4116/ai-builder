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

	knowledge := GetKnowledge(query)

	q := strings.ToLower(query)

	/*
	   TOP REQUESTS
	*/

	if strings.Contains(q, "топ") {

		return `
🔥 Лучшие результаты:

` + knowledge
	}

	/*
	   EDUCATION
	*/

	if strings.Contains(q, "что такое") {

		return `
📘 Объяснение:

` + knowledge
	}

	/*
	   WORD MEANING
	*/

	if strings.Contains(q, "значение") {

		return `
📖 Значение:

` + knowledge
	}

	/*
	   GENERAL
	*/

	return `
🌐 Ответ:

` + knowledge
}
