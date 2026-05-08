package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

type Site struct {
	ID      string `json:"id"`
	Prompt  string `json:"prompt"`
	Created string `json:"created"`
	URL     string `json:"url"`
}

var sites []Site

func main() {

	token := os.Getenv("DISCORD_TOKEN")

	if token == "" {
		log.Fatal("DISCORD_TOKEN missing")
	}

	loadSites()

	dg, err := discordgo.New("Bot " + token)

	if err != nil {
		log.Fatal(err)
	}

	dg.AddHandler(messageHandler)

	dg.Identify.Intents = discordgo.IntentsGuildMessages |
		discordgo.IntentsMessageContent

	err = dg.Open()

	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", home)

	http.HandleFunc("/site/", renderSite)

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	go func() {
		log.Println("Web running on :" + port)
		http.ListenAndServe(":"+port, nil)
	}()

	log.Println("Bot online")

	select {}
}

func messageHandler(
	s *discordgo.Session,
	m *discordgo.MessageCreate,
) {

	if m.Author.Bot {
		return
	}

	content := strings.ToLower(m.Content)

	if isWebsiteRequest(content) {

		id := randomID()

		url := os.Getenv("RAILWAY_PUBLIC_DOMAIN")

		if url == "" {
			url = "localhost:8080"
		}

		siteURL := fmt.Sprintf(
			"https://%s/site/%s",
			url,
			id,
		)

		site := Site{
			ID:      id,
			Prompt:  m.Content,
			Created: time.Now().String(),
			URL:     siteURL,
		}

		sites = append(sites, site)

		saveSites()

		s.ChannelMessageSend(
			m.ChannelID,
			"🚀 Website generated:\n"+siteURL,
		)

		return
	}

	answer := smartAnswer(content)

	s.ChannelMessageSend(
		m.ChannelID,
		answer,
	)
}

func smartAnswer(input string) string {

	if strings.Contains(input, "привет") {
		return "👋 Привет. Я AI Builder нового поколения."
	}

	if strings.Contains(input, "пока") {
		return "👋 До встречи."
	}

	if strings.Contains(input, "лучшие фильмы") {
		return `🎬 Топ фильмов:

1. Interstellar
2. The Dark Knight
3. Fight Club
4. Inception
5. Blade Runner 2049
6. Dune
7. The Matrix
8. Whiplash
9. Parasite
10. Oppenheimer`
	}

	if strings.Contains(input, "что значит") {

		word := strings.ReplaceAll(
			input,
			"что значит",
			"",
		)

		return "📖 Значение:\n" +
			strings.TrimSpace(word) +
			" — термин или понятие, требующее контекста."
	}

	return "🧠 Я анализирую запрос: " + input
}

func isWebsiteRequest(text string) bool {

	keywords := []string{
		"создай",
		"website",
		"site",
		"landing",
		"лендинг",
		"сайт",
		"make website",
	}

	for _, k := range keywords {

		if strings.Contains(text, k) {
			return true
		}
	}

	return false
}

func renderSite(
	w http.ResponseWriter,
	r *http.Request,
) {

	id := strings.TrimPrefix(
		r.URL.Path,
		"/site/",
	)

	var prompt string

	for _, s := range sites {

		if s.ID == id {
			prompt = s.Prompt
		}
	}

	if prompt == "" {
		http.NotFound(w, r)
		return
	}

	html := buildNextGenWebsite(prompt)

	w.Header().Set(
		"Content-Type",
		"text/html",
	)

	fmt.Fprint(w, html)
}

