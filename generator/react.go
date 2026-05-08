package generator

func GenerateReactPage(title string) string {

	return `
export default function Page() {
	return (
		<div className="min-h-screen bg-black text-white">

			<section className="max-w-7xl mx-auto px-6 py-32">

				<h1 className="text-8xl font-bold tracking-tight">
					` + title + `
				</h1>

				<p className="text-zinc-400 text-xl mt-6">
					Premium AI generated experience.
				</p>

			</section>

		</div>
	)
}
`
}
