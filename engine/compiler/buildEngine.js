
export async function buildProject(
  projectName
){

  return {

    success:true,

    buildPath:

    "/runtime/builds/" +
    projectName
  }
}
