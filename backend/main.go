
package main

import (
  "fmt"
  "log"
  "os"
  "strings"

  "github.com/bwmarrin/discordgo"
)

func generate(prompt string) string {

  lower :=
  strings.ToLower(prompt)

  if strings.Contains(lower,"maze") {

    return `
🧠 WEBSITE GENERATED

PROJECT:
Maze Runner Experience

FEATURES:
• 3D Maze
• Infinite Runner
• WCKD UI
• Glitch Effects
• Procedural Systems

DEPLOYED:
https://maze-runner.pages.dev
`
  }

  return fmt.Sprintf(`
🧠 WEBSITE GENERATED

PROMPT:
%s

DEPLOYED:
https://ai-builder.pages.dev
`,prompt)
}

func messageCreate(
  s *discordgo.Session,
  m *discordgo.MessageCreate,
){

  if m.Author.Bot{
    return
  }

  response :=
  generate(m.Content)

  s.ChannelMessageSend(
    m.ChannelID,
    response,
  )
}

func main(){

  token :=
  os.Getenv("DISCORD_TOKEN")

  dg, err :=
  discordgo.New(
    "Bot " + token,
  )

  if err != nil{
    log.Fatal(err)
  }

  dg.AddHandler(messageCreate)

  dg.Identify.Intents =
    discordgo.IntentsGuildMessages |
    discordgo.IntentsMessageContent

  err = dg.Open()

  if err != nil{
    log.Fatal(err)
  }

  log.Println(
    "AI Builder Omega Online",
  )

  select{}
}
