package brain

import "unicode"

func DetectRussian(text string) bool {

	russian := 0
	english := 0

	for _, r := range text {

		if unicode.Is(unicode.Cyrillic, r) {
			russian++
		}

		if unicode.IsLetter(r) &&
			r <= unicode.MaxASCII {
			english++
		}
	}

	return russian >= english
}
