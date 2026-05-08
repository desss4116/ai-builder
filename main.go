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

	/*
	   DISCORD TOKEN
	*/

	token := os.Getenv("DISCORD_TOKEN")

	if token == "" {
		log.Fatal("DISCORD_TOKEN missing")
	}

	/*
	   CREATE DISCORD SESSION
	*/

	dg, err := discordgo.New("Bot " + token)

	if err != nil {
		log.Fatal(err)
	}

	/*
	   ADD MESSAGE HANDLER
	*/

	dg.AddHandler(messageCreate)

	/*
	   BOT INTENTS
	*/

	dg.Identify.Intents =
		discordgo.IntentsGuildMessages |
			discordgo.IntentsDirectMessages |
			discordgo.IntentsMessageContent

	/*
	   OPEN CONNECTION
	*/

	err = dg.Open()

	if err != nil {
		log.Fatal(err)
	}

	log.Println("🚀 AI Builder Online")

	/*
	   KEEP ALIVE
	*/

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

		_, sendErr := s.ChannelMessageSend(
			m.ChannelID,
			"🚀 NEXT.JS PROJECT GENERATED\n\n"+project,
		)

		if sendErr != nil {
			log.Println(sendErr)
		}

		return
	}

	/*
	   REAL GOOGLE SEARCH
	*/

	searchResult :=
	internet.LiveSearch(message)

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

	_, sendErr := s.ChannelMessageSend(
		m.ChannelID,
		answer,
	)

	if sendErr != nil {
		log.Println(sendErr)
	}
}
