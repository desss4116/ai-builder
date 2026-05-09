
console.log(
  "AI Builder Discord Runtime Online"
)

async function handlePrompt(
  prompt
){

  console.log(
    "PROCESSING:",
    prompt
  )

  return {

    success:true,

    url:
    "https://generated.pages.dev"
  }
}

module.exports = {
  handlePrompt
}
