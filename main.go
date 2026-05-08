package main

import (
	"ai-builder/brain"
	"ai-builder/generator"
	"ai-builder/internet"
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

	if m.Author.Bot {
		return
	}

	message := strings.TrimSpace(m.Content)

	if message == "" {
		return
	}

	intent := brain.DetectIntent(message)

	/*
	   WEBSITE GENERATION
	*/

	if intent == "website" {

		project := generator.GenerateNextProject(message)

		s.ChannelMessageSend(
			m.ChannelID,
			"🚀 NEXT.JS APP GENERATED\n\n"+project,
		)

		return
	}

	/*
	   INTERNET PIPELINE
	*/

	searchResult := internet.SearchWeb(message)

	scraped := internet.ScrapeContent(searchResult)

	summary := internet.Summarize(scraped)

	/*
	   AI REASONING
	*/

	finalAnswer := brain.Think(
		message,
		summary,
	)

	/*
	   SEND
	*/

	s.ChannelMessageSend(
		m.ChannelID,
		finalAnswer,
	)
}
