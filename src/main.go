package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql" // 必须这样引入
	"github.com/jmoiron/sqlx"
)

type User = struct {
	Email    string
	Account  string
	Nickname string
	Password string
	Status   int8
}

// ResInfo 表示的是返回的信息
type ResInfo = struct {
	status  int8
	message string
	data    string
}

func main() {
	database, err := sqlx.Open("mysql", "root:123457986@tcp(127.0.0.1:3306)/questionnaire_sql")
	if err != nil {
		return
	}
	fmt.Println(database)
	var users []User
	// 查询具有某个特定account的 user
	err = database.Select(&users, "select * from user where account = ?", "admin")
	if err != nil {
		fmt.Println("err", err)
		return
	}
	fmt.Println(users)
	// 登陆，通过account和password来判定
	var data []User
	err = database.Select(&data, "select * from user where account = ? and password = ?", "admin", "admin")
	if err != nil {
		fmt.Println("login err -> ", err)
		return
	}
	fmt.Println(len(data)) // 通过长度是否为0可以判断是否存在
	// 注册
	_, err = database.Exec("insert into user(account, password, nickname, email, status)values(?,?,?,?,?)", "admin_test", "admin_test", "nickname", "123@qq.com", 1)

	if err != nil {
		fmt.Println("register err -> ", err)
		//return
	}

	// 删除
	_, err = database.Exec("delete from user where account = ?", "admin_test")
	if err != nil {
		fmt.Println("del err -> ", err)
		// return
	}
}
