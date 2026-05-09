package main

import (
	"ai-builder/brain"
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

	session, err := discordgo.New(
		"Bot " + token,
	)

	if err != nil {
		log.Fatal(err)
	}

	session.AddHandler(messageCreate)

	session.Identify.Intents =
		discordgo.IntentsGuildMessages |
			discordgo.IntentsDirectMessages |
			discordgo.IntentsMessageContent

	err = session.Open()

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

	query := strings.TrimSpace(
		m.Content,
	)

	if query == "" {
		return
	}

	s.ChannelMessageSend(
		m.ChannelID,
		"🧠 Анализирую запрос...",
	)

	// SEARCH
	results := internet.Search(query)

	// AI SYNTHESIS
	answer := brain.Synthesize(
		query,
		results,
	)

	if len(answer) > 1900 {
		answer = answer[:1900]
	}

	_, err := s.ChannelMessageSend(
		m.ChannelID,
		answer,
	)

	if err != nil {
		log.Println(err)
	}
}