func buildNextGenWebsite(prompt string) string {

	title := strings.Title(prompt)

	gradient := randomGradient()

	return `
<!DOCTYPE html>
<html>
<head>

<meta charset="UTF-8"/>

<meta name="viewport"
content="width=device-width, initial-scale=1.0"/>

<title>` + title + `</title>

<script src="https://cdn.tailwindcss.com"></script>

<link rel="preconnect"
href="https://fonts.googleapis.com">

<link href="https://fonts.googleapis.com/css2?family=Inter:wght@300;400;500;600;700;800;900&display=swap"
rel="stylesheet">

<style>

*{
margin:0;
padding:0;
box-sizing:border-box;
}

body{
font-family:'Inter',sans-serif;
background:black;
color:white;
overflow-x:hidden;
}

.hero{
min-height:100vh;
display:flex;
align-items:center;
justify-content:center;
padding:80px 24px;
position:relative;
overflow:hidden;
}

.glow{
position:absolute;
width:700px;
height:700px;
background:` + gradient + `;
filter:blur(140px);
opacity:0.3;
border-radius:999px;
animation:float 8s ease-in-out infinite;
}

.container{
max-width:1300px;
width:100%;
position:relative;
z-index:2;
}

.title{
font-size:clamp(56px,10vw,140px);
font-weight:900;
line-height:0.9;
letter-spacing:-6px;
max-width:1100px;
}

.subtitle{
margin-top:40px;
font-size:24px;
color:#b4b4b4;
max-width:700px;
line-height:1.6;
}

.buttons{
display:flex;
gap:20px;
margin-top:50px;
flex-wrap:wrap;
}

.btn{
padding:18px 42px;
border-radius:999px;
font-size:18px;
font-weight:700;
border:none;
cursor:pointer;
transition:0.4s;
}

.primary{
background:white;
color:black;
}

.primary:hover{
transform:translateY(-4px) scale(1.03);
}

.secondary{
background:transparent;
border:1px solid rgba(255,255,255,0.15);
color:white;
backdrop-filter:blur(20px);
}

.secondary:hover{
background:rgba(255,255,255,0.05);
}

.grid{
display:grid;
grid-template-columns:repeat(auto-fit,minmax(320px,1fr));
gap:28px;
padding:80px 24px 140px;
max-width:1300px;
margin:auto;
}

.card{
background:rgba(255,255,255,0.04);
border:1px solid rgba(255,255,255,0.08);
backdrop-filter:blur(20px);
border-radius:36px;
padding:40px;
transition:0.4s;
}

.card:hover{
transform:translateY(-8px);
background:rgba(255,255,255,0.07);
}

.card h2{
font-size:42px;
font-weight:800;
margin-bottom:20px;
}

.card p{
font-size:18px;
line-height:1.7;
color:#b4b4b4;
}

@keyframes float{
0%{
transform:translateY(0px);
}
50%{
transform:translateY(-30px);
}
100%{
transform:translateY(0px);
}
}

</style>

</head>

<body>

<section class="hero">

<div class="glow"></div>

<div class="container">

<h1 class="title">
` + title + `
</h1>

<p class="subtitle">
Award-winning next generation digital experience powered by intelligent design systems and premium UI architecture.
</p>

<div class="buttons">

<button class="btn primary">
Launch Experience
</button>

<button class="btn secondary">
Explore Platform
</button>

</div>

</div>

</section>

<section class="grid">

<div class="card">
<h2>Luxury UI</h2>
<p>
Emotion-driven premium interface with cinematic spacing and ultra modern composition system.
</p>
</div>

<div class="card">
<h2>Motion Engine</h2>
<p>
Smooth interactions, premium animations and immersive experience architecture.
</p>
</div>

<div class="card">
<h2>AI Systems</h2>
<p>
Advanced intelligent generation pipeline inspired by modern product ecosystems.
</p>
</div>

</section>

</body>
</html>
`
}

func randomGradient() string {

	gradients := []string{
		"linear-gradient(135deg,#ff0080,#7928ca)",
		"linear-gradient(135deg,#00c6ff,#0072ff)",
		"linear-gradient(135deg,#00ff87,#60efff)",
		"linear-gradient(135deg,#f7971e,#ffd200)",
		"linear-gradient(135deg,#7928ca,#ff0080)",
	}

	return gradients[rand.Intn(len(gradients))]
}

func randomID() string {

	letters := "abcdefghijklmnopqrstuvwxyz1234567890"

	result := ""

	for i := 0; i < 12; i++ {

		result += string(
			letters[rand.Intn(len(letters))],
		)
	}

	return result
}

func home(
	w http.ResponseWriter,
	r *http.Request,
) {

	fmt.Fprint(
		w,
		"AI Builder running",
	)
}

func saveSites() {

	file, _ := json.MarshalIndent(
		sites,
		"",
		"  ",
	)

	os.WriteFile(
		"data/sites.json",
		file,
		0644,
	)
}

func loadSites() {

	file, err := os.ReadFile(
		"data/sites.json",
	)

	if err != nil {
		return
	}

	json.Unmarshal(file, &sites)
}
