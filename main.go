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

		s.ChannelMessageSend(
			m.ChannelID,
			"🧠 Анализирую запрос...",
		)

		results :=
			internet.Search(
				m.Content,
			)

		answer :=
			brain.Synthesize(
				m.Content,
				results,
			)

		msg :=
			"🌐 Ответ:\n\n" +
				answer

		s.ChannelMessageSend(
			m.ChannelID,
			msg,
		)
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
