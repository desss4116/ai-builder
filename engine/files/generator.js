
import fs from "fs"

import path from "path"

export async function generateFiles(

  projectName,
  prompt

){

  const base =

  path.join(
    process.cwd(),
    "runtime/projects",
    projectName
  )

  fs.mkdirSync(
    base,
    {
      recursive:true
    }
  )

  fs.writeFileSync(

    path.join(
      base,
      "README.md"
    ),

`
# Generated Project

Prompt:

${prompt}
`
  )

  fs.writeFileSync(

    path.join(
      base,
      "project.json"
    ),

    JSON.stringify({

      projectName,

      prompt,

      createdAt:
      Date.now()

    },null,2)
  )

  return {

    success:true,

    path:base
  }
}
