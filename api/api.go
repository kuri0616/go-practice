package api

import (
	"database/sql"
	"github.com/gorilla/mux"
	"github.com/yourname/reponame/controllers"
	"github.com/yourname/reponame/services"
	"net/http"
)

func NewRouter(db *sql.DB) *mux.Router {
	ser := services.NewMyAppService(db)
	aCon := controllers.NewArticleController(ser)
	cCon := controllers.NewCommentController(ser)
	r := mux.NewRouter()
	r.HandleFunc("/article", aCon.PostArticleHandler).Methods(http.MethodPost)
	r.HandleFunc("/article/list", aCon.GetArticleListHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/{id:[0-9]+}", aCon.GetArticleHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/nice", aCon.PostArticleNiceHandler).Methods(http.MethodPost)
	r.HandleFunc("/comment", cCon.PostCommentHandler).Methods(http.MethodPost)
	return r
}
