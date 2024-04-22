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
	var reqArticle models.Article
	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
		http.Error(w, "failed to decode json", http.StatusBadRequest)
		return
	}

	article := reqArticle
	json.NewEncoder(w).Encode(article)
}

func GetArticleListHandler(w http.ResponseWriter, req *http.Request) {
	articles := []models.Article{models.Article1, models.Article2}
	json.NewEncoder(w).Encode(articles)
}

func GetArticleHandler(w http.ResponseWriter, req *http.Request) {
	articleId, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
	}
	article := models.Article1
	article.ID = articleId
	json.NewEncoder(w).Encode(article)
}

func PostArticleNiceHandler(w http.ResponseWriter, req *http.Request) {
	article := models.Article1
	article.NiceNum++
	json.NewEncoder(w).Encode(article)
}

func PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	var reqComment models.Comment
	if err := json.NewDecoder(req.Body).Decode(&reqComment); err != nil {
		http.Error(w, "failed to decode json", http.StatusBadRequest)
		return
	}
	comment := reqComment
	json.NewEncoder(w).Encode(comment)

}
