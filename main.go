package main

import (
	"ai-builder/brain"
	"ai-builder/generator"
	"ai-builder/internet"
	"ai-builder/summarizer"
	"log"
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func main() {

	token := os.Getenv("DISCORD_TOKEN")

	if token == "" {
		log.Fatal("DISCORD_TOKEN missing")
	}

	dg, err := discordgo.New("Bot " + token)

	if err != nil {
		log.Fatal(err)
	}

	dg.AddHandler(messageCreate)

	dg.Identify.Intents =
		discordgo.IntentsGuildMessages |
			discordgo.IntentsDirectMessages |
			discordgo.IntentsMessageContent

	err = dg.Open()

	if err != nil {
		log.Fatal(err)
	}

	log.Println("🚀 AI Builder Online")

	select {}
}

func messageCreate(
	s *discordgo.Session,
	m *discordgo.MessageCreate,
) {

	/*
	   IGNORE BOTS
	*/

	if m.Author.Bot {
		return
	}

	/*
	   CLEAN MESSAGE
	*/

	message := strings.TrimSpace(m.Content)

	if message == "" {
		return
	}

	/*
	   DETECT INTENT
	*/

	intent := brain.DetectIntent(message)

	/*
	   WEBSITE GENERATION
	*/

	if intent == "website" {

		project :=
			generator.GenerateNextProject(message)

		s.ChannelMessageSend(
			m.ChannelID,
			"🚀 NEXT.JS PROJECT GENERATED\n\n"+project,
		)

		return
	}

	/*
	   REAL GOOGLE SEARCH
	*/

	searchResult :=
		internet.GoogleSearch(message)

	/*
	   SCRAPING / CLEANING
	*/

	cleanContent :=
		internet.ScrapeContent(searchResult)

	/*
	   AI SUMMARY
	*/

	summary :=
		summarizer.Summarize(cleanContent)

	/*
	   REASONING
	*/

	answer :=
		brain.Think(
			message,
			summary,
		)

	/*
	   FALLBACK
	*/

	if strings.TrimSpace(answer) == "" {

		answer = `
❌ Не удалось получить ответ.

Попробуй:
- уточнить запрос
- написать конкретнее
- использовать английский язык
`
	}

	/*
	   SEND MESSAGE
	*/

	_, err = s.ChannelMessageSend(
		m.ChannelID,
		answer,
	)

	if err != nil {
		log.Println(err)
	}
}
