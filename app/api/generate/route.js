
import {generateProject}
from "../../../engine/generation/generator"

export async function POST(req){

  const body = await req.json()

  const result =
    await generateProject(body.prompt)

  return Response.json(result)
}
