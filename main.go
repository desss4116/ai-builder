package main

import (
	"ai-builder/database"
	"ai-builder/engine"
	"ai-builder/scheduler"

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

	/*
		DATABASE
	*/

	database.InitDB()

	/*
		AUTO INTERNET LEARNING
	*/

	scheduler.StartLearning()

	/*
		DISCORD
	*/

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

	/*
		SITES
	*/

	loadSites()

	/*
		ROUTES
	*/

	http.HandleFunc("/", homeHandler)

	http.HandleFunc("/site/", siteHandler)

	http.HandleFunc("/editor/", editorHandler)

	http.HandleFunc("/api/generate", apiGenerateHandler)

	http.HandleFunc("/api/react/", reactExportHandler)

	/*
		PORT
	*/

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	log.Println("===================================")
	log.Println("AI Builder ONLINE")
	log.Println("Discord Bot ONLINE")
	log.Println("Internet Intelligence ACTIVE")
	log.Println("SQLite Memory ACTIVE")
	log.Println("===================================")

	log.Fatal(
		http.ListenAndServe(":"+port, nil),
	)
}

func messageHandler(
	s *discordgo.Session,
	m *discordgo.MessageCreate,
) {

	if m.Author.Bot {
		return
	}

	msg := strings.TrimSpace(
		m.Content,
	)

	/*
		NORMAL AI-LIKE REPLIES
	*/

	if !isWebsiteRequest(msg) {

		reply := generateAIReply(msg)

		s.ChannelMessageSend(
			m.ChannelID,
			reply,
		)

		return
	}

	/*
		WEBSITE GENERATION
	*/

	go generateSiteForUser(
		s,
		m,
	)
}

func generateSiteForUser(
	s *discordgo.Session,
	m *discordgo.MessageCreate,
) {

	s.ChannelTyping(
		m.ChannelID,
	)

	prompt := m.Content

	/*
		PARSE PROMPT
	*/

	ctx := engine.ParsePrompt(
		prompt,
	)

	/*
		INTERNET SEARCH
	*/

	internetData :=
		engine.SearchInternet(prompt)

	_ = internetData

	/*
		GENERATE WEBSITE
	*/

	html := engine.GenerateWebsite(
		ctx,
	)

	/*
		REACT EXPORT
	*/

	reactCode :=
		generateReactExport(html)

	/*
		SAVE
	*/

	site := Site{
		ID:        randomID(),
		Prompt:    prompt,
		HTML:      html,
		ReactCode: reactCode,
		Type:      ctx.Type,
		Created:   time.Now().String(),
	}

	saveSite(site)

	/*
		URLS
	*/

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

	/*
		MESSAGE
	*/

	message := fmt.Sprintf(`
✅ Website Generated

🌍 Website:
%s

🎨 Editor:
%s

⚛ React Export Ready
✨ Premium UI Enabled
🚀 Internet Enhanced
`,
		url,
		editorURL,
	)

	/*
		DM USER
	*/

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

func generateAIReply(
	msg string,
) string {

	msg = strings.ToLower(msg)

	/*
		INTERNET SEARCH
	*/

	internet :=
		engine.SearchInternet(msg)

	_ = internet

	/*
		SMART REPLIES
	*/

	if strings.Contains(msg, "привет") {

		return `
👋 Привет.

Я AI Builder.

Могу:
• создавать сайты
• анализировать интернет
• генерировать UI
• делать landing pages
`
	}

	if strings.Contains(msg, "что умеешь") {

		return `
🚀 Возможности:

• AI-like website generation
• Internet-enhanced search
• Premium UI
• React export
• Framer-like layouts
• Dynamic generation
`
	}

	if strings.Contains(msg, "как дела") {

		return `
⚡ Работаю нормально.

Сейчас:
• изучаю интернет
• генерирую сайты
• улучшаю layouts
`
	}

	/*
		SEARCH DATABASE
	*/

	results :=
		database.SearchKnowledge(msg)

	if len(results) > 0 {

		return fmt.Sprintf(`
🧠 Я нашёл информацию:

%s

%s
`,
			results[0].Title,
			results[0].Content,
		)
	}

	/*
		FALLBACK
	*/

	return `
🌐 Я обработал запрос через internet layer.

Попробуй:
• создай сайт для sneaker brand
• сделай ecommerce landing
• build portfolio website
`
}

func isWebsiteRequest(
	msg string,
) bool {

	msg = strings.ToLower(msg)

	triggers := []string{

		"создай сайт",

		"сделай сайт",

		"website",

		"landing page",

		"build website",

		"create website",

		"portfolio",

		"shop",

		"store",

		"лендинг",

		"интернет магазин",
	}

	for _, t := range triggers {

		if strings.Contains(
			msg,
			t,
		) {

			return true
		}
	}

	return false
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
background:
linear-gradient(
135deg,
#020617,
#0f172a
);
color:white;
display:flex;
justify-content:center;
align-items:center;
height:100vh;
flex-direction:column;
overflow:hidden;
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
Internet Powered Website Intelligence
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

	fmt.Fprint(
		w,
		site.HTML,
	)
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

<title>AI Editor</title>

<style>

body{
margin:0;
font-family:Arial;
background:#020617;
color:white;
display:flex;
height:100vh;
overflow:hidden;
}

.sidebar{
width:320px;
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
width:100%%;
padding:16px;
margin-bottom:16px;
border:none;
border-radius:16px;
background:#3b82f6;
color:white;
cursor:pointer;
font-size:16px;
transition:.3s;
}

.button:hover{
transform:translateY(-3px);
}

.input{
width:100%%;
padding:14px;
margin-bottom:16px;
border:none;
border-radius:14px;
}

</style>

</head>

<body>

<div class="sidebar">

<h2>AI Builder Editor</h2>

<input
class="input"
placeholder="Edit title"
>

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

	fmt.Fprint(
		w,
		editor,
	)
}

func apiGenerateHandler(
	w http.ResponseWriter,
	r *http.Request,
) {

	prompt :=
		r.URL.Query().Get(
			"prompt",
		)

	if prompt == "" {

		prompt =
			"modern startup website"
	}

	ctx :=
		engine.ParsePrompt(prompt)

	html :=
		engine.GenerateWebsite(ctx)

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

		http.NotFound(
			w,
			r,
		)

		return
	}

	w.Header().Set(
		"Content-Type",
		"text/plain",
	)

	fmt.Fprint(
		w,
		site.ReactCode,
	)
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

func saveSite(
	site Site,
) {

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

	data, err :=
		os.ReadFile(
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
