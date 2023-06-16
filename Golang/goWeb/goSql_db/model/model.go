package model

import (
	"fmt"
	"golang/goWeb/goSql_db/utils"
)

type User struct {
	Username string
	Password string
}

func (user *User) AddUser() error {

	// 1. 写 sql 语句
	sqlStr := "insert into users(username,password) values(?,?)"

	// 2. 预编译
	inStmt, err := utils.Db.Prepare(sqlStr)
	if err != nil {
		fmt.Println("预编译出现异常", err)
		return err
	}

	defer utils.Db.Close()

	// 执行

	_, err2 := inStmt.Exec("charles_dong", "123123")

	if err2 != nil {
		fmt.Println("执行出现异常,", err2)
		return err2
	}
	return nil
}

func (user *User) AddUser2() error {
	// 1. 写 sql 语句
	sqlStr := "insert into users(username,password) values(?,?)"

	//执行

	_, err := utils.Db.Exec(sqlStr, "qiong", "000000")
	defer utils.Db.Close()
	if err != nil {
		fmt.Println("执行出现异常", err)
		return err
	}

	return nil
}

// GetQueryRow() 查询返回一行的结果
func (user *User) GetQueryRow() (*User, error) {

	// sql 语句
	sqlStr := "select id,username,password,status from users where username = ?"

	// 执行
	row := utils.Db.QueryRow(sqlStr, user.Username)
	var id int
	var username string
	var password string
	var status string
	err := row.Scan(&id,&username, &password, &status)
	if err != nil {
		return nil, err
	}
	fmt.Println(id)
	fmt.Println(status)
	u := &User{
		Username: username,
		Password: password,
		
	}
	return u, nil
}
