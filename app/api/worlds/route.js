
import {
  generateWorld
}
from "../../../engine/worlds/worldGenerator"

export async function GET(){

  return Response.json(

    generateWorld()
  )
}
