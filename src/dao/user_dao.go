package dao

// 形成一个良好的习惯，写完的每一个函数都需要验证函数的正确性

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"questionnaireGo/src/entity"
)

// LoginRight 登陆，通过account和password来判定
func LoginRight(account string, password string) bool {
	var data []entity.User
	err := GetDBConnect().Select(&data, "select * from user where account = ? and password = ?", account, password)
	if err != nil || len(data) == 0 {
		return false
	}
	return true
}

// Register 注册
func Register(user entity.User) bool {
	_, err := GetDBConnect().Exec("insert into user(account, password, nickname, email, status)values(?,?,?,?,?)",
		user.Account, user.Password, user.Nickname, user.Email, user.Status)
	if err != nil {
		//fmt.Println("register err -> ", err)
		return false
	}
	return true
}

// DelUserByAccount 根据传输来的account删除user
func DelUserByAccount(account string) bool {
	var result sql.Result
	var err error
	var rowAffect int64
	var db *sqlx.DB = GetDBConnect()
	result, err = db.Exec("delete from user where account = ?", account)
	if err != nil {
		return false
	}
	if rowAffect, _ = result.RowsAffected(); rowAffect == 0 {
		//fmt.Println("del err -> ", err)
		return false
	}
	// 问卷和模板归admin所有的所有
	//db.Exec("delete from questionnaire where account = ?", account)
	//db.Exec("delete from template where account = ?", account)
	//管理员的账号前缀为admin
	db.Exec("update questionnaire set account = ? where account = ?", "admin", account)
	db.Exec("update template set account = ? where account = ?", "admin", account)
	return true
}

// ModifyUserStatus 修改user的status
func ModifyUserStatus(account string, status int8) bool {
	var result sql.Result
	var err error
	var rowAffect int64
	result, err = GetDBConnect().Exec("update user set status = ? where account = ?", status, account)
	if err != nil {
		return false
	}
	if rowAffect, _ = result.RowsAffected(); rowAffect == 0 {
		//fmt.Println("ModifyUserStatus err -> ", err)
		return false
	}
	return true
}

// ModifyUserAllInfo  修改user的info,此时需要的数据是原来的account以及最新的user对象
func ModifyUserAllInfo(account string, user entity.User) bool {
	_, err := GetDBConnect().Exec("update user set account = ?, password = ?, nickname = ?, email = ?  where account = ?",
		user.Account, user.Password, user.Nickname, user.Email, account)
	if err != nil {
		//fmt.Println("ModifyUserStatus err -> ", err)
		return false
	}
	return true
}

// GetUserInfoById 根据account查询user
func GetUserInfoById(account string) []entity.User {
	var users []entity.User
	err := GetDBConnect().Select(&users, "select * from user where account = ?", account)
	if err != nil {
		return nil
	}
	return users
}
