
import fs from "fs"
import path from "path"

export async function generateProjectFiles(
  projectName,
  prompt
){

  const base =
  path.join(
    process.cwd(),
    "runtime/projects",
    projectName
  )

  fs.mkdirSync(base,{
    recursive:true
  })

  fs.writeFileSync(

    path.join(base,"README.md"),

`
# Generated Website

Prompt:

${prompt}
`
  )

  return {
    success:true,
    path:base
  }
}
