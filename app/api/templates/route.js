
export async function GET(){

  return Response.json({

    success:true,

    templates:[

      "maze-runner",

      "cyberpunk",

      "cinematic-saas",

      "aaa-game"
    ]
  })
}
