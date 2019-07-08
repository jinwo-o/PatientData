package driver

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// DB ...
type DB struct {
	SQL *sql.DB
}

// DBConn ...
var dbConn = &DB{}


// ConnectSQL ...
func ConnectSQL(host, port, uname, pass, dbname string) (*DB, error) {
	// d, err := sql.Open("mysql", "root:root1234@tcp(127.0.0.1)/patients")
	d, err := sql.Open("mysql", "docker:docker@tcp(db:3306)/patients")

	if err != nil {
		panic(err)
	}
	dbConn.SQL = d
	return dbConn, err
}

