package components

func Scripts() string {

	return `
<script>

const menuBtn =
document.getElementById(
'menuBtn'
)

const navLinks =
document.getElementById(
'navLinks'
)

menuBtn.onclick = () => {

navLinks.classList.toggle(
'active'
)

}

const observer =
new IntersectionObserver(
entries => {

entries.forEach(entry => {

if(entry.isIntersecting){

entry.target.classList.add(
'show'
)

}

})

}
)

document
.querySelectorAll('.fade-up')
.forEach(el => {

observer.observe(el)

})

document
.querySelectorAll('button')
.forEach(btn => {

btn.addEventListener(
'mouseenter',
() => {

btn.style.transform =
'scale(1.05)'

}
)

btn.addEventListener(
'mouseleave',
() => {

btn.style.transform =
'scale(1)'

}
)

})

</script>
`
}
