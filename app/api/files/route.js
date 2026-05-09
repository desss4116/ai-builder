
import fs from "fs"

import path from "path"

export async function GET(){

  const runtime =

  path.join(
    process.cwd(),
    "runtime/projects"
  )

  const files =

  fs.readdirSync(runtime)

  return Response.json({

    success:true,

    files
  })
}
