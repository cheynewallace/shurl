package main

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"html/template"
	"log"
	"net/http"
	"os"
)

var db *sql.DB

func redirHandler(w http.ResponseWriter, r *http.Request) {
	page, err := queryPage(r.URL.Path[1:])
	if err != nil || page.LongURL == "" {
		fmt.Fprintf(w, "No Matches Found")
		return
	}
	fmt.Printf("Redirecting %s To %s\n", page.ShortURL, page.LongURL)
	http.Redirect(w, r, page.LongURL, 301)
	return
}

func fileRedirHandler(w http.ResponseWriter, r *http.Request) {
	p2, err := loadPage(r.URL.Path[1:])
	if err != nil {
		fmt.Fprintf(w, "No Matches Found")
		return
	}

	fmt.Println(fmt.Sprintf("%s Redirecting To: %s", r.URL.Path[1:], string(p2.LongURL)))
	http.Redirect(w, r, string(p2.LongURL), 301)
	return
}

func newHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("views/new.html")
	t.Execute(w, nil)
}

func createHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	page := &Page{ShortURL: r.FormValue("shortpath"), LongURL: r.FormValue("longurl")}
	err := page.save()
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	fmt.Println(fmt.Sprintf("Created: %s Redirects To %s", page.ShortURL, page.LongURL))
	renderTemplate(w, "views/create.html", page)
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	if t, err := template.ParseFiles(tmpl); err != nil {
		http.Error(w, err.Error(), 500)
	} else {
		t.Execute(w, p)
	}
}

func serveResource(w http.ResponseWriter, req *http.Request) {
	log.Fatal("DGFDSG")
	path := "../../public" + req.URL.Path
	http.ServeFile(w, req, path)
}

func main() {
	// Init Database
	// https://godoc.org/github.com/lib/pq
	var err error
	db, err = sql.Open("postgres", fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("SHURL_DB_HOST"), os.Getenv("SHURL_DB_USER"), os.Getenv("SHURL_DB_PASS"), os.Getenv("SHURL_DB_NAME")))
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()
	s := http.StripPrefix("/public/", http.FileServer(http.Dir("./public/")))
	r.HandleFunc("/new", newHandler).Methods("GET")
	r.HandleFunc("/create", createHandler).Methods("POST")
	r.PathPrefix("/public/").Handler(s)
	r.HandleFunc("/{id:[a-z0-9]{3,8}}", redirHandler)
	http.Handle("/", r)

	http.ListenAndServe(":8080", nil)
}
