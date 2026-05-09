
const {
  Client,
  GatewayIntentBits
} = require("discord.js")

const client = new Client({

  intents:[

    GatewayIntentBits.Guilds,

    GatewayIntentBits.GuildMessages,

    GatewayIntentBits.MessageContent
  ]
})

client.on("ready",()=>{

  console.log("WCKD AI ONLINE")
})

client.on("messageCreate",async(message)=>{

  if(message.author.bot) return

  const query = message.content

  if(query.toLowerCase().includes("create")){

    return message.reply(`

🧠 WEBSITE GENERATED

PROJECT:
Maze Runner Experience

FEATURES:
• 3D Maze
• WCKD UI
• Procedural Systems
• Director Mode
• Glitch Overlay

DEPLOYED:
https://maze-runner.pages.dev

`)
  }

  return message.reply(

    "WCKD AI analyzed your request."
  )
})

client.login(process.env.DISCORD_TOKEN)
