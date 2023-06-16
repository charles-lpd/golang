package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/thinkerou/favicon"
)

func Hello(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"msg": "你好",
	})
}

func GetURLQuery(ctx *gin.Context) {
	// user?user=charles&id=
	ctx.DefaultQuery("id", "qiong")           // ""
	ctx.DefaultQuery("lastname", "qiong")     // qiong
	user := ctx.DefaultQuery("user", "qiong") //charles

	fmt.Println(user)
	ctx.JSON(200, gin.H{
		"user": user,
	})
}

// 获取地址上 param 的值
func GetURLParams(ctx *gin.Context) {
	// userInfo/:id/:username
	id := ctx.Param("id")
	username := ctx.Param("username")
	fmt.Println(id, username, "12312312")
	ctx.JSON(200, gin.H{
		"id":       id,
		"username": username,
	})
}

type Account struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Msg     string `json:"msg"`
	Address string `json:"address"`
}

// 前端给后端传递 json 数据
func PostJsonFormData(ctx *gin.Context) {
	// 获取 POST 接口 form-data 或 x-www-form-urlencoded 传递的的值
	name := ctx.PostForm("name")                // 取不到返回空字符串
	ageStr := ctx.DefaultPostForm("age", "100") // 取不到返回默认值 100
	// 因为返回 stirng类型 所以需要将 age 转为 int 类型
	age, err := strconv.Atoi(ageStr)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"status": 4001,
			"error":  "err_request_params",
			"msg":    "age incorrect",
		})
		return
	}
	msg, msg_err := ctx.GetPostForm("msg")
	if !msg_err { // 取到 true 未取到 false
		ctx.JSON(http.StatusBadGateway, gin.H{
			"status": 4001,
			"error":  "err_request_params",
			"msg":    "msg incorrect",
		})
		return
	}
	address, add_err := ctx.GetPostForm("address")
	if !add_err {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"status": 4001,
			"error":  "err_request_params",
			"msg":    "address incorrect",
		})
		return
	}

	m := Account{
		name, age, msg, address,
	}

	fmt.Println(name)

	ctx.JSON(http.StatusOK, m)
}

func PostJsonRaw(ctx *gin.Context) {
	data, err := ctx.GetRawData()
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"status": 4001,
			"error":  err.Error(),
		})
	}
	var m Account
	marshal_err := json.Unmarshal(data, &m)
	if marshal_err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"status": 4001,
			"error":  marshal_err.Error(),
		})
	}
	ctx.JSON(http.StatusOK, m)
}
func main() {
	// 创建一个服务
	ginServer := gin.Default()

	// 配置前端 网页标题头像
	ginServer.Use(favicon.New("./favicon.ico"))

	// 访问地址，处理我们的请求， Request Response
	ginServer.GET("/hello", Hello)

	// 获取地址栏参数 // user?user=charles&id=666
	ginServer.GET("/user", GetURLQuery)

	// 获取地址栏参数 // userInfo/1/charles
	ginServer.GET("/userinfo/:id/:username", GetURLParams)

	// 前端给后端传递 json form-data 数据
	ginServer.POST("/jsonformdata", PostJsonFormData)
	// 前端给后端传递 json raw 数据
	ginServer.POST("/jsonraw", PostJsonRaw)
	// 服务器端口
	ginServer.Run(":8082")

}
