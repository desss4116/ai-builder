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
	if err != nil { log.Fatal(err) }

	dg.AddHandler(messageHandler)
	dg.Identify.Intents = discordgo.IntentsGuildMessages | discordgo.IntentsMessageContent
	if err = dg.Open(); err != nil { log.Fatal(err) }

	// Роутинг
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { http.ServeFile(w, r, "index.html") })
	http.HandleFunc("/editor/", func(w http.ResponseWriter, r *http.Request) { http.ServeFile(w, r, "editor.html") })
	http.HandleFunc("/site/", siteServeHandler)

	port := os.Getenv("PORT")
	if port == "" { port = "8080" }
	log.Printf("🚀 AI.OS Core Online | Port: %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot { return }
	input := strings.TrimSpace(m.Content)

	if isWebsiteRequest(strings.ToLower(input)) {
		s.ChannelTyping(m.ChannelID)
		html := engine.GenerateWebsite(input)
		id := fmt.Sprintf("%d", time.Now().UnixNano())
		database.SaveGeneratedSite(id, input, html)

		domain := os.Getenv("DOMAIN")
		msg := fmt.Sprintf("✨ **Website Intelligence Engine**\n\n✅ Продукт готов: %s/site/%s\n🎨 Редактор: %s/editor/%s", domain, id, domain, id)
		s.ChannelMessageSend(m.ChannelID, msg)
	} else {
		s.ChannelTyping(m.ChannelID)
		// Умный поиск вместо простых ответов
		data := engine.SearchInternet(input)
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("🧠 **AI Intelligence Layer**\n\n%s", data))
	}
}

func siteServeHandler(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/site/")
	html := database.GetSiteHTML(id) 
	if html == "" { http.NotFound(w, r); return }
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, html)
}

func isWebsiteRequest(msg string) bool {
	triggers := []string{"сайт", "website", "build", "landing", "создай", "сделай"}
	for _, t := range triggers { if strings.Contains(msg, t) { return true } }
	return false
}
