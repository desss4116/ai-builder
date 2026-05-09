
import {
  generateStory
}
from "../../../engine/storytelling/storyEngine"

export async function GET(){

  return Response.json(

    generateStory(
      "cinematic"
    )
  )
}
