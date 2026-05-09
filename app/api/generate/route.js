
import {v4 as uuid}
from "uuid"

import {
  analyzePrompt
}
from "../../../engine/reasoning/analyzer"

import {
  generateFiles
}
from "../../../engine/files/generator"

import {
  deployProject
}
from "../../../engine/deployment/deployer"

import {
  createProject
}
from "../../../engine/projects/projectEngine"

export async function POST(req){

  try{

    const body =
    await req.json()

    const prompt =
    body.prompt

    const analysis =
    analyzePrompt(prompt)

    const projectName =

    "project-" +
    uuid()

    await createProject(
      projectName
    )

    await generateFiles(
      projectName,
      prompt
    )

    const deployment =
    await deployProject(
      projectName
    )

    return Response.json({

      success:true,

      projectName,

      analysis,

      deployment,

      prompt
    })

  }catch(error){

    return Response.json({

      success:false,

      error:error.message
    })
  }
}
