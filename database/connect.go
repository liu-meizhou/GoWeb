//Package database 获取连接数据库的引用
package database

/*
  使用go查询postgres 数据库
*/

import (
    "fmt"
    "database/sql"
    
    //_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

const (
    host     = "193.112.207.30"
    port     = 5432
    user     = "postgres"
    password = "liu123456"
    dbname   = "postgres"
)

var db *sql.DB

//初始化数据库连接池
func init(){
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
        "password=%s dbname=%s sslmode=disable",
        host, port, user, password, dbname)
    db, _ = sql.Open("postgres", psqlInfo)
    db.SetMaxOpenConns(10)
    db.SetMaxIdleConns(8)
}

// GetConnect 单例模式获取连接
func GetConnect() (*sql.DB){
	return db
}
