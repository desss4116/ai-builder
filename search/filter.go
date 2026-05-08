package search

import "strings"

func IsBadContent(text string) bool {

	badPatterns := []string{
		"404",
		"not found",
		"create account",
		"enable javascript",
		"cloudflare",
		"access denied",
		"captcha",
		"body { display:none",
		"jump to content",
		"ошибка 404",
		"страница не найдена",
		"доступ запрещен",
		"зарегистрируйтесь",
		"signin",
		"log in",
		"login",
		"subscription required",
		"forbidden",
	}

	lower := strings.ToLower(text)

	for _, b := range badPatterns {

		if strings.Contains(lower, b) {
			return true
		}
	}

	if len(text) < 300 {
		return true
	}

	return false
}
