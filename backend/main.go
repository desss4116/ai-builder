
package main

import (
  "ai-builder/backend/discord"
  "log"
  "os"
)

func main(){

  token := os.Getenv("DISCORD_TOKEN")

  if token == ""{
    log.Fatal("DISCORD_TOKEN missing")
  }

  err := discord.StartBot(token)

  if err != nil{
    log.Fatal(err)
  }

  log.Println("AI Builder Backend Online")

  select{}
}
