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
		apperrors.ErrorHandler(w, req, err)
		return
	}
	resArticle, err := c.service.PostArticleService(reqArticle)
	if err != nil {
		err = apperrors.InsertDataFailed.Wrap(err, "failed to insert article")
		apperrors.ErrorHandler(w, req, err)
		return
	}

	if err := json.NewEncoder(w).Encode(resArticle); err != nil {
		err = apperrors.ReqBodyEncodeFailed.Wrap(err, "failed to encode json")
		apperrors.ErrorHandler(w, req, err)
		return
	}
}

func (c *ArticleController) GetArticleListHandler(w http.ResponseWriter, req *http.Request) {
	queryMap := req.URL.Query()
	if p, OK := queryMap["page"]; OK && len(p) > 0 {
		page, err := strconv.Atoi(p[0])
		if err != nil {
			err = apperrors.BadParam.Wrap(err, "invalid page number")
			apperrors.ErrorHandler(w, req, err)
			return
		}
		articleList, err := c.service.ArticleListService(page)
		if err != nil {
			err = apperrors.GetDataFailed.Wrap(err, "failed to get article list")
			apperrors.ErrorHandler(w, req, err)
			return
		}
		if err := json.NewEncoder(w).Encode(articleList); err != nil {
			err = apperrors.ReqBodyEncodeFailed.Wrap(err, "failed to encode json")
			apperrors.ErrorHandler(w, req, err)
			return
		}
	}
}

func (c *ArticleController) GetArticleHandler(w http.ResponseWriter, req *http.Request) {
	articleId, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		err = apperrors.BadParam.Wrap(err, "invalid article id")
		apperrors.ErrorHandler(w, req, err)
		return
	}
	article, err := c.service.GetArticleService(articleId)
	if err != nil {
		err = apperrors.GetDataFailed.Wrap(err, "failed to get article")
		apperrors.ErrorHandler(w, req, err)
		return
	}

	if err := json.NewEncoder(w).Encode(article); err != nil {
		err = apperrors.ReqBodyEncodeFailed.Wrap(err, "failed to encode json")
		apperrors.ErrorHandler(w, req, err)
		return
	}
}

func (c *ArticleController) PostArticleNiceHandler(w http.ResponseWriter, req *http.Request) {
	articleId, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		err = apperrors.BadParam.Wrap(err, "invalid article id")
		apperrors.ErrorHandler(w, req, err)
		return
	}
	article, err := c.service.PostNiceService(articleId)
	if err != nil {
		apperrors.ErrorHandler(w, req, err)
		http.Error(w, "failed to update nice", http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(article); err != nil {
		err = apperrors.ReqBodyEncodeFailed.Wrap(err, "failed to encode json")
		apperrors.ErrorHandler(w, req, err)
		return
	}
}
