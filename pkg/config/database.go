package config

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

func GetDatabaseConnection() *sql.DB {
	config := mysql.Config{
		User:                 db_user,
		Passwd:               db_pwd,
		Net:                  "tcp",
		Addr:                 db_host + ":" + db_port,
		DBName:               db_name,
		AllowNativePasswords: true,
	}

	db, err := sql.Open("mysql", config.FormatDSN())
	if err != nil {
		log.Fatalf("Can't connect to database: %s", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Database not connected: %s", err)
	}

	return db
}
