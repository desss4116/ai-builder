
import {reason}
from "../../../engine/ai/reasoner"

export async function POST(req){

  const body = await req.json()

  const result = await reason(body.message)

  return Response.json(result)
}
