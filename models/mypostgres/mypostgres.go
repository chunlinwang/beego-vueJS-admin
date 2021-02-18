package mypostgres

import (
	"database/sql"

	_ "github.com/lib/pq" // import postgres driver.
)

var session *sql.DB

// Conn return mysql connection.
func Conn() *sql.DB {
	return session
}

//func init() {
//	url := beego.AppConfig.String("postgres::url")
//
//	db, err := sql.Open("postgres", url)
//	if err != nil {
//		panic(err)
//	}
//
//	if err := db.Ping(); err != nil {
//		panic(err)
//	}
//
//	session = db
//}
