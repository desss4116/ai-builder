package main

import (
	"ai-builder/database"
	"ai-builder/engine"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"://github.com"
)

func main() {
	database.InitDB()
	
	dg, err := discordgo.New("Bot " + os.Getenv("DISCORD_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}

	dg.AddHandler(messageHandler)
	dg.Identify.Intents = discordgo.IntentsGuildMessages | discordgo.IntentsMessageContent

	if err = dg.Open(); err != nil {
		log.Fatal(err)
	}

	// Маршруты
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { http.ServeFile(w, r, "index.html") })
	http.HandleFunc("/site/", siteServeHandler)
	http.HandleFunc("/editor/", func(w http.ResponseWriter, r *http.Request) { http.ServeFile(w, r, "editor.html") })

	port := os.Getenv("PORT")
	if port == "" { port = "8080" }
	log.Printf("🚀 Intelligence OS System Online | Port: %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot { return }

	input := strings.TrimSpace(m.Content)
	lowInput := strings.ToLower(input)

	// 1. ЛОГИКА СОЗДАНИЯ САЙТА (Строгие триггеры)
	if isWebsiteRequest(lowInput) {
		s.ChannelTyping(m.ChannelID)
		
		// Генерация
		html := engine.GenerateWebsite(input)
		id := fmt.Sprintf("%d", time.Now().UnixNano())
		database.SaveGeneratedSite(id, input, html)

		domain := os.Getenv("DOMAIN")
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("✨ **Web-OS Engine:** Сайт готов!\n🌍 Ссылка: %s/site/%s", domain, id))
		return
	}

	// 2. ИНТЕЛЛЕКТУАЛЬНЫЙ ПОИСК (Для всех остальных вопросов)
	s.ChannelTyping(m.ChannelID)
	
	// Вызываем твой метод поиска
	searchResult := engine.SearchInternet(input)
	
	if len(searchResult) < 5 {
		s.ChannelMessageSend(m.ChannelID, "🔍 Информационный слой пуст. Попробуйте другой запрос.")
		return
	}

	// Отправляем РЕАЛЬНЫЙ ответ из интернета
	header := "🧠 **AI Research Layer:**\n\n"
	s.ChannelMessageSend(m.ChannelID, header + searchResult)
}

func isWebsiteRequest(msg string) bool {
    // Убираем лишние слова, оставляем только четкие команды на создание
	triggers := []string{"создай сайт", "сделай сайт", "build website", "create site", "сгенерируй лендинг"}
	for _, t := range triggers {
		if strings.Contains(msg, t) {
			return true
		}
	}
	return false
}

func siteServeHandler(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/site/")
	html := database.GetSiteHTML(id) 
	if html == "" {
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, html)
}
