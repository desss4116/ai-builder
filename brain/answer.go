package brain

import (
	"strings"
)

func GenerateAnswer(query string, internetData string) string {

	q := strings.ToLower(query)

	if strings.Contains(q, "топ") {
		return BuildTopAnswer(query, internetData)
	}

	if strings.Contains(q, "что такое") {
		return BuildEducationAnswer(query, internetData)
	}

	if strings.Contains(q, "значение слова") {
		return BuildDictionaryAnswer(query, internetData)
	}

	return BuildGeneralAnswer(query, internetData)
}

func BuildTopAnswer(query string, data string) string {
	return "🔥 Лучшие результаты по запросу:\n\n" + data
}

func BuildEducationAnswer(query string, data string) string {
	return "📘 Объяснение:\n\n" + data
}

func BuildDictionaryAnswer(query string, data string) string {
	return "📖 Значение слова:\n\n" + data
}

func BuildGeneralAnswer(query string, data string) string {
	return "🌐 Ответ:\n\n" + data
}
