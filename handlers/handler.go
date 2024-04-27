package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/yourname/reponame/models"
	"github.com/yourname/reponame/services"
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
	resArticle, err := services.PostArticleSerrvice(reqArticle)
	if err != nil {
		http.Error(w, "failed to insert article", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(resArticle); err != nil {
		http.Error(w, "failed to encode json", http.StatusInternalServerError)
		return
	}
}

func GetArticleListHandler(w http.ResponseWriter, req *http.Request) {
	queryMap := req.URL.Query()
	if p, OK := queryMap["page"]; OK && len(p) > 0 {
		page, err := strconv.Atoi(p[0])
		if err != nil {
			http.Error(w, "invalid page number", http.StatusBadRequest)
			return
		}
		articleList, err := services.ArticleListHandler(page)
		if err != nil {
			http.Error(w, "failed to get article list", http.StatusInternalServerError)
			return
		}
		if err := json.NewEncoder(w).Encode(articleList); err != nil {
			http.Error(w, "failed to encode json", http.StatusInternalServerError)
			return
		}
	}
}

func GetArticleHandler(w http.ResponseWriter, req *http.Request) {
	articleId, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		http.Error(w, "invalid article id", http.StatusBadRequest)
		return
	}
	article, err := services.GetArticleService(articleId)
	if err != nil {
		http.Error(w, "failed to get article", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(article); err != nil {
		http.Error(w, "failed to encode json", http.StatusInternalServerError)
		return
	}
}

func PostArticleNiceHandler(w http.ResponseWriter, req *http.Request) {
	articleId, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		http.Error(w, "invalid article id", http.StatusBadRequest)
		return
	}
	article, err := services.PostNiceService(articleId)
	if err != nil {
		http.Error(w, "failed to update nice", http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(article); err != nil {
		http.Error(w, "failed to encode json", http.StatusInternalServerError)
		return
	}
}

func PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	var reqComment models.Comment
	if err := json.NewDecoder(req.Body).Decode(&reqComment); err != nil {
		http.Error(w, "failed to decode json", http.StatusBadRequest)
		return
	}
	comment, err := services.PostCommentService(reqComment)
	if err != nil {
		http.Error(w, "failed to insert comment", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(comment); err != nil {
		http.Error(w, "failed to encode json", http.StatusInternalServerError)
		return
	}
}
