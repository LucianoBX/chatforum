package main

import (
	. "github.com/chatforum/routes"
	"log"
	"net/http"
)

// 通过指定端口启动 Web 服务器
func startWebServer(port string) {
	r := NewRouter() // 通过 router.go 中定义的路由器来分发请求

	// 处理静态资源文件
	assets := http.FileServer(http.Dir("public"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", assets))

	http.Handle("/", r) // 将路由器应用到HTTP服务器
	log.Println("Starting HTTP service at " + port)
	err := http.ListenAndServe(":"+port, nil) // 启动协程监听请求

	if err != nil {
		log.Println("An error occured starting HTTP listener at port " + port)
		log.Println("Error: " + err.Error())
	}
}

func main() {

	startWebServer("8080")
}
