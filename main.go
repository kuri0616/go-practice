package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/yourname/reponame/api"
	"log"
	"net/http"
	"os"
)

func main() {
	var (
		dbUser     = os.Getenv("DB_USER")
		dbPass     = os.Getenv("DB_PASS")
		dbDatabase = os.Getenv("DB_NAME")
		dbConn     = fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPass, dbDatabase)
	)

	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		log.Fatal(err)
	}

	r := api.NewRouter(db)
	log.Println("server start at port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
