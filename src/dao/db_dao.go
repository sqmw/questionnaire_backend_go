package dao

import (
	"github.com/jmoiron/sqlx"
	"sync"
)
import _ "github.com/go-sql-driver/mysql" // 必须这样引入

var db *sqlx.DB = nil

func GetDBConnect() *sqlx.DB {
	// 在这里加上锁，防止死锁的现象发生
	locker := sync.Mutex{}
	locker.Lock()
	var err error
	if db == nil {
		db, err = sqlx.Connect("mysql", "root:123457986@tcp(127.0.0.1:3306)/questionnaire_sql")
		db.SetMaxIdleConns(20)
		db.SetConnMaxIdleTime(5)
		if err != nil {
			return nil
		}
	}
	locker.Unlock()
	return db
}
