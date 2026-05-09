package main

import (
	"ai-builder/brain"
	"ai-builder/internet"
	"ai-builder/memory"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func main() {

	token := os.Getenv("DISCORD_TOKEN")

	if token == "" {
		log.Fatal("DISCORD_TOKEN not found")
	}

	session, err :=
		discordgo.New(
			"Bot " + token,
		)

	if err != nil {
		log.Fatal(err)
	}

	session.Identify.Intents =
		discordgo.IntentsGuildMessages |
			discordgo.IntentsDirectMessages |
			discordgo.IntentsMessageContent

	session.AddHandler(messageCreate)

	err = session.Open()

	if err != nil {
		log.Fatal(err)
	}

	log.Println("AI Builder Online")

	stop := make(chan os.Signal, 1)

	signal.Notify(
		stop,
		syscall.SIGINT,
		syscall.SIGTERM,
		os.Interrupt,
	)

	<-stop

	session.Close()
}

func messageCreate(
	s *discordgo.Session,
	m *discordgo.MessageCreate,
) {

	if m.Author.Bot {
		return
	}

	query := m.Content

	if query == "" {
		return
	}

	_, _ = s.ChannelMessageSend(
		m.ChannelID,
		"🧠 Анализирую запрос...",
	)

	// INTERNET SEARCH

	results :=
		internet.LiveSearch(query)

	if len(results) == 0 {

		_, _ = s.ChannelMessageSend(
			m.ChannelID,
			"❌ Ничего не найдено.",
		)

		return
	}

	// AI SYNTHESIS

	answer :=
		brain.Synthesize(
			query,
			results,
		)

	// MEMORY

	memory.Save(answer)

	if len(answer) > 1900 {
		answer = answer[:1900]
	}

	_, sendErr :=
		s.ChannelMessageSend(
			m.ChannelID,
			answer,
		)

	if sendErr != nil {
		log.Println(sendErr)
	}
}
