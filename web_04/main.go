package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Person struct {
	Name string
	Tags []string
}

func main() {
	http.HandleFunc("/post", func(rw http.ResponseWriter, r *http.Request) {
		// 读取查询参数
		fmt.Fprintln(rw, r.URL.RawQuery)
		fmt.Fprintln(rw, r.URL.Query())
		fmt.Fprintln(rw, r.URL.Query()["id"])
		fmt.Fprintln(rw, r.URL.Query().Get("id"))

		// 读取 Hander 参数
		fmt.Fprintln(rw, r.Header)
		fmt.Fprintln(rw, r.Header["My-Header"])
		fmt.Fprintln(rw, r.Header.Get("My-Header"))

		// 读取 request body 内容
		body := make([]byte, r.ContentLength)
		r.Body.Read(body)
		fmt.Fprintln(rw, string(body))

		// r.ParseForm()
		// r.Form["key"]
		// r.FormValue("key")
		// r.PostForm["key"]
		// r.PostFormValue("key")
		// r.FormFile("key")

		// r.ParseMultipartForm(1024)
		// r.MultipartForm["key"]
		// r.MultipartForm.File["uploadfile"][0].Open()

		// r.MultipartReader() // 以 steam 读取
	})

	// 返回 json 数据
	http.HandleFunc("/json", func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type", "application/json")
		p := &Person{
			Name: "haha",
			Tags: []string{"aa", "bb", "cc"},
		}
		j, _ := json.Marshal(p)
		//rw.WriteHeader(http.StatusOK)
		rw.Write(j)
	})

	http.ListenAndServe(":8080", nil)
}
