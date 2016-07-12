package main

import (
	"time"
)

// Page is the base object used for redirection
type Page struct {
	ID        int8
	ShortURL  string
	LongURL   string
	CreatedAt time.Time
}

func (p *Page) save() error {
	_, err := db.Exec("INSERT INTO pages (shorturl, longurl) VALUES ($1, $2)", p.ShortURL, p.LongURL)
	return err
}

func queryPage(path string) (*Page, error) {
	var id int8
	var shorturl string
	var longurl string
	var createdAt string
	err := db.QueryRow("SELECT * FROM pages WHERE shorturl = $1", path).Scan(&id, &shorturl, &longurl, &createdAt)
	if err != nil {
		return nil, err
	}
	return &Page{ID: id, LongURL: longurl, ShortURL: shorturl}, nil
}
