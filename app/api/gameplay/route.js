
export async function GET(){

  return Response.json({

    success:true,

    gameplay:{

      combat:true,

      stealth:true,

      procedural:true,

      ai:true
    }
  })
}
