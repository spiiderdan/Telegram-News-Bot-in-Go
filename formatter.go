package main

import (
	"regexp"
	"strings"
)

// Remove all HTML tags from the input
func stripHTML(input string) string {
	re := regexp.MustCompile(`<[^>]*>`)
	return re.ReplaceAllString(input, "")
}

// Escape characters that Telegram HTML formatting requires
func escapeHTML(s string) string {
	replacer := strings.NewReplacer(
		"&", "&amp;",
		"<", "&lt;",
		">", "&gt;",
		`"`, "&quot;",
	)	
	return replacer.Replace(s)
}

// FormatNewsMessage builds the final message in Telegram HTML format
func FormatNewsMessage(title, description, link string) string {
	cleanTitle := escapeHTML(title)
	cleanDesc := escapeHTML(stripHTML(description))

	if len(cleanDesc) > 300 {
		cleanDesc = cleanDesc[:297] + "..."
	}

	return "<b>" + cleanTitle + "</b>\n\n" +
		cleanDesc + "\n\n" +
		"<a href=\"" + link + "\">Read full article</a>"
}
