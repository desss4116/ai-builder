
import {
  analyzePrompt
}
from "../../../engine/reasoning/analyzer"

export async function POST(req){

  const body =
  await req.json()

  return Response.json(

    analyzePrompt(
      body.prompt
    )
  )
}
