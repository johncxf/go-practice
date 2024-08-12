// 一个简单的 http 服务
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// http://127.0.0.1:8888/hello?name=word
	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		params := request.URL.Query()
		fmt.Fprintf(writer, "hello, %s", params.Get("name"))
	})
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		log.Fatalf("启动 HTTP 服务器失败: %v", err)
	}
}
