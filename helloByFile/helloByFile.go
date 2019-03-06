package main

import (
	"log"
	"net/http"
	"path/filepath"

	"text/template"
)

// templ is once template.
type templateHandler struct {
	filename string
	templ    *template.Template
}

// ServHTTP process HTTP Request
func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	t.templ.Execute(w, r)
}

func main() {
	// router
	http.Handle("/", &templateHandler{filename: "hello.html"})

	// web server start
	if err := http.ListenAndServe(":8090", nil); err != nil {
		log.Fatal("LinstenAndServe:", err)
	}
}
