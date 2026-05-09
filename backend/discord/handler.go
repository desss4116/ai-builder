
package discord

import (
  "ai-builder/backend/engine"
  "github.com/bwmarrin/discordgo"
)

func MessageHandler(
  s *discordgo.Session,
  m *discordgo.MessageCreate,
){

  if m.Author.Bot{
    return
  }

  response := engine.ProcessQuery(m.Content)

  s.ChannelMessageSend(
    m.ChannelID,
    response,
  )
}
