package db

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var (
	config = InitConfiguration("root", "", "localhost", "3306", "go-lang-database")
	dns    = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.USERNAME, config.PASSWORD, config.HOST, config.PORT, config.DB_NAME)
)

func SetupDatabase() (*sql.DB, error) {

	db, err := sql.Open("mysql", dns)
	if err != nil {
		return db, err
	}

	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db, nil
}
