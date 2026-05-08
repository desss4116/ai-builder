package main

import (
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
	   DISCORD SESSION
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
		   IGNORE BOT
		*/

		if m.Author.Bot {
			return
		}

		message :=
			strings.TrimSpace(
				m.Content,
			)

		if message == "" {
			return
		}

		/*
		   INTERNET SEARCH
		*/

		searchResult :=
			internet.LiveSearch(message)

		/*
		   EMPTY RESULT
		*/

		if strings.TrimSpace(
			searchResult,
		) == "" {

			searchResult =
				"❌ Ничего не найдено."
		}

		/*
		   SEND RESPONSE
		*/

		_, err := s.ChannelMessageSend(
			m.ChannelID,
			"🌐 Ответ:\n\n"+searchResult,
		)

		if err != nil {

			fmt.Println(err)
		}
	})

	/*
	   OPEN DISCORD
	*/

	err = dg.Open()

	if err != nil {

		fmt.Println(err)

		return
	}

	fmt.Println(
		"AI Builder Discord Bot Started",
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
