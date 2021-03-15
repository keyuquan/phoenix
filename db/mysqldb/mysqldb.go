package mysqldb

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strings"
	"time"
)

// 打开mysql 数据库连接
func OpenDB(userName string, password string, ip string, port string, dbName string) (*sql.DB, error) {
	path := strings.Join([]string{userName, ":", password, "@tcp(", ip, ":", port, ")/", dbName, "?tls=skip-verify&autocommit=true"}, "")
	DB, err := sql.Open("mysql", path)
	if err != nil {
		panic(err)
	}
	DB.SetConnMaxLifetime(time.Minute * 3)
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(10)
	if err := DB.Ping(); err != nil {
		fmt.Println("opon database fail")
	}
	fmt.Println("connnect success")
	return DB, err
}
