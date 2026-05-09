
export async function deployProject(
  projectName
){

  return {

    success:true,

    url:
    "https://" +
    projectName +
    ".pages.dev"
  }
}
