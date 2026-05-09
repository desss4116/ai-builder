
export function analyzePrompt(
  prompt
){

  const lower =
  prompt.toLowerCase()

  let type =
  "website"

  let cinematic =
  false

  let game =
  false

  let technologies = [

    "Next.js",

    "React",

    "Tailwind",

    "Framer Motion"
  ]

  if(
    lower.includes("game")
  ){

    game = true

    type = "game"

    technologies.push(
      "Three.js"
    )

    technologies.push(
      "PixiJS"
    )
  }

  if(
    lower.includes("maze")
  ){

    cinematic = true

    type = "maze-runner"

    technologies.push(
      "Procedural Generation"
    )
  }

  if(
    lower.includes("cinematic")
  ){

    cinematic = true
  }

  return {

    success:true,

    type,

    game,

    cinematic,

    technologies,

    confidence:0.98
  }
}
