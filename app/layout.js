
import "./globals.css"

export const metadata = {
  title:"AI Builder",
  description:"Autonomous Website Intelligence System"
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
