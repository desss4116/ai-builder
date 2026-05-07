package main

import (
	"ai-builder/engine"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/bwmarrin/discordgo"
)

type Site struct {
	ID        string `json:"id"`
	Prompt    string `json:"prompt"`
	HTML      string `json:"html"`
	ReactCode string `json:"react_code"`
	Type      string `json:"type"`
	Created   string `json:"created"`
}

var (
	sites = map[string]Site{}
	mutex sync.RWMutex
)

func main() {

	token := os.Getenv("DISCORD_TOKEN")

	if token == "" {
		log.Fatal("DISCORD_TOKEN missing")
	}

	dg, err := discordgo.New("Bot " + token)

	if err != nil {
		log.Fatal(err)
	}

	dg.Identify.Intents =
		discordgo.IntentsGuildMessages |
			discordgo.IntentsDirectMessages |
			discordgo.IntentsMessageContent

	dg.AddHandler(messageHandler)

	err = dg.Open()

	if err != nil {
		log.Fatal(err)
	}

	loadSites()

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/site/", siteHandler)
	http.HandleFunc("/editor/", editorHandler)
	http.HandleFunc("/api/generate", apiGenerateHandler)
	http.HandleFunc("/api/react/", reactExportHandler)

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	log.Println("===================================")
	log.Println("AI Builder V2 ONLINE")
	log.Println("Discord Bot ONLINE")
	log.Println("Server running on :", port)
	log.Println("===================================")

	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func messageHandler(
	s *discordgo.Session,
	m *discordgo.MessageCreate,
) {

	if m.Author.Bot {
		return
	}

	msg := strings.TrimSpace(m.Content)

	if !isWebsiteRequest(msg) {

		s.ChannelMessageSend(
			m.ChannelID,
			randomReply(),
		)

		return
	}

	go generateSiteForUser(s, m)
}

func generateSiteForUser(
	s *discordgo.Session,
	m *discordgo.MessageCreate,
) {

	s.ChannelTyping(m.ChannelID)

	prompt := m.Content

	ctx := engine.ParsePrompt(prompt)

	/*
		INTERNET ENHANCED GENERATION
	*/

	trends := engine.FetchWebsite(
		"https://www.apple.com",
	)

	_ = trends

	/*
		LAYOUT ENGINE
	*/

	html := engine.GenerateWebsite(ctx)

	/*
		REACT EXPORT
	*/

	reactCode := generateReactExport(html)

	site := Site{
		ID:        randomID(),
		Prompt:    prompt,
		HTML:      html,
		ReactCode: reactCode,
		Type:      ctx.Type,
		Created:   time.Now().String(),
	}

	saveSite(site)

	domain := os.Getenv("DOMAIN")

	url := fmt.Sprintf(
		"%s/site/%s",
		domain,
		site.ID,
	)

	editorURL := fmt.Sprintf(
		"%s/editor/%s",
		domain,
		site.ID,
	)

	message := fmt.Sprintf(`
✅ Website Generated

🌍 Live Website:
%s

🎨 Visual Editor:
%s

⚛ React Export Ready
🚀 Premium UI Generated
✨ Animations Enabled
`,
		url,
		editorURL,
	)

	dm, err := s.UserChannelCreate(
		m.Author.ID,
	)

	if err == nil {

		s.ChannelMessageSend(
			dm.ID,
			message,
		)

		s.ChannelMessageSend(
			m.ChannelID,
			"📩 Website sent in DM",
		)

	} else {

		s.ChannelMessageSend(
			m.ChannelID,
			message,
		)
	}
}

func isWebsiteRequest(msg string) bool {

	msg = strings.ToLower(msg)

	triggers := []string{
		"создай сайт",
		"сделай сайт",
		"website",
		"landing page",
		"build website",
		"create website",
		"shop",
		"store",
		"portfolio",
		"лендинг",
	}

	for _, t := range triggers {

		if strings.Contains(msg, t) {
			return true
		}
	}

	return false
}

func randomReply() string {

	replies := []string{
		"👋 Напиши: создай сайт для sneaker store",
		"🚀 Я умею создавать premium websites",
		"🎨 Попробуй: build website for coffee shop",
		"⚡ Я генерирую Framer-like сайты",
		"🔥 Могу сделать ecommerce landing",
	}

	return replies[
		time.Now().UnixNano()%int64(len(replies))
	]
}

func homeHandler(
	w http.ResponseWriter,
	r *http.Request,
) {

	html := `
<!DOCTYPE html>

<html>

<head>

<title>AI Builder</title>

<style>

body{
margin:0;
font-family:Arial;
background:#020617;
color:white;
display:flex;
justify-content:center;
align-items:center;
height:100vh;
flex-direction:column;
}

h1{
font-size:72px;
margin-bottom:20px;
}

p{
opacity:.7;
font-size:22px;
}

</style>

</head>

<body>

<h1>AI Builder</h1>

<p>
Discord AI Website Generator
</p>

</body>

</html>
`

	fmt.Fprint(w, html)
}

func siteHandler(
	w http.ResponseWriter,
	r *http.Request,
) {

	id := strings.TrimPrefix(
		r.URL.Path,
		"/site/",
	)

	mutex.RLock()

	site, ok := sites[id]

	mutex.RUnlock()

	if !ok {

		http.NotFound(w, r)

		return
	}

	w.Header().Set(
		"Content-Type",
		"text/html",
	)

	fmt.Fprint(w, site.HTML)
}

func editorHandler(
	w http.ResponseWriter,
	r *http.Request,
) {

	id := strings.TrimPrefix(
		r.URL.Path,
		"/editor/",
	)

	mutex.RLock()

	site, ok := sites[id]

	mutex.RUnlock()

	if !ok {

		http.NotFound(w, r)

		return
	}

	editor := fmt.Sprintf(`
<!DOCTYPE html>

<html>

<head>

<title>Visual Editor</title>

<style>

body{
margin:0;
font-family:Arial;
background:#0f172a;
color:white;
display:flex;
height:100vh;
overflow:hidden;
}

.sidebar{
width:300px;
background:#111827;
padding:20px;
overflow:auto;
}

.preview{
flex:1;
background:white;
}

iframe{
width:100%%;
height:100%%;
border:none;
}

.button{
display:block;
width:100%%;
padding:16px;
margin-bottom:16px;
border:none;
background:#3b82f6;
color:white;
border-radius:12px;
cursor:pointer;
font-size:16px;
}

.input{
width:100%%;
padding:14px;
margin-bottom:16px;
border:none;
border-radius:12px;
}

</style>

</head>

<body>

<div class="sidebar">

<h2>AI Builder Editor</h2>

<input class="input" placeholder="Edit title">

<button class="button">
Change Layout
</button>

<button class="button">
Add Section
</button>

<button class="button">
Export React
</button>

<button class="button">
Publish
</button>

</div>

<div class="preview">

<iframe src="/site/%s"></iframe>

</div>

</body>

</html>
`,
		site.ID,
	)

	fmt.Fprint(w, editor)
}

func apiGenerateHandler(
	w http.ResponseWriter,
	r *http.Request,
) {

	prompt := r.URL.Query().Get(
		"prompt",
	)

	if prompt == "" {

		prompt = "modern startup website"
	}

	ctx := engine.ParsePrompt(prompt)

	html := engine.GenerateWebsite(ctx)

	site := Site{
		ID:      randomID(),
		Prompt:  prompt,
		HTML:    html,
		Type:    ctx.Type,
		Created: time.Now().String(),
	}

	saveSite(site)

	domain := os.Getenv("DOMAIN")

	url := fmt.Sprintf(
		"%s/site/%s",
		domain,
		site.ID,
	)

	json.NewEncoder(w).Encode(
		map[string]string{
			"url": url,
		},
	)
}

func reactExportHandler(
	w http.ResponseWriter,
	r *http.Request,
) {

	id := strings.TrimPrefix(
		r.URL.Path,
		"/api/react/",
	)

	mutex.RLock()

	site, ok := sites[id]

	mutex.RUnlock()

	if !ok {

		http.NotFound(w, r)

		return
	}

	w.Header().Set(
		"Content-Type",
		"text/plain",
	)

	fmt.Fprint(w, site.ReactCode)
}

func generateReactExport(
	html string,
) string {

	return fmt.Sprintf(`
export default function GeneratedWebsite(){

return(

<div>

%s

</div>

)

}
`,
		html,
	)
}

func randomID() string {

	return fmt.Sprintf(
		"%d",
		time.Now().UnixNano(),
	)
}

func saveSite(site Site) {

	mutex.Lock()

	defer mutex.Unlock()

	sites[site.ID] = site

	data, _ := json.MarshalIndent(
		sites,
		"",
		"  ",
	)

	os.WriteFile(
		"sites.json",
		data,
		0644,
	)
}

func loadSites() {

	data, err := os.ReadFile(
		"sites.json",
	)

	if err != nil {
		return
	}

	json.Unmarshal(
		data,
		&sites,
	)
}
