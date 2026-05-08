package brain

import "strings"

func GetKnowledge(query string) string {

	q := strings.ToLower(query)

	/*
	   MOVIES
	*/

	if strings.Contains(q, "фильм") {

		return `
1. Interstellar
2. Oppenheimer
3. Dune Part Two
4. Blade Runner 2049
5. The Batman
6. Joker
7. Avatar
8. John Wick 4
9. Napoleon
10. The Creator
`
	}

	/*
	   PROGRAMMING
	*/

	if strings.Contains(q, "golang") {

		return `
Golang — язык программирования,
созданный Google.

Главные преимущества:
- высокая скорость
- многопоточность
- простота
- scalability
- backend performance
`
	}

	/*
	   APPLE
	*/

	if strings.Contains(q, "apple") {

		return `
Apple — одна из самых влиятельных
технологических компаний мира.

Известна:
- premium UX
- минимализмом
- экосистемой
- smooth animations
- high-end product design
`
	}

	/*
	   NETFLIX
	*/

	if strings.Contains(q, "netflix") {

		return `
Netflix использует:
- cinematic UI
- personalization
- recommendation systems
- dark premium design
- adaptive streaming
`
	}

	/*
	   AI
	*/

	if strings.Contains(q, "ии") ||
		strings.Contains(q, "ai") {

		return `
Искусственный интеллект —
это системы, способные:
- анализировать
- обучаться
- принимать решения
- генерировать контент
- автоматизировать задачи
`
	}

	/*
	   DEFAULT
	*/

	return `
Тема проанализирована.

Сейчас knowledge engine
ещё развивается,
но система уже умеет:
- анализировать запросы
- искать связи
- маршрутизировать intent
- строить AI pipeline
`
}
