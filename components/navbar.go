package components

func Navbar() string {

	return `
<nav class="navbar">

<div class="container nav-inner">

<div class="logo">
AI Builder
</div>

<div
class="nav-links"
id="navLinks"
>

<a href="#features">
Features
</a>

<a href="#pricing">
Pricing
</a>

<a href="#faq">
FAQ
</a>

</div>

<div
class="mobile-menu"
id="menuBtn"
>
☰
</div>

</div>

</nav>
`
}
