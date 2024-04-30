package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/yourname/reponame/controllers"
	"github.com/yourname/reponame/services"
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

	ser := services.NewMyAppService(db)
	con := controllers.NewMyAppController(ser)
	r := mux.NewRouter()

	r.HandleFunc("/article", con.PostArticleHandler).Methods(http.MethodPost)
	r.HandleFunc("/article/list", con.GetArticleListHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/{id:[0-9]+}", con.GetArticleHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/nice", con.PostArticleNiceHandler).Methods(http.MethodPost)
	r.HandleFunc("/comment", con.PostCommentHandler).Methods(http.MethodPost)

	log.Println("server start at port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
