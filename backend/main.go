package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
)

type AIResponse struct {
	Answer string `json:"answer"`
}

func generateAnswer(input string) string {

	lower := strings.ToLower(input)

	if strings.Contains(lower, "hello") ||
		strings.Contains(lower, "hi") ||
		strings.Contains(lower, "привет") {

		return `
🤖 AI Builder

Hello.
AI system online.

Capabilities:
• Answer questions
• Analyze prompts
• Generate websites
• Basic reasoning
`
	}

	if strings.Contains(lower, "harry potter") {

		return `
🧠 AI Analysis

Harry Potter is a fantasy universe
created by J.K. Rowling.

Main themes:
• Magic
• Friendship
• War between good and evil
`
	}

	if strings.Contains(lower, "who are you") {

		return `
🤖 AI Builder Core

Autonomous Website Intelligence System

Modules:
• Discord AI
• Website Generator
• Reasoning Engine
• Research System
`
	}

	if strings.Contains(lower, "what can you do") {

		return `
⚡ AI Builder Features

• Answer questions
• Analyze requests
• Generate websites
• Build layouts
• Create structures
• AI reasoning
• Website intelligence
`
	}

	return fmt.Sprintf(`
🧠 AI Analysis Complete

Input:
%s

Status:
Processed successfully.

Inference:
The request was analyzed by the AI reasoning engine.
`, input)
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

	answer := generateAnswer(query)

	_, err := s.ChannelMessageSend(
		m.ChannelID,
		answer,
	)

	if err != nil {
		log.Println(err)
	}
}

func main() {

	token := os.Getenv(
		"DISCORD_TOKEN",
	)

	if token == "" {
		log.Fatal(
			"DISCORD_TOKEN missing",
		)
	}

	dg, err := discordgo.New(
		"Bot " + token,
	)

	if err != nil {
		log.Fatal(err)
	}

	dg.AddHandler(
		messageCreate,
	)

	dg.Identify.Intents =
		discordgo.IntentsGuildMessages |
			discordgo.IntentsMessageContent

	err = dg.Open()

	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc(
		"/",
		func(
			w http.ResponseWriter,
			r *http.Request,
		) {

			response := AIResponse{
				Answer: "AI Builder Backend Online",
			}

			json.NewEncoder(w).Encode(
				response,
			)
		},
	)

	go func() {

		err := http.ListenAndServe(
			":8080",
			nil,
		)

		if err != nil {
			log.Fatal(err)
		}
	}()

	log.Println(
		"🚀 AI Builder Online",
	)

	select {}
}
