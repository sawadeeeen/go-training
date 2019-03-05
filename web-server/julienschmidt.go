package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"html"
	"log"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func main() {
	router := httprouter.New()
	router.GET("/:path", Index)

	log.Fatal(http.ListenAndServe(":8080", router))
}
