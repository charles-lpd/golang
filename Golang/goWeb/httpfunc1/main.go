package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// 创建处理器函数
func handler(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		http.Error(w, "Method 错误", http.StatusMethodNotAllowed)
		return
	}

	user := User{
		Username: "chalres",
		Password: "123456",
	}
	jsonBytes, err := json.Marshal(user)
	if err != nil {
		fmt.Println("json 数据失败", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)
}

func main() {
	// 当用户访问以下路由时，调用 handler 函数
	http.HandleFunc("/", handler)

	// 创建路由
	http.ListenAndServe(":8085", nil)
}
