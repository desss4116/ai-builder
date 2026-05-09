
import fs from "fs"

import path from "path"

export async function exportProject(
  projectName
){

  const zipPath =

  path.join(
    process.cwd(),
    "runtime/zips",
    projectName + ".zip"
  )

  fs.writeFileSync(
    zipPath,
    "ZIP_PLACEHOLDER"
  )

  return {

    success:true,

    zipPath
  }
}
