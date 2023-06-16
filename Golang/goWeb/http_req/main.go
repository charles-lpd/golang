package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "你发送的请求地址是", r.URL.Path)
	// 返回所有字符串
	fmt.Fprintln(w, "发送的的查询字符串是，", r.URL.RawQuery)
	// 查询指定字符的键值
	user := r.URL.Query().Get("user")
	fmt.Fprintln(w, "查询user", user)

	fmt.Fprintln(w, "请求头中的信息", r.Header)
}

func main() {

	http.HandleFunc("/hello", handler)

	http.ListenAndServe(":8080", nil)
}
