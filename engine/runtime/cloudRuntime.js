
export class CloudRuntime{

  constructor(){

    this.providers = [

      "Cloudflare",

      "Railway",

      "Vercel"
    ]
  }

  deploy(project){

    return {

      success:true,

      provider:
      this.providers[0],

      url:
      "https://" +
      project +
      ".pages.dev"
    }
  }
}
