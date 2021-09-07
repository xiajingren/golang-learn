package main

import (
	"log"
	"net/http"
	"text/template"
)

func main() {
	ts := loadTemplates()
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		filename := r.URL.Path[1:]
		t := ts.Lookup(filename)
		if t != nil {
			err := t.Execute(rw, "hello word!")
			if err != nil {
				log.Fatalln(err.Error())
			}
		} else {
			http.NotFound(rw, r)
		}
	})

	http.Handle("/css/", http.FileServer(http.Dir("static")))
	http.Handle("/img/", http.FileServer(http.Dir("static")))

	http.ListenAndServe(":8080", nil)
}

func loadTemplates() *template.Template {
	ts := template.New("templates")
	template.Must(ts.ParseGlob("templates/*.html"))
	return ts
}
