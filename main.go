package main

import (
	"ai-builder/brain"
	"ai-builder/internet"
	"fmt"
	"log"
	"net/http"
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

		fmt.Println("DISCORD_TOKEN not found")

		return
	}

	/*
	   CREATE DISCORD SESSION
	*/

	dg, err := discordgo.New(
		"Bot " + token,
	)

	if err != nil {

		fmt.Println(err)

		return
	}

	/*
	   MESSAGE HANDLER
	*/

	dg.AddHandler(func(
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

		message :=
			strings.TrimSpace(
				m.Content,
			)

		if message == "" {
			return
		}

		/*
		   REAL INTERNET SEARCH
		*/

		searchResult :=
			internet.LiveSearch(
				message,
			)

		/*
		   EMPTY RESULT
		*/

		if strings.TrimSpace(
			searchResult,
		) == "" {

			searchResult =
				"❌ Информация не найдена."
		}

		/*
		   FORMAT AI ANSWER
		*/

		finalAnswer :=
			brain.FormatAnswer(
				message,
				searchResult,
			)

		/*
		   SEND MESSAGE
		*/

		_, err := s.ChannelMessageSend(
			m.ChannelID,
			finalAnswer,
		)

		if err != nil {

			fmt.Println(err)
		}
	})

	/*
	   DISCORD INTENTS
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

		fmt.Println(err)

		return
	}

	fmt.Println(
		"🚀 AI Builder Online",
	)

	/*
	   OPTIONAL WEB SERVER
	*/

	http.HandleFunc("/", func(
		w http.ResponseWriter,
		r *http.Request,
	) {

		fmt.Fprint(
			w,
			"AI Builder Running",
		)
	})

	go func() {

		log.Fatal(
			http.ListenAndServe(
				":8080",
				nil,
			),
		)
	}()

	/*
	   KEEP ALIVE
	*/

	select {}
}
