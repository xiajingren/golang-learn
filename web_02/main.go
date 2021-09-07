package main

import "net/http"

type myHandler struct{}

func (m myHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("myHandler"))
}

func main() {
	// 方法1 实际上是调用了 http.Server这个结构体 的 ListenAndServe方法
	//http.ListenAndServe(":8080", nil)
	//http.ListenAndServeTLS ...

	// 方法2 配置更灵活
	// server := http.Server{
	// 	Addr:    ":8080",
	// 	Handler: nil,
	// }
	// server.ListenAndServe()
	//server.ListenAndServeTLS ...

	// 以上2种方法是等效的

	// 使用自定义Handler
	// mh := myHandler{}
	// server := http.Server{
	// 	Addr:    ":8080",
	// 	Handler: mh,
	// }
	// server.ListenAndServe()

	mh := myHandler{}
	server := http.Server{
		Addr:    ":8080",
		Handler: nil, //使用 DefaultServeMux
	}
	// 为路径 /my 指定 handler
	http.Handle("/my", mh)
	// 使用 http.HandleFunc，http.HandleFunc内部使用http.HandlerFunc类型转换（http.HandlerFunc是一个函数类型，本身就是一个handler），将具有适当签名的函数转换为一个handler 本质还是调用 http.Handle
	http.HandleFunc("/home", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("home"))
	})
	server.ListenAndServe()
}
