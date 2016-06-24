package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
)

func redirHandler(w http.ResponseWriter, r *http.Request) {
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
	page := &Page{ShortPath: r.FormValue("shortpath"), LongURL: r.FormValue("longurl")}
	page.save()
	fmt.Println(fmt.Sprintf("%s - %s", page.ShortPath, page.LongURL))
	renderTemplate(w, "views/create.html", page)
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	if t, err := template.ParseFiles(tmpl); err != nil {
		http.Error(w, err.Error(), 500)
	} else {
		t.Execute(w, p)
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/new", newHandler)
	r.HandleFunc("/create", createHandler)
	r.HandleFunc("/{id:[a-z0-9]{4}}", redirHandler)
	http.Handle("/", r)

	http.ListenAndServe(":8080", nil)
}
