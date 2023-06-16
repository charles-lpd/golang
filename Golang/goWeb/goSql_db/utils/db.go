package utils

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var (
	Db *sql.DB
	err error
)

func init(){
	// 用户名:密码@tcp(主机:端口)/数据库名
	Db, err = sql.Open("mysql", "root:Lpeidong0916..@tcp(localhost:3306)/my_db_01")

	if err != nil {
		panic(err.Error())
	}

	// 测试连接是否成功
	err = Db.Ping()

	if err != nil {
		panic(err.Error())
	}
	fmt.Println("成功连接到数据库")
}
