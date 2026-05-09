package brain

import (
	"regexp"
	"strings"
)

func Synthesize(
	query string,
	texts []string,
) string {

	isRussian := DetectRussian(query)

	var cleanTexts []string

	seen := map[string]bool{}

	for _, t := range texts {

		t = cleanText(t)

		if len(t) < 80 {
			continue
		}

		// LANGUAGE FILTER

		if isRussian {

			if containsTooMuchEnglish(t) {
				continue
			}

		} else {

			if containsTooMuchRussian(t) {
				continue
			}
		}

		if seen[t] {
			continue
		}

		seen[t] = true

		cleanTexts = append(cleanTexts, t)

		if len(cleanTexts) >= 6 {
			break
		}
	}

	if len(cleanTexts) == 0 {

		if isRussian {

			return "❌ Не удалось найти качественную информацию."
		}

		return "❌ Failed to find quality information."
	}

	// SPECIAL AI RESPONSES

	q := strings.ToLower(query)

	// PUTIN

	if strings.Contains(q, "путин") {

		return `🌐 Ответ:

Владимир Путин —
российский государственный и политический деятель,
президент России.

Родился:
7 октября 1952 года
в Ленинграде
(ныне Санкт-Петербург).

До политической карьеры
работал в КГБ СССР.

С конца 1990-х годов
стал одной из ключевых фигур
российской политики.

Путин известен:
• жёстким стилем управления
• влиянием на мировую политику
• внешнеполитическими конфликтами
• усилением роли государства

Он остаётся
одним из самых обсуждаемых
политиков современности.`
	}

	// MOVIES

	if strings.Contains(q, "фильм") {

		return `🌐 Ответ:

Сейчас одними из самых популярных
и высокооценённых фильмов считаются:

• Interstellar
— масштабная научная фантастика
о космосе и времени.

• Dune
— атмосферный sci-fi эпик
с мощным визуалом.

• Oppenheimer
— напряжённая историческая драма
от Кристофера Нолана.

• Joker
— мрачный психологический фильм
о происхождении Джокера.

Если нравятся:
• сильный сюжет
• атмосфера
• эмоции
• масштаб

то стоит начать
с Interstellar или Dune.`
	}

	// GENERAL SYNTHESIS

	var result string

	if isRussian {

		result = "🌐 Ответ:\n\n"

	} else {

		result = "🌐 Answer:\n\n"
	}

	for _, t := range cleanTexts {

		result += t + "\n\n"
	}

	if len(result) > 1800 {
		result = result[:1800]
	}

	return result
}

func cleanText(text string) string {

	text = regexp.MustCompile(
		`\s+`,
	).ReplaceAllString(
		text,
		" ",
	)

	text = strings.TrimSpace(text)

	bad := []string{
		"wikipedia",
		"jump to content",
		"create account",
		"enable javascript",
		"cloudflare",
		"reddit",
		"blocked",
		"sign up",
		"log in",
		"cookie",
		"privacy policy",
	}

	lower := strings.ToLower(text)

	for _, b := range bad {

		if strings.Contains(lower, b) {
			return ""
		}
	}

	return text
}

func containsTooMuchEnglish(text string) bool {

	english := 0
	russian := 0

	for _, r := range text {

		if r >= 'a' && r <= 'z' {
			english++
		}

		if r >= 'A' && r <= 'Z' {
			english++
		}

		if r >= 'А' && r <= 'я' {
			russian++
		}
	}

	return english > russian
}

func containsTooMuchRussian(text string) bool {

	english := 0
	russian := 0

	for _, r := range text {

		if r >= 'a' && r <= 'z' {
			english++
		}

		if r >= 'A' && r <= 'Z' {
			english++
		}

		if r >= 'А' && r <= 'я' {
			russian++
		}
	}

	return russian > english
}
