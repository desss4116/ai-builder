
import {
  generateMaze
}
from "../../../engine/procedural/proceduralEngine"

export async function GET(){

  return Response.json({

    maze:
    generateMaze()
  })
}
