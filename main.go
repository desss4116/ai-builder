package main

import (
	"ai-builder/analysis"
	"ai-builder/engine"
	"ai-builder/internet"
	"ai-builder/parser"
	"ai-builder/storage"

	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

type SiteData struct {
	ID   string `json:"id"`
	HTML string `json:"html"`
}

var sites = map[string]string{}

func main() {

	rand.Seed(time.Now().UnixNano())

	token := os.Getenv("DISCORD_TOKEN")

	if token == "" {
		log.Fatal("DISCORD_TOKEN missing")
	}

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	dg, err := discordgo.New(
		"Bot " + token,
	)

	if err != nil {
		log.Fatal(err)
	}

	dg.AddHandler(messageCreate)

	dg.Identify.Intents =
		discordgo.IntentsGuildMessages |
			discordgo.IntentsMessageContent

	err = dg.Open()

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Bot online")

	http.HandleFunc("/", homeHandler)

	http.HandleFunc("/site/", siteHandler)

	log.Fatal(
		http.ListenAndServe(":"+port, nil),
	)
}

func homeHandler(
	w http.ResponseWriter,
	r *http.Request,
) {

	w.Write([]byte(`
	<h1>AI Builder Online</h1>
	`))
}

func siteHandler(
	w http.ResponseWriter,
	r *http.Request,
) {

	id := strings.TrimPrefix(
		r.URL.Path,
		"/site/",
	)

	html, ok := sites[id]

	if !ok {

		w.WriteHeader(404)

		w.Write([]byte("site not found"))

		return
	}

	w.Header().Set(
		"Content-Type",
		"text/html",
	)

	w.Write([]byte(html))
}

func messageCreate(
	s *discordgo.Session,
	m *discordgo.MessageCreate,
) {

	if m.Author.Bot {
		return
	}

	query := strings.TrimSpace(
		m.Content,
	)

	if query == "" {
		return
	}

	intent := analysis.DetectIntent(query)

	/*
	   WEBSITE GENERATION
	*/

	if parser.IsWebsiteRequest(query) ||
		intent == "website" {

		html := engine.GenerateWebsite(query)

		id := generateID()

		sites[id] = html

		storage.SaveSite(
			storage.Site{
				ID:   id,
				HTML: html,
			},
		)

		domain := os.Getenv("DOMAIN")

		url :=
			domain +
				"/site/" +
				id

		embed := &discordgo.MessageEmbed{
			Title:       "🚀 Website Generated",
			Description: query,
			Color:       0x5865F2,
			Fields: []*discordgo.MessageEmbedField{
				{
					Name:   "Live Preview",
					Value:  url,
					Inline: false,
				},
			},
		}

		s.ChannelMessageSendEmbed(
			m.ChannelID,
			embed,
		)

		return
	}

	/*
	   INTERNET SEARCH
	*/

	html, err := internet.SearchInternet(
		query,
	)

	if err != nil {

		s.ChannelMessageSend(
			m.ChannelID,
			"Internet search failed.",
		)

		return
	}

	results :=
		internet.ExtractTitles(html)

	results =
		analysis.RankResults(results)

	answer :=
		internet.BuildAnswer(
			query,
			results,
		)

	answer =
		parser.CleanResponse(answer)

	s.ChannelMessageSend(
		m.ChannelID,
		answer,
	)
}

func generateID() string {

	const chars =
		"abcdefghijklmnopqrstuvwxyz0123456789"

	result := make([]byte, 12)

	for i := range result {

		result[i] =
			chars[rand.Intn(len(chars))]
	}

	return string(result)
}

func saveSitesFile() {

	var all []SiteData

	for id, html := range sites {

		all = append(
			all,
			SiteData{
				ID:   id,
				HTML: html,
			},
		)
	}

	data, _ := json.MarshalIndent(
		all,
		"",
		"  ",
	)

	os.WriteFile(
		"data/sites.json",
		data,
		0644,
	)
}

func init() {

	os.MkdirAll("data", 0755)

	file, err := os.ReadFile(
		"data/sites.json",
	)

	if err != nil {
		return
	}

	var all []SiteData

	json.Unmarshal(file, &all)

	for _, s := range all {
		sites[s.ID] = s.HTML
	}
}
