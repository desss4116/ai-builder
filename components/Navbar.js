
export default function Navbar() {

  return (
    <nav
      style={{
        width:"100%",
        padding:"20px 40px",
        display:"flex",
        justifyContent:"space-between",
        alignItems:"center",
        borderBottom:"1px solid rgba(255,255,255,0.1)"
      }}
    >

      <h2>AI Builder</h2>

      <button
        style={{
          background:"#2563eb",
          color:"white",
          border:"none",
          padding:"12px 20px",
          borderRadius:"10px",
          cursor:"pointer"
        }}
      >
        Dashboard
      </button>

    </nav>
  )
}
