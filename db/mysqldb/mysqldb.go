package mysqldb

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

// 打开mysql 数据库连接
func OpenDB(path string) (*sql.DB, error) {
	DB, err := sql.Open("mysql", path)
	if err != nil {
		panic(err)
	}
	DB.SetConnMaxLifetime(time.Minute * 3)
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(10)
	return DB, err
}
