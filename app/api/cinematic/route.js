
export async function GET(){

  return Response.json({

    success:true,

    cinematic:{

      bloom:true,

      lensFlare:true,

      volumetricFog:true,

      chromaticAberration:true
    }
  })
}
