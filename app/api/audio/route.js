
import {
  createAudioSystem
}
from "../../../engine/audio/audioEngine"

export async function GET(){

  return Response.json(

    createAudioSystem()
  )
}
