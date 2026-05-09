package brain

import "strings"

func Recommend(query string) string {

	q := strings.ToLower(query)

	// MOVIES

	if strings.Contains(q, "фильм") {

		return `🎬 Рекомендации:

• Interstellar
— глубокая sci-fi драма
о космосе и времени.

• Dune
— масштабная фантастика
с мощной атмосферой.

• Joker
— мрачный психологический фильм.

• Oppenheimer
— напряжённая историческая драма.

• Blade Runner 2049
— стильный киберпанк
с невероятной атмосферой.`
	}

	// GAMES

	if strings.Contains(q, "игр") {

		return `🎮 Рекомендации:

• Red Dead Redemption 2
• Cyberpunk 2077
• Elden Ring
• GTA V
• The Witcher 3

Лучший старт:
Red Dead Redemption 2.`
	}

	return ""
}
