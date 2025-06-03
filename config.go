package main

const TELEGRAM_APITOKEN = "7766314037:AAHoJCIYfg9eucbkoZF_xqJHIUnhgkD-Gv8" // Go News Bot token
const TELEGRAM_CHATID = -1002504264864 // Go News Channel ID
var RSS_FEED_URLS  = [] string{
	"https://tr.cointelegraph.com/rss", // Cointelegraph
	"https://www.coindesk.com/arc/outboundfeeds/rss/", // Coindesk
	"https://bitcoinmagazine.com/.rss/full/", // Bitcoinmagezine
	"https://ethereumworldnews.com/feed/", // Ethereum World News
	"https://www.newsbtc.com/feed/", // NewsBTC
}