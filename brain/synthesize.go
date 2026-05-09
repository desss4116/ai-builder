package brain

import (
	"strings"
)

func Synthesize(
	query string,
	texts []string,
) string {

	q := strings.ToLower(query)

	// MOVIES

	if strings.Contains(q, "фильм") ||
		strings.Contains(q, "кино") {

		return `🌐 Ответ:

Топ фильмов, которые сейчас считаются лучшими:

• Interstellar
— культовая sci-fi драма Кристофера Нолана.

• Dune
— масштабная фантастика с невероятной атмосферой.

• Spider-Man: No Way Home
— один из лучших фильмов Marvel.

• Joker
— психологический триллер о происхождении Джокера.

• Oppenheimer
— мощная историческая драма.

Если любишь:
• глубокий сюжет
• атмосферу
• масштаб
• эмоции

то начни с:
Interstellar или Dune.`
	}

	// GAMES

	if strings.Contains(q, "игр") ||
		strings.Contains(q, "игры") {

		return `🌐 Ответ:

Лучшие современные игры:

• Red Dead Redemption 2
• Cyberpunk 2077
• Elden Ring
• GTA V
• The Witcher 3

Если любишь:
• открытый мир
• сюжет
• атмосферу

то начни с:
Red Dead Redemption 2.`
	}

	// SPIDER MAN

	if strings.Contains(q, "человек-паук") {

		return `🌐 Ответ:

Человек-паук —
один из самых популярных
супергероев Marvel.

Настоящее имя:
Питер Паркер.

После укуса радиоактивного паука
он получает:
• сверхсилу
• паучье чутьё
• возможность лазать по стенам

Лучшие фильмы:
• Spider-Man 2
• No Way Home
• Into the Spider-Verse

Также стоит посмотреть:
• Daredevil
• Batman Begins
• Invincible`
	}

	// GENERAL SYNTHESIS

	var final []string

	seen := map[string]bool{}

	for _, t := range texts {

		t = strings.TrimSpace(t)

		if len(t) < 80 {
			continue
		}

		if seen[t] {
			continue
		}

		seen[t] = true

		if len(t) > 350 {
			t = t[:350]
		}

		final = append(final, "• "+t)

		if len(final) >= 8 {
			break
		}
	}

	if len(final) == 0 {
		return "❌ Информация не найдена."
	}

	result := strings.Join(
		final,
		"\n\n",
	)

	if len(result) > 1800 {
		result = result[:1800]
	}

	return "🌐 Ответ:\n\n" + result
}
