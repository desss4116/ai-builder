package brain

import (
	"strings"
)

func Synthesize(
	query string,
	texts []string,
) string {

	if len(texts) == 0 {
		return "❌ Информация не найдена."
	}

	queryLower :=
		strings.ToLower(query)

	// CLEAN TEXTS

	cleaned := []string{}

	for _, t := range texts {

		t = strings.ReplaceAll(t, "\n", " ")
		t = strings.TrimSpace(t)

		if len(t) < 40 {
			continue
		}

		lower := strings.ToLower(t)

		// FILTER TRASH

		bad := []string{
			"cookie",
			"privacy",
			"blocked",
			"sign up",
			"youtube",
			"rutube",
			"playlist",
			"javascript",
			"404",
			"login",
		}

		skip := false

		for _, b := range bad {

			if strings.Contains(lower, b) {
				skip = true
			}
		}

		if skip {
			continue
		}

		// ONLY SAME LANGUAGE STYLE

		if isEnglish(queryLower) {

			if containsRussian(t) {
				continue
			}

		} else {

			if containsEnglishHeavy(t) {
				continue
			}
		}

		cleaned = append(cleaned, t)
	}

	if len(cleaned) == 0 {
		return "❌ Нормальная информация не найдена."
	}

	// KNOWLEDGE FUSION

	final :=
		"🌐 Ответ:\n\n"

	used := map[string]bool{}

	for _, t := range cleaned {

		sentences :=
			strings.Split(t, ".")

		for _, s := range sentences {

			s = strings.TrimSpace(s)

			if len(s) < 35 {
				continue
			}

			lower :=
				strings.ToLower(s)

			if used[lower] {
				continue
			}

			used[lower] = true

			final += s + ".\n\n"

			if len(final) > 1400 {
				return final
			}
		}
	}

	return final
}

func containsRussian(s string) bool {

	for _, r := range s {

		if r >= 'А' && r <= 'я' {
			return true
		}
	}

	return false
}

func containsEnglishHeavy(s string) bool {

	count := 0

	for _, r := range s {

		if (r >= 'A' && r <= 'Z') ||
			(r >= 'a' && r <= 'z') {

			count++
		}
	}

	return count > 25
}

func isEnglish(s string) bool {

	count := 0

	for _, r := range s {

		if (r >= 'A' && r <= 'Z') ||
			(r >= 'a' && r <= 'z') {

			count++
		}
	}

	return count > 2
}
