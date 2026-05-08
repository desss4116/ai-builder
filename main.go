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
		log.Fatal("DISCORD_TOKEN not found")
	}

	dg, err := discordgo.New("Bot " + token)

	if err != nil {
		log.Fatal(err)
	}

	dg.AddHandler(messageCreate)

	dg.Identify.Intents = discordgo.IntentsGuildMessages |
		discordgo.IntentsDirectMessages |
		discordgo.IntentsMessageContent

	err = dg.Open()

	if err != nil {
		log.Fatal(err)
	}

	log.Println("🚀 AI Builder System Online")

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
	   WEBSITE MODE
	*/

	if intent == "website" {

		project := generator.GenerateNextProject(message)

		s.ChannelMessageSend(
			m.ChannelID,
			"🚀 NEXT.JS PROJECT GENERATED\n\n"+project,
		)

		return
	}

	/*
	   INTERNET SEARCH
	*/

	result := internet.SearchWeb(message)

	/*
	   AI ANSWER ENGINE
	*/

	finalAnswer := brain.GenerateAnswer(
		message,
		result,
	)

	/*
	   SEND RESPONSE
	*/

	_, err := s.ChannelMessageSend(
		m.ChannelID,
		finalAnswer,
	)

	if err != nil {
		log.Println(err)
	}
}
