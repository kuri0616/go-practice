package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, world!\n")
	}

	postArticleHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Posting Article...\n")
	}

	getArticleListHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Article List\n")
	}

	getArticleHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Article No.1\n")
	}

	postArticleNiceHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Posting Nice...\n")
	}

	postCommentHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Posting Comment...\n")
	}

	http.HandleFunc("/article", postArticleHandler)
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/article/list", getArticleListHandler)
	http.HandleFunc("/article/1", getArticleHandler)
	http.HandleFunc("/article/nice", postArticleNiceHandler)
	http.HandleFunc("/comment", postCommentHandler)

	log.Println("server start at port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
