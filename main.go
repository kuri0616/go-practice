package main

import (
	"github.com/yourname/reponame/handlers"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/article", handlers.PostArticleHandler)
	http.HandleFunc("/", handlers.HelloHandler)
	http.HandleFunc("/article/list", handlers.GetArticleListHandler)
	http.HandleFunc("/article/1", handlers.GetArticleHandler)
	http.HandleFunc("/article/nice", handlers.PostArticleNiceHandler)
	http.HandleFunc("/comment", handlers.PostCommentHandler)

	log.Println("server start at port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
