
import {
  createScene
}
from "../../../engine/scenes/sceneEngine"

export async function GET(){

  return Response.json(

    createScene()
  )
}
