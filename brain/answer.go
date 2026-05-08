package brain

import (
	"strings"
)

func GenerateAnswer(query string, data string) string {

	q := strings.ToLower(query)

	// MOVIES

	if strings.Contains(q, "человек-паук") {

		return `🌐 Ответ:

Человек-паук —
один из самых популярных
супергероев Marvel.

Настоящее имя:
Питер Паркер.

После укуса радиоактивного паука
он получает сверхспособности:
• паучье чутьё
• сверхсилу
• ловкость
• возможность лазать по стенам

Лучшие фильмы:
• Spider-Man 2 (2004)
• No Way Home (2021)
• Into the Spider-Verse

Жанры:
Action / Sci-Fi / Superhero

Если нравятся фильмы Marvel,
также стоит посмотреть:
• Batman Begins
• Daredevil
• The Boys
• Invincible`
	}

	// GAMES

	if strings.Contains(q, "лучшие игры") {

		return `🌐 Ответ:

Топ современных игр:

1. Red Dead Redemption 2
— лучший open world и сюжет.

2. Cyberpunk 2077
— футуристический RPG-экшен.

3. Elden Ring
— один из лучших souls-like проектов.

4. GTA V
— культовый open world.

5. The Witcher 3
— легендарная RPG.

Если любишь:
• сюжет
• атмосферу
• открытый мир
то начни с:
Red Dead Redemption 2.`
	}

	// SERIALS

	if strings.Contains(q, "сериал") {

		return `🌐 Ответ:

Лучшие современные сериалы:

• Breaking Bad
• Better Call Saul
• Dark
• Mr. Robot
• The Last of Us
• True Detective

Если любишь:
• напряжение
• сюжетные повороты
• умные диалоги

начни с:
Breaking Bad.`
	}

	// fallback

	if len(data) > 1500 {
		data = data[:1500]
	}

	return "🌐 Ответ:\n\n" + data
}
