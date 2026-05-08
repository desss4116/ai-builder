package summarizer

import "strings"

func Summarize(content string) string {

	if len(content) > 1500 {
		content = content[:1500]
	}

	content = strings.TrimSpace(content)

	return `
🧠 AI Summary:

` + content
}
