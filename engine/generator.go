package engine

import (
	"fmt"
)

func GenerateWebsite(title string) string {

	html := fmt.Sprintf(`
<!DOCTYPE html>
<html>
<head>

<meta charset="UTF-8">
<meta name="viewport" content="width=device-width,initial-scale=1.0">

<title>%s</title>

%s
%s

</head>

<body>

<section class="hero fade-up">

<div class="container">

<h1>%s</h1>

<p>
Premium AI Generated Website
</p>

<a href="#start" class="btn">
Get Started
</a>

</div>

</section>

<section class="container">

<div class="card fade-up">
	<h2>Modern UI</h2>
	<p>Internet-inspired intelligent generation.</p>
</div>

<div class="card fade-up">
	<h2>Adaptive Design</h2>
	<p>Dynamic composition engine.</p>
</div>

</section>

</body>
</html>
`,
		title,
		GlobalStyles(),
		Animations(),
		title,
	)

	return html
}
