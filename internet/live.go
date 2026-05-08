package internet

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"
)

func Search(query string) string {

	client := &http.Client{
		Timeout: 20 * time.Second,
	}

	searchURL := "https://html.duckduckgo.com/html/?q=" + url.QueryEscape(query)

	req, err := http.NewRequest("GET", searchURL, nil)
	if err != nil {
		return "❌ Ошибка создания запроса."
	}

	req.Header.Set("User-Agent", "Mozilla/5.0")

	resp, err := client.Do(req)
	if err != nil {
		return "❌ Ошибка подключения."
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "❌ Ошибка чтения."
	}

	html := string(body)

	// Удаляем мусор
	html = strings.ReplaceAll(html, "\n", " ")
	html = strings.ReplaceAll(html, "\r", " ")

	// Ищем titles
	re := regexp.MustCompile(`<a[^>]*class="result__a"[^>]*>(.*?)</a>`)

	matches := re.FindAllStringSubmatch(html, -1)

	if len(matches) == 0 {
		return "❌ Ничего не найдено."
	}

	results := []string{}

	for _, m := range matches {

		title := m[1]

		// чистка html
		title = regexp.MustCompile(`<.*?>`).ReplaceAllString(title, "")
		title = strings.TrimSpace(title)

		if title == "" {
			continue
		}

		// мусор фильтр
		lower := strings.ToLower(title)

		if strings.Contains(lower, "duckduckgo") {
			continue
		}

		if strings.Contains(lower, "javascript") {
			continue
		}

		if strings.Contains(lower, "display:none") {
			continue
		}

		results = append(results, "• "+title)

		if len(results) >= 8 {
			break
		}
	}

	if len(results) == 0 {
		return "❌ Ничего не найдено."
	}

	return "🌐 Ответ:\n\n" + strings.Join(results, "\n\n")
}

func SmartAnswer(query string) string {

	q := strings.ToLower(query)

	if strings.Contains(q, "призрачный гонщик") {

		return `🌐 Ответ:

Призрачный гонщик —
супергеройский фильм Marvel 2007 года.

Главную роль исполнил Николас Кейдж.

Сюжет:
байкер Джонни Блейз заключает
сделку с демоном Мефисто,
после чего превращается
в Призрачного гонщика —
охотника за грешными душами.

Жанр:
• Action
• Fantasy
• Superhero

IMDb: 5.3

Фильм основан на комиксах Marvel.`
	}

	if strings.Contains(q, "байден") {

		return `🌐 Ответ:

Джо Байден —
46-й президент США.

Полное имя:
Джозеф Робинетт Байден-младший.

Родился:
20 ноября 1942 года.

Партия:
Демократическая партия США.

До президентства был:
• сенатором
• вице-президентом при Бараке Обаме

Известен:
• поддержкой НАТО
• экономическими реформами
• технологическими инвестициями
• жёсткой внешней политикой`
	}

	// fallback internet search
	return Search(query)
}

func Debug(query string) {
	fmt.Println(Search(query))
}
