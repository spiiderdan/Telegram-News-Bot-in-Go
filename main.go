package main

import (
	"log"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/mmcdole/gofeed"
)

func main() {
	// Initialize Telegram bot
	bot, err := tgbotapi.NewBotAPI(TELEGRAM_APITOKEN)
	if err != nil {
		log.Panic(err)
	}
	log.Printf("Authorized on account %s", bot.Self.UserName)

	// Initialize SQLite database
	initDB()

	// RSS feed parser
	parser := gofeed.NewParser()

	for {
		for _, feedURL := range RSS_FEED_URLS {
			feed, err := parser.ParseURL(feedURL)
			if err != nil {
				log.Printf("Failed to fetch RSS (%s): %v", feedURL, err)
				continue
			}
			
			log.Println("Checked:", feedURL) //checks if the news has been sent before

			if len(feed.Items) > 0 {
				item := feed.Items[0]

				// Use GUID or fallback to Link
				articleID := item.GUID
				if articleID == "" {
					articleID = item.Link
				}

				if !articleAlreadySent(feedURL, articleID) {
					message := "<b>" + escapeHTML(item.Title) + "</b>\n" + item.Link

					msg := tgbotapi.NewMessage(TELEGRAM_CHATID, message)
					msg.ParseMode = "HTML"

					_, err := bot.Send(msg)
					if err != nil {
						log.Println("Failed to send message:", err)
					} else {
						log.Println("Message sent:", item.Title)
						markArticleAsSent(feedURL, articleID)
					}
				}
			}
		}

		time.Sleep(1 * time.Minute)
	}
}
