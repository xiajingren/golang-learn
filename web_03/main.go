package main

import "net/http"

func main() {
	// 静态文件 server

	// 1
	// http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
	// 	http.ServeFile(rw, r, "wwwroot"+r.URL.Path)
	// })
	// http.ListenAndServe(":8080", nil)

	// 2
	http.ListenAndServe(":8080", http.FileServer(http.Dir("wwwroot")))
}
