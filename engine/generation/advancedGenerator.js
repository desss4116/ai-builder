
import {
  analyzePrompt
}
from "../reasoning/analyzer"

import {
  generateWorld
}
from "../worlds/worldGenerator"

import {
  generateStory
}
from "../storytelling/storyEngine"

export async function advancedGenerate(
  prompt
){

  const analysis =
  analyzePrompt(prompt)

  const world =
  generateWorld()

  const story =
  generateStory(
    analysis.type
  )

  return {

    success:true,

    analysis,

    world,

    story,

    generatedAt:
    Date.now()
  }
}
