package main

import (
	"net/http"
	"text/template"
)

func main() {
	server := http.Server{
		Addr: ":8080",
	}

	http.HandleFunc("/hello", func(rw http.ResponseWriter, r *http.Request) {
		// 1
		// t := template.New("index.html")
		// t, _ = t.ParseFiles("template/index.html")

		// 2
		t, _ := template.ParseFiles("template/index.html")

		// 3 解析所有符合规则的模板
		//t, _ := template.ParseGlob("template/*.html")

		// 4 直接解析字符串
		// t := template.New("index.html")
		// t.Parse("<html><head><title>Document</title></head><body>{{.}}</body></html>")

		t.Execute(rw, "Hello Word!  "+t.Name())
	})

	server.ListenAndServe()
}
