package brain

import (
	"strings"
)

func GenerateAnswer(query string, data string) string {

	q := strings.ToLower(query)

	if strings.Contains(q, "человек-паук") {

		return `🌐 Ответ:

Человек-паук —
один из самых известных
супергероев Marvel.

Настоящее имя:
Питер Паркер.

После укуса радиоактивного паука
он получает способности:
• сверхсилу
• паучье чутьё
• возможность лазать по стенам

Лучшие фильмы:
• Spider-Man 2
• No Way Home
• Into the Spider-Verse

Если нравятся:
• Marvel
• супергерои
• экшен

то также стоит посмотреть:
• Batman Begins
• Daredevil
• Invincible`
	}

	if strings.Contains(q, "лучшие фильмы") {

		return `🌐 Ответ:

Топ современных фильмов:

1. Interstellar
— научная фантастика Кристофера Нолана.

2. Oppenheimer
— мощная историческая драма.

3. Dune
— один из лучших sci-fi фильмов последних лет.

4. Joker
— психологический триллер.

5. Spider-Man: No Way Home
— лучший современный Marvel crossover.

Если любишь:
• атмосферу
• масштаб
• сюжет

начни с:
Interstellar.`
	}

	if strings.Contains(q, "лучшие игры") {

		return `🌐 Ответ:

Топ современных игр:

1. Red Dead Redemption 2
2. Cyberpunk 2077
3. Elden Ring
4. GTA V
5. The Witcher 3

Если любишь:
• открытый мир
• сюжет
• атмосферу

начни с:
Red Dead Redemption 2.`
	}

	if len(data) > 1700 {
		data = data[:1700]
	}

	return "🌐 Ответ:\n\n" + data
}
