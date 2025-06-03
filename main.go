package main

import (
	"log"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/mmcdole/gofeed"
)

// lastPublishedMap keeps track of latest sent news per feed
var lastPublishedMap = make(map[string]string)

func main() {
	bot, err := tgbotapi.NewBotAPI(TELEGRAM_APITOKEN)
	if err != nil {
		panic(err)
	}
	log.Printf("Authorized on account %s", bot.Self.UserName)

	parser := gofeed.NewParser()

	for {
		for _, feedURL := range RSS_FEED_URLS {
			feed, err := parser.ParseURL(feedURL)
			if err != nil {
				log.Printf("Failed to fetch RSS (%s): %v", feedURL, err)
				continue
			}

			if len(feed.Items) > 0 {
				item := feed.Items[0]
				lastSent := lastPublishedMap[feedURL]

				if item.Published != lastSent {
					lastPublishedMap[feedURL] = item.Published

					message := FormatNewsMessage(item.Title, item.Description, item.Link)


					msg := tgbotapi.NewMessage(TELEGRAM_CHATID, message)
					msg.ParseMode = "HTML"

					_, err := bot.Send(msg)
					if err != nil {
						log.Println("Failed to send message:", err)
					} else {
						log.Println("Message sent:", item.Title)
					}
				}
			}
		}

		time.Sleep(5 * time.Minute)
	}
}