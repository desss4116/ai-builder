
package discord

import "github.com/bwmarrin/discordgo"

func StartBot(token string) error {

  session, err := discordgo.New("Bot " + token)

  if err != nil{
    return err
  }

  session.AddHandler(MessageHandler)

  session.Identify.Intents =
    discordgo.IntentsGuildMessages |
    discordgo.IntentsDirectMessages |
    discordgo.IntentsMessageContent

  return session.Open()
}
