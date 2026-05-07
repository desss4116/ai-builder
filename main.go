package main

import (
	"ai-builder/database"
	"ai-builder/engine"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"://github.com"
)

func main() {
	database.InitDB()
	
	token := os.Getenv("DISCORD_TOKEN")
	dg, err := discordgo.New("Bot " + token)
	if err != nil { log.Fatal(err) }

	dg.AddHandler(messageHandler)
	dg.Identify.Intents = discordgo.IntentsGuildMessages | discordgo.IntentsMessageContent
	
	if err = dg.Open(); err != nil { log.Fatal(err) }

	// HTTP Сервер для Railway
	http.HandleFunc("/site/", func(w http.ResponseWriter, r *http.Request) {
		id := strings.TrimPrefix(r.URL.Path, "/site/")
		// Здесь логика выдачи HTML из базы данных
		html := database.GetSiteHTML(id) 
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprint(w, html)
	})

	port := os.Getenv("PORT")
	if port == "" { port = "8080" }
	log.Printf("AI Builder Online on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot { return }
	msg := strings.ToLower(m.Content)

	if isWebsiteRequest(msg) {
		s.ChannelTyping(m.ChannelID)
		html := engine.GenerateWebsite(m.Content)
		id := database.SaveGeneratedSite(m.Content, html) // Возвращает ID
		
		domain := os.Getenv("DOMAIN")
		s.ChannelMessageSend(m.ChannelID, "✨ Сайт готов: " + domain + "/site/" + id)
	} else {
		// Просто отвечаем на вопросы, используя поиск
		info := engine.SearchInternet(m.Content)
		s.ChannelMessageSend(m.ChannelID, "🧠 Анализ сети: " + info)
	}
}

func isWebsiteRequest(msg string) bool {
	triggers := []string{"создай сайт", "сделай сайт", "website", "build", "create", "лендинг"}
	for _, t := range triggers {
		if strings.Contains(msg, t) { return true }
	}
	return false
}
