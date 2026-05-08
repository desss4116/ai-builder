package main

import (
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
		log.Fatal("DISCORD_TOKEN not found")
	}

	session, err := discordgo.New("Bot " + token)

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

	// Ignore bots
	if m.Author.Bot {
		return
	}

	query := strings.TrimSpace(m.Content)

	if query == "" {
		return
	}

	// Thinking message
	s.ChannelMessageSend(
		m.ChannelID,
		"🧠 Анализирую запрос...",
	)

	// Search
	results := internet.Search(query)

	if len(results) == 0 {

		s.ChannelMessageSend(
			m.ChannelID,
			"❌ Ничего не найдено.",
		)

		return
	}

	// Multi-source extraction
	var combined string

	maxSources := 3

	if len(results) < maxSources {
		maxSources = len(results)
	}

	for i := 0; i < maxSources; i++ {

		url := results[i].URL

		html := internet.FetchPage(url)

		if html == "" {
			continue
		}

		clean := internet.ExtractCleanText(html)

		if clean == "" {
			continue
		}

		combined += clean + "\n\n"
	}

	// No extracted content
	if combined == "" {

		s.ChannelMessageSend(
			m.ChannelID,
			"❌ Не удалось извлечь информацию.",
		)

		return
	}

	// Build AI answer
	answer := summarizer.BuildAnswer(combined)

	// Discord limit
	if len(answer) > 1900 {
		answer = answer[:1900]
	}

	// Send answer
	s.ChannelMessageSend(
		m.ChannelID,
		answer,
	)
}
