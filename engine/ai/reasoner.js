
export async function reason(query){

  const lower = query.toLowerCase()

  let category = "general"

  if(lower.includes("maze")){

    category = "maze"
  }

  if(lower.includes("game")){

    category = "game"
  }

  if(lower.includes("website")){

    category = "website"
  }

  return {

    success:true,

    category,

    response:
      "WCKD AI analyzed your request and generated tactical response.",

    recommendations:[

      "Enable cinematic mode",

      "Activate procedural systems",

      "Increase immersion layer"
    ]
  }
}
