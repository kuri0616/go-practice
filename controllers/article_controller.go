package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/yourname/reponame/apperrors"
	"github.com/yourname/reponame/controllers/services"
	"github.com/yourname/reponame/models"
	"net/http"
	"strconv"
)

type ArticleController struct {
	service services.ArticleServicer
}

func NewArticleController(s services.ArticleServicer) *ArticleController {
	return &ArticleController{service: s}
}

func (c *ArticleController) PostArticleHandler(w http.ResponseWriter, req *http.Request) {
	var reqArticle models.Article
	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
		err = apperrors.ReqBodyDecodeFailed.Wrap(err, "failed to decode json")
		http.Error(w, "failed to decode json", http.StatusBadRequest)
		return
	}
	resArticle, err := c.service.PostArticleService(reqArticle)
	if err != nil {
		err = apperrors.InsertDataFailed.Wrap(err, "failed to insert article")
		http.Error(w, "failed to insert article", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(resArticle); err != nil {
		err = apperrors.ReqBodyEncodeFailed.Wrap(err, "failed to encode json")
		http.Error(w, "failed to encode json", http.StatusInternalServerError)
		return
	}
}

func (c *ArticleController) GetArticleListHandler(w http.ResponseWriter, req *http.Request) {
	queryMap := req.URL.Query()
	if p, OK := queryMap["page"]; OK && len(p) > 0 {
		page, err := strconv.Atoi(p[0])
		if err != nil {
			err = apperrors.BadParam.Wrap(err, "invalid page number")
			http.Error(w, "invalid page number", http.StatusBadRequest)
			return
		}
		articleList, err := c.service.ArticleListHandler(page)
		if err != nil {
			err = apperrors.GetDataFailed.Wrap(err, "failed to get article list")
			http.Error(w, "failed to get article list", http.StatusInternalServerError)
			return
		}
		if err := json.NewEncoder(w).Encode(articleList); err != nil {
			err = apperrors.ReqBodyEncodeFailed.Wrap(err, "failed to encode json")
			http.Error(w, "failed to encode json", http.StatusInternalServerError)
			return
		}
	}
}

func (c *ArticleController) GetArticleHandler(w http.ResponseWriter, req *http.Request) {
	articleId, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		err = apperrors.BadParam.Wrap(err, "invalid article id")
		http.Error(w, "invalid article id", http.StatusBadRequest)
		return
	}
	article, err := c.service.GetArticleService(articleId)
	if err != nil {
		err = apperrors.GetDataFailed.Wrap(err, "failed to get article")
		http.Error(w, "failed to get article", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(article); err != nil {
		err = apperrors.ReqBodyEncodeFailed.Wrap(err, "failed to encode json")
		http.Error(w, "failed to encode json", http.StatusInternalServerError)
		return
	}
}

func (c *ArticleController) PostArticleNiceHandler(w http.ResponseWriter, req *http.Request) {
	articleId, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		err = apperrors.BadParam.Wrap(err, "invalid article id")
		http.Error(w, "invalid article id", http.StatusBadRequest)
		return
	}
	article, err := c.service.PostNiceService(articleId)
	if err != nil {
		err = apperrors.UpdateDataFailed.Wrap(err, "failed to update nice")
		http.Error(w, "failed to update nice", http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(article); err != nil {
		err = apperrors.ReqBodyEncodeFailed.Wrap(err, "failed to encode json")
		http.Error(w, "failed to encode json", http.StatusInternalServerError)
		return
	}
}
