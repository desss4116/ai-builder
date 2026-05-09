
export async function POST(req) {

  const body = await req.json()

  return Response.json({
    success:true,
    prompt:body.prompt || null,
    message:"AI Builder Generator Online"
  })
}
