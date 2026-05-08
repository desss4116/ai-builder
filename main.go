package main

import (
	"ai-builder/brain"
	"ai-builder/internet"
	"ai-builder/memory"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func main() {

	token := os.Getenv(
		"DISCORD_TOKEN",
	)

	if token == "" {

		fmt.Println(
			"DISCORD_TOKEN missing",
		)

		return
	}

	dg, err := discordgo.New(
		"Bot " + token,
	)

	if err != nil {

		fmt.Println(err)

		return
	}

	dg.AddHandler(func(
		s *discordgo.Session,
		m *discordgo.MessageCreate,
	) {

		/*
		   IGNORE BOT
		*/

		if m.Author.Bot {
			return
		}

		query :=
			strings.TrimSpace(
				m.Content,
			)

		if query == "" {
			return
		}

		/*
		   SAVE MEMORY
		*/

		memory.Save(query)

		/*
		   SEARCH
		*/

		search :=
			internet.LiveSearch(query)

		/*
		   SCRAPE
		*/

		scraped :=
			internet.ScrapePage(
				"https://duckduckgo.com",
			)

		/*
		   REASONING
		*/

		answer :=
			brain.Reason(
				query,
				search,
				scraped,
			)

		/*
		   SEND
		*/

		s.ChannelMessageSend(
			m.ChannelID,
			answer,
		)
	})

	dg.Identify.Intents =
		discordgo.IntentsGuildMessages |
			discordgo.IntentsDirectMessages |
			discordgo.IntentsMessageContent

	err = dg.Open()

	if err != nil {

		fmt.Println(err)

		return
	}

	fmt.Println(
		"🚀 AI Builder Online",
	)

	http.HandleFunc("/", func(
		w http.ResponseWriter,
		r *http.Request,
	) {

		fmt.Fprint(
			w,
			"AI Builder Running",
		)
	})

	http.ListenAndServe(
		":8080",
		nil,
	)
}
