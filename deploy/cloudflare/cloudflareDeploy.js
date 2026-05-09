
export async function deployToCloudflare(
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
