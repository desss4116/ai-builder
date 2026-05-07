package engine

func GlobalStyles() string {

	return `
<style>

*{
margin:0;
padding:0;
box-sizing:border-box;
}

html{
scroll-behavior:smooth;
}

body{
font-family:
-apple-system,
BlinkMacSystemFont,
sans-serif;

background:#020617;
color:white;
overflow-x:hidden;
}

.container{
width:min(1200px,92%);
margin:auto;
}

section{
padding:120px 0;
}

h1{
font-size:72px;
line-height:1;
font-weight:800;
}

h2{
font-size:54px;
margin-bottom:24px;
}

p{
font-size:18px;
line-height:1.8;
opacity:.8;
}

.grid{
display:grid;
grid-template-columns:
repeat(auto-fit,minmax(320px,1fr));
gap:24px;
}

.card{
background:
rgba(255,255,255,.05);

backdrop-filter:
blur(20px);

border:
1px solid rgba(255,255,255,.08);

padding:32px;

border-radius:32px;

transition:.3s;
}

.card:hover{
transform:
translateY(-8px);
}

.btn{
padding:
18px 34px;

border:none;

border-radius:999px;

background:white;

color:black;

font-weight:700;

cursor:pointer;

transition:.3s;
}

.btn:hover{
transform:
scale(1.05);
}

.navbar{
position:fixed;
top:0;
left:0;
width:100%;
z-index:999;

background:
rgba(0,0,0,.5);

backdrop-filter:
blur(20px);

border-bottom:
1px solid rgba(255,255,255,.08);
}

.nav-inner{
height:80px;

display:flex;
align-items:center;
justify-content:space-between;
}

.logo{
font-size:24px;
font-weight:800;
}

.nav-links{
display:flex;
gap:32px;
}

.nav-links a{
color:white;
text-decoration:none;
opacity:.8;
transition:.3s;
}

.nav-links a:hover{
opacity:1;
}

.mobile-menu{
display:none;
font-size:32px;
cursor:pointer;
}

.hero{
min-height:100vh;

display:flex;
align-items:center;
text-align:center;
}

.hero-content{
max-width:900px;
margin:auto;
}

.hero p{
margin:
32px auto;
max-width:700px;
}

.features{
background:
linear-gradient(
180deg,
transparent,
rgba(255,255,255,.03)
);
}

.product-card img{
width:100%;
border-radius:24px;
margin-bottom:20px;
}

.testimonial{
font-size:20px;
line-height:1.7;
}

.price{
font-size:64px;
font-weight:800;
margin:20px 0;
}

.footer{
padding:80px 0;
border-top:
1px solid rgba(255,255,255,.08);
}

.fade-up{
opacity:0;
transform:
translateY(40px);
transition:1s;
}

.fade-up.show{
opacity:1;
transform:
translateY(0);
}

@media(max-width:768px){

h1{
font-size:52px;
}

h2{
font-size:38px;
}

.nav-links{
position:fixed;
top:80px;
right:-100%;
width:260px;
height:100vh;

background:#020617;

flex-direction:column;
padding:40px;

transition:.4s;
}

.nav-links.active{
right:0;
}

.mobile-menu{
display:block;
}

section{
padding:90px 0;
}

}

</style>
`
}
