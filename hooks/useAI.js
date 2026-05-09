
export function useAI() {

  async function generate(prompt) {

    const res = await fetch("/api/generate",{
      method:"POST",
      headers:{
        "Content-Type":"application/json"
      },
      body:JSON.stringify({
        prompt
      })
    })

    return res.json()
  }

  return {
    generate
  }
}
