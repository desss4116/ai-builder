
import {
  createEnemyAI
}
from "../../../engine/enemies/enemyEngine"

export async function GET(){

  return Response.json(

    createEnemyAI()
  )
}
