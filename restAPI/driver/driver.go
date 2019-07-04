package driver

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// DB ...
type DB struct {
	SQL *sql.DB
}

// DBConn ...
var dbConn = &DB{}

// Check for any idle connections
// func TestonBorrow(c redis.Conn, t time.Time) error {
// 	if time.Since(t) < time.Minute {
// 		return nil
// 	}
// 	_, err := c.Do("PING")
// 	return err
// }

// ConnectSQL ...
func ConnectSQL(host, port, uname, pass, dbname string) (*DB, error) {
	dbSource := fmt.Sprintf(
		// "root:%s@tcp(%s:%s)/%s?charset=utf8",
		"root:%s@tcp(%s:%s)/%s",
		pass,
		host,
		port,
		dbname,
	)
	d, err := sql.Open("mysql", dbSource)
	if err != nil {
		panic(err)
	}
	dbConn.SQL = d
	// this fix didnt work
	// dbConn.SQL.SetConnMaxLifetime(time.Minute * 5)
	// dbConn.SQL.SetMaxIdleConns(0)
	// dbConn.SQL.SetMaxOpenConns(5)
	return dbConn, err
}
