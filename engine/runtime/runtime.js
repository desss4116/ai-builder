
import fs from "fs"

import path from "path"

export async function bootRuntime(
  projectName
){

  const runtimePath =

  path.join(
    process.cwd(),
    "runtime/projects",
    projectName
  )

  const exists =
  fs.existsSync(runtimePath)

  if(!exists){

    return {

      success:false,

      error:
      "Runtime project not found"
    }
  }

  return {

    success:true,

    runtimePath,

    status:"ONLINE"
  }
}
