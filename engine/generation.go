package engine

import (
	"fmt"
	"strings"
)

func GenerateWebsite(prompt string) string {

	style := "from-purple-600 to-indigo-700"

	lower := strings.ToLower(prompt)

	if strings.Contains(lower, "luxury") {
		style = "from-yellow-400 to-yellow-700"
	}

	if strings.Contains(lower, "nike") {
		style = "from-red-500 to-orange-500"
	}

	return fmt.Sprintf(`
<!DOCTYPE html>
<html>

<head>

<meta charset="UTF-8">
<meta name="viewport" content="width=device-width,initial-scale=1.0">

%s
%s

<link rel="preconnect" href="https://fonts.googleapis.com">

<link href="https://fonts.googleapis.com/css2?family=Inter:wght@400;600;700;800&display=swap" rel="stylesheet">

<style>

body{
	font-family:Inter,sans-serif;
	background:#050505;
	color:white;
}

</style>

</head>

<body>

<div id="root"></div>

<script>

const e = React.createElement

function App(){

	return e(
		'div',
		{
			className:'min-h-screen bg-black text-white'
		},

		e(
			'section',
			{
				className:'px-8 py-32 text-center'
			},

			e(
				'h1',
				{
					className:'text-7xl font-extrabold leading-tight bg-gradient-to-r %s bg-clip-text text-transparent'
				},
				'%s'
			),

			e(
				'p',
				{
					className:'mt-8 text-zinc-400 text-2xl max-w-3xl mx-auto'
				},
				'Next generation AI website experience.'
			),

			e(
				'div',
				{
					className:'mt-12 flex justify-center gap-6 flex-wrap'
				},

				e(
					'a',
					{
						href:'#',
						className:'px-8 py-4 rounded-2xl bg-white text-black font-bold'
					},
					'Launch'
				),

				e(
					'a',
					{
						href:'#',
						className:'px-8 py-4 rounded-2xl border border-white/20'
					},
					'Preview'
				),
			),
		),
	)
}

ReactDOM.createRoot(
	document.getElementById('root')
).render(
	e(App)
)

</script>

</body>
</html>
`,
		ReactRoot(),
		Tailwind(),
		style,
		prompt,
	)
}
