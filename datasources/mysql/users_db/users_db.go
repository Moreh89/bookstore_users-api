package users_db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var (
	Client *sql.DB
	// username = os.Getenv("mysql_users_username")
	// pass     = os.Getenv("mysql_users_password")
	// hostname = os.Getenv("mysql_users_hostname")
	// port     = os.Getenv("mysql_users_port")
	// schema   = os.Getenv("mysql_users_schema")
	username = "root"
	pass     = "123"
	hostname = "127.0.0.1"
	port     = "3660"
	schema   = "users_db"
)

func init() {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s",
		username, pass, hostname, schema)
	var err error
	Client, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
	if err = Client.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("database successfully configured")
}
