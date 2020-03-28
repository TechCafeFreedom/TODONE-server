package mysql

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func CreateSQLInstance() *sql.DB {
	dbuser := os.Getenv("MYSQL_USER")
	dbpassword := os.Getenv("MYSQL_PASSWORD")
	dbhost := os.Getenv("MYSQL_HOST")
	dbname := os.Getenv("MYSQL_DATABASE")

	dataSource := fmt.Sprintf("%s:%s@%s/%s?parseTime=true", dbuser, dbpassword, dbhost, dbname)

	db, err := sql.Open("mysql", dataSource)
	if err != nil {
		panic(err.Error())
	}
	return db
}
