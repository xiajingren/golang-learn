package main

import (
	"net/http"
	"text/template"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		filename := r.URL.Path[1:]
		t, err := template.ParseFiles("templates/layout.html", "templates/"+filename, "templates/foot.html")
		if err != nil {
			http.NotFound(rw, r)
			return
		}

		//t.Execute(rw, nil) // 默认执行layout.html 那么index.html,home.html中就无需写{{ template "layout.html"}}
		t.ExecuteTemplate(rw, filename, nil)
	})

	http.ListenAndServe(":8080", nil)
}
