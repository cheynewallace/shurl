package main

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"time"
)

// Page is the short to long url conversion object
type Page struct {
	ID        int64
	ShortPath string
	LongURL   string
	CreatedAt time.Time
}

func (p *Page) save() error {
	filename := p.shortURLHASH()
	return ioutil.WriteFile("redirs/" + filename, []byte(p.LongURL), 0600)
}

func (p *Page) shortURLHASH() string {
	return fmt.Sprintf("%x", md5.Sum([]byte(p.ShortPath)))
}

func loadPage(path string) (*Page, error) {
	filename := fmt.Sprintf("%x", md5.Sum([]byte(path)))
	longurl, err := ioutil.ReadFile("redirs/" + filename)
	if err != nil {
		return nil, err
	}
	return &Page{ShortPath: path, LongURL: string(longurl)}, nil
}
