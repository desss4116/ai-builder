package main

import (
	"ai-builder/analysis"
	"ai-builder/engine"
	"ai-builder/internet"
	"ai-builder/parser"
	"ai-builder/storage"

	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

var sites = map[string]string{}

func main() {

	token := os.Getenv("DISCORD_TOKEN")

	if token == "" {
		log.Fatal("DISCORD_TOKEN missing")
	}

	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatal(err)
	}

	dg.Identify.Intents =
		discordgo.IntentsGuildMessages |
			discordgo.IntentsMessageContent

	dg.AddHandler(messageCreate)

	err = dg.Open()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Bot online")

	http.HandleFunc("/site/", serveSite)

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func messageCreate(
	s *discordgo.Session,
	m *discordgo.MessageCreate,
) {

	if m.Author.Bot {
		return
	}

	content := strings.TrimSpace(m.Content)

	// WEBSITE REQUEST
	if parser.IsWebsiteRequest(content) {

		go handleWebsiteRequest(s, m, content)

		return
	}

	// NORMAL CHAT
	go handleChatRequest(s, m, content)
}

func handleChatRequest(
	s *discordgo.Session,
	m *discordgo.MessageCreate,
	query string,
) {

	html, err := internet.SearchInternet(query)
	if err != nil {

		s.ChannelMessageSend(
			m.ChannelID,
			"Ошибка поиска в интернете.",
		)

		return
	}

	results := internet.ExtractTitles(html)

	response := parser.BuildChatResponse(
		query,
		results,
	)

	s.ChannelMessageSend(
		m.ChannelID,
		response,
	)
}

func handleWebsiteRequest(
	s *discordgo.Session,
	m *discordgo.MessageCreate,
	query string,
) {

	s.ChannelMessageSend(
		m.ChannelID,
		"🌐 Анализирую интернет и создаю сайт...",
	)

	// INTERNET SEARCH
	searchHTML, _ := internet.SearchInternet(query)

	// ANALYSIS
	style := analysis.DetectStyle(searchHTML)

	_ = style

	// GENERATE SITE
	siteHTML := engine.GenerateWebsite(query)

	// SITE ID
	siteID := fmt.Sprintf(
		"%d",
		time.Now().UnixNano(),
	)

	sites[siteID] = siteHTML

	// SAVE
	storage.SaveSite(storage.Site{
		ID:    siteID,
		HTML:  siteHTML,
		Title: query,
	})

	// DOMAIN
	domain := os.Getenv("DOMAIN")

	if domain == "" {
		domain = "http://localhost:8080"
	}

	link := fmt.Sprintf(
		"%s/site/%s",
		domain,
		siteID,
	)

	s.ChannelMessageSend(
		m.ChannelID,
		"✅ Сайт создан:\n"+link,
	)
}

func serveSite(
	w http.ResponseWriter,
	r *http.Request,
) {

	id := strings.TrimPrefix(
		r.URL.Path,
		"/site/",
	)

	html, ok := sites[id]

	if !ok {

		http.NotFound(w, r)

		return
	}

	w.Header().Set(
		"Content-Type",
		"text/html",
	)

	fmt.Fprint(w, html)
}
