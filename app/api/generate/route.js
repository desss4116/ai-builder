
import {v4 as uuid} from "uuid"

import {
  generateProjectFiles
}
from "../../../engine/files/generator"

import {
  deployProject
}
from "../../../engine/deploy/deployer"

export async function POST(req){

  const body =
  await req.json()

  const prompt =
  body.prompt

  const projectName =
  "project-" +
  uuid()

  await generateProjectFiles(
    projectName,
    prompt
  )

  const deploy =
  await deployProject(
    projectName
  )

  return Response.json({

    success:true,

    project:projectName,

    url:deploy.url,

    prompt
  })
}
