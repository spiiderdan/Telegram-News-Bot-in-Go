# 🤖 Telegram News Bot (in Go)

A lightweight Telegram bot written in Go that fetches the latest articles from multiple RSS feeds and automatically sends them to a Telegram group or channel.

It uses `gofeed` to parse RSS feeds and the official Telegram Bot API to post formatted updates. Sent articles are tracked and saved locally using a JSON file to avoid duplicate messages — even after restarting the bot.

---

## 🚀 Features

- Monitors multiple RSS feeds
- Sends formatted Telegram HTML messages (title, summary, link)
- Lightweight — no external database required
- HTML-safe message formatting

---

## 🛠 Configuration

The actual configuration (Telegram bot token, chat ID, and RSS feed URLs) is **not included in the repo** for security reasons.

To run the bot, create a `config.go` file in the root of the project with this content:

```go
package main

const TELEGRAM_APITOKEN = "your-telegram-bot-token"
const TELEGRAM_CHATID int64 = -12345678910 // your group or channel chat ID

var RSS_FEED_URLS = []string{
	"news feed RSS 1"
	"news feed RSS 2"
	...
}
```

> 💡 Tip: If you don't know your group’s chat ID, use a temporary script with `GetUpdates()` to extract it after sending a message in the group.

---

## 🧠 How It Works

- `main.go` initializes the bot, loads previously sent articles, and starts checking RSS feeds.
- Each feed is parsed and checked for new items.
- The latest article per feed is stored in `state.json` so the same article isn’t re-sent.
- `formatter.go` strips and escapes HTML to make Telegram-safe messages.
- Messages are sent using Telegram's `HTML` parse mode.

---

## 📦 Requirements

- Go 1.18 or higher
- A Telegram bot token from [@BotFather](https://t.me/botfather)
- (Optional) A Telegram group or channel to post into

---

## 🏃‍♂️ Running the Bot

```bash
go run .
```

The bot will start, authenticate with Telegram, and begin posting updates every 5 minutes.

---

## 🔒 Security

This project excludes the `config.go` file from version control to avoid leaking secrets (API token and chat ID).  
Make sure you do the same if you fork or modify this bot.

---

## 📚 References

- [Telegram Bot API](https://core.telegram.org/bots/api) – Official Telegram API documentation
- [go-telegram-bot-api](https://github.com/go-telegram-bot-api/telegram-bot-api) – Go client for the Telegram Bot API
- [gofeed](https://github.com/mmcdole/gofeed) – Go RSS feed parser

---

## 📄 License

MIT License — free to use, modify, and distribute.
