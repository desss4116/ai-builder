
import fs from "fs"

import path from "path"

export async function createProject(
  name
){

  const target =

  path.join(
    process.cwd(),
    "runtime/projects",
    name
  )

  fs.mkdirSync(
    target,
    {
      recursive:true
    }
  )

  return {

    success:true,

    target
  }
}
