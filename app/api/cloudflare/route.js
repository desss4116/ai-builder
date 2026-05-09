
export async function POST(req){

  const body =
  await req.json()

  return Response.json({

    success:true,

    deployed:true,

    url:
    "https://" +
    body.project +
    ".pages.dev"
  })
}
