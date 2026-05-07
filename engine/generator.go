package engine

import (
	"ai-builder/components"
	"fmt"
)

type Context struct {
	Prompt string
	Type string
}

func GenerateWebsite(
	ctx Context,
) string {

	return fmt.Sprintf(`

<!DOCTYPE html>

<html>

<head>

<meta charset="UTF-8">

<meta
name="viewport"
content="width=device-width, initial-scale=1.0"
/>

<title>%s</title>

%s

</head>

<body>

%s

%s

%s

%s

%s

%s

%s

%s

</body>

</html>

`,
		ctx.Prompt,

		GlobalStyles(),

		components.Navbar(),

		components.Hero(ctx.Prompt),

		components.Features(),

		components.Products(),

		components.Testimonials(),

		components.Pricing(),

		components.FAQ(),

		components.Footer(),

		components.Scripts(),
	)
}
