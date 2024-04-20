package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/yourname/reponame/models"
	"io"
	"net/http"
	"strconv"
)

func HelloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, world!\n")
}

func PostArticleHandler(w http.ResponseWriter, req *http.Request) {
	article := models.Article1
	jsonData, err := json.Marshal(article)
	if err != nil {
		http.Error(w, "failed to encode json", http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)

}

func GetArticleListHandler(w http.ResponseWriter, req *http.Request) {
	articles := []models.Article{models.Article1, models.Article2}
	jsonData, err := json.Marshal(articles)
	if err != nil {
		http.Error(w, "failed to encode json", http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)
}

func GetArticleHandler(w http.ResponseWriter, req *http.Request) {
	articleId, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		http.Error(w, "Invalid query parameter", http.StatusBadRequest)
		return
	}
	article := models.Article1
	article.ID = articleId
	jsonData, err := json.Marshal(article)
	if err != nil {
		http.Error(w, "failed to encode json", http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)

}

func PostArticleNiceHandler(w http.ResponseWriter, req *http.Request) {
	article := models.Article1
	article.NiceNum++
	jsonData, err := json.Marshal(article)
	if err != nil {
		http.Error(w, "failed to encode json", http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)
}

func PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	comment := models.Comment1
	jsonData, err := json.Marshal(comment)
	if err != nil {
		http.Error(w, "failed to encode json", http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)
}
