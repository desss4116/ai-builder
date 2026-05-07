package components

func Hero(title string) string {

	return `
<section class="hero">

<div class="container">

<div class="hero-content fade-up">

<h1>` + title + `</h1>

<p>
Premium AI Generated Website
with modern UI,
animations and interactions.
</p>

<button
class="btn"
onclick="
document
.getElementById('features')
.scrollIntoView({
behavior:'smooth'
})
"
>
Explore
</button>

</div>

</div>

</section>
`
}
