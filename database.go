package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func initDB() {
	var err error
	db, err = sql.Open("sqlite3", "rssbot.db")
	if err != nil {
		log.Fatal("Failed to open database:", err)
	}

	createTable := `
	CREATE TABLE IF NOT EXISTS sent_articles (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		feed_url TEXT NOT NULL,
		article_id TEXT NOT NULL,
		UNIQUE(feed_url, article_id)
	);`

	_, err = db.Exec(createTable)
	if err != nil {
		log.Fatal("Failed to create table:", err)
	}
}

func articleAlreadySent(feedURL, articleID string) bool {
	var count int
	err := db.QueryRow(`SELECT COUNT(*) FROM sent_articles WHERE feed_url = ? AND article_id = ?`, feedURL, articleID).Scan(&count)
	if err != nil {
		log.Println("DB check error:", err)
		return false
	}
	return count > 0
}

func markArticleAsSent(feedURL, articleID string) {
	_, err := db.Exec(`INSERT OR IGNORE INTO sent_articles (feed_url, article_id) VALUES (?, ?)`, feedURL, articleID)
	if err != nil {
		log.Println("DB insert error:", err)
	}
}
