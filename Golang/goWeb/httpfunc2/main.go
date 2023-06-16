package main

import (
	"fmt"
	"net/http"
)

type Myhandler struct{}

// 创建处理器处理请求
func (m *Myhandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "通过自己创建的处理器处理请求")
}
func main() {

	
	myhandler := Myhandler{}
	http.Handle("/myHandler", &myhandler)

	http.ListenAndServe(":8080", nil)
}
