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

	q := strings.ToLower(query)

	if strings.Contains(q, "фильм") {

		return `
🎬 Топ фильмы:

1. Interstellar
2. Oppenheimer
3. Dune Part Two
4. Blade Runner 2049
5. Joker
6. The Batman
7. Avatar
8. John Wick 4
9. Napoleon
10. The Creator
`
	}

	if strings.Contains(q, "apple") {

		return `
🍎 Apple:

Apple — одна из самых влиятельных
технологических компаний мира.

Главные особенности:
- premium UI/UX
- минимализм
- экосистема
- плавные анимации
- high-end branding
`
	}

	if strings.Contains(q, "netflix") {

		return `
🎬 Netflix UI Architecture:

- cinematic layout
- dark gradients
- hover previews
- recommendation engine
- adaptive personalization
- smooth transitions
`
	}

	if strings.Contains(q, "что такое") {

		return `
📘 Объяснение:

` + data
	}

	return `
🌐 AI Analysis:

` + data
}
