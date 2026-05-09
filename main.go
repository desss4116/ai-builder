package main

import (
	"ai-builder/brain"
	"ai-builder/internet"
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
)

func main() {

	token :=
		os.Getenv(
			"DISCORD_TOKEN",
		)

	dg, err :=
		discordgo.New(
			"Bot " + token,
		)

	if err != nil {
		panic(err)
	}

	dg.AddHandler(func(
		s *discordgo.Session,
		m *discordgo.MessageCreate,
	) {

		if m.Author.Bot {
			return
		}

		go handleMessage(s, m)
	})

	err = dg.Open()

	if err != nil {
		panic(err)
	}

	fmt.Println(
		"AI Builder Online",
	)

	select {}
}

func handleMessage(
	s *discordgo.Session,
	m *discordgo.MessageCreate,
) {

	s.ChannelMessageSend(
		m.ChannelID,
		"🧠 Анализирую запрос...",
	)

	results :=
		internet.Search(
			m.Content,
		)

	if len(results) == 0 {

		s.ChannelMessageSend(
			m.ChannelID,
			"❌ Ничего не найдено.",
		)

		return
	}

	answer :=
		brain.Synthesize(
			m.Content,
			results,
		)

	if answer == "" {
		answer = "❌ Не удалось сформировать ответ."
	}

	msg :=
		"🌐 Ответ:\n\n" +
			answer

	if len(msg) > 1900 {
		msg = msg[:1900]
	}

	s.ChannelMessageSend(
		m.ChannelID,
		msg,
	)
}
