package main

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"time"
)

// Page is the short to long url conversion object
type Page struct {
	ID        int8
	ShortURL  string
	LongURL   string
	CreatedAt time.Time
}

// DB Based
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

// File Store Based
func (p *Page) saveFile() error {
	filename := p.shortURLHASH()
	return ioutil.WriteFile("redirs/"+filename, []byte(p.LongURL), 0600)
}

func (p *Page) shortURLHASH() string {
	return fmt.Sprintf("%x", md5.Sum([]byte(p.ShortURL)))
}

func loadPage(path string) (*Page, error) {
	filename := fmt.Sprintf("%x", md5.Sum([]byte(path)))
	longurl, err := ioutil.ReadFile("redirs/" + filename)
	if err != nil {
		return nil, err
	}
	return &Page{ShortURL: path, LongURL: string(longurl)}, nil
}
