package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/yourname/reponame/models"
	"github.com/yourname/reponame/services"
	"net/http"
	"strconv"
)

type MyAppController struct {
	service *services.MyAppService
}

func NewMyAppController(s *services.MyAppService) *MyAppController {
	return &MyAppController{service: s}
}

func (c *MyAppController) PostArticleHandler(w http.ResponseWriter, req *http.Request) {
	var reqArticle models.Article
	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
		http.Error(w, "failed to decode json", http.StatusBadRequest)
		return
	}
	resArticle, err := c.service.PostArticleService(reqArticle)
	if err != nil {
		http.Error(w, "failed to insert article", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(resArticle); err != nil {
		http.Error(w, "failed to encode json", http.StatusInternalServerError)
		return
	}
}

func (c *MyAppController) GetArticleListHandler(w http.ResponseWriter, req *http.Request) {
	queryMap := req.URL.Query()
	if p, OK := queryMap["page"]; OK && len(p) > 0 {
		page, err := strconv.Atoi(p[0])
		if err != nil {
			http.Error(w, "invalid page number", http.StatusBadRequest)
			return
		}
		articleList, err := c.service.ArticleListHandler(page)
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

func (c *MyAppController) GetArticleHandler(w http.ResponseWriter, req *http.Request) {
	articleId, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		http.Error(w, "invalid article id", http.StatusBadRequest)
		return
	}
	article, err := c.service.GetArticleService(articleId)
	if err != nil {
		http.Error(w, "failed to get article", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(article); err != nil {
		http.Error(w, "failed to encode json", http.StatusInternalServerError)
		return
	}
}

func (c *MyAppController) PostArticleNiceHandler(w http.ResponseWriter, req *http.Request) {
	articleId, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		http.Error(w, "invalid article id", http.StatusBadRequest)
		return
	}
	article, err := c.service.PostNiceService(articleId)
	if err != nil {
		http.Error(w, "failed to update nice", http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(article); err != nil {
		http.Error(w, "failed to encode json", http.StatusInternalServerError)
		return
	}
}

func (c *MyAppController) PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	var reqComment models.Comment
	if err := json.NewDecoder(req.Body).Decode(&reqComment); err != nil {
		http.Error(w, "failed to decode json", http.StatusBadRequest)
		return
	}
	comment, err := c.service.PostCommentService(reqComment)
	if err != nil {
		http.Error(w, "failed to insert comment", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(comment); err != nil {
		http.Error(w, "failed to encode json", http.StatusInternalServerError)
		return
	}
}
