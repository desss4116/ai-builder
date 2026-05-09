
import {
  generateShader
}
from "../../../engine/shaders/shaderEngine"

export async function GET(){

  return Response.json({

    shader:
    generateShader()
  })
}
