
import "./globals.css"

export const metadata = {
  title:"AI Builder Omega",
  description:"Autonomous Website Factory"
}

export default function RootLayout({children}){

  return (
    <html lang="en">
      <body>
        {children}
      </body>
    </html>
  )
}
