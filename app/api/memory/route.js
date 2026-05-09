
import {
  getMemory
}
from "../../../engine/memory/memoryEngine"

export async function GET(){

  return Response.json({

    success:true,

    memory:
    getMemory()
  })
}
