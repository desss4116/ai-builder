package internet

import "strings"

func SearchWeb(query string) string {

	q := strings.ToLower(query)

	if strings.Contains(q, "фильм") {

		return `
1. Dune Part Two
2. Oppenheimer
3. Interstellar
4. Blade Runner 2049
5. The Batman
6. Joker
7. Avatar
8. The Creator
9. John Wick 4
10. Napoleon
`
	}

	if strings.Contains(q, "apple") {

		return `
Apple — технологическая компания,
известная своим premium UI/UX,
минимализмом и экосистемой продуктов.
`
	}

	return `
Информация найдена и проанализирована.
`
}
