package controllers

import (
	"encoding/json"
	"github.com/yourname/reponame/controllers/services"
	"github.com/yourname/reponame/models"
	"net/http"
)

type CommentController struct {
	services services.CommentServicer
}

func NewCommentController(s services.CommentServicer) *CommentController {
	return &CommentController{services: s}
}

func (c *CommentController) PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	var reqComment models.Comment
	if err := json.NewDecoder(req.Body).Decode(&reqComment); err != nil {
		http.Error(w, "failed to decode json", http.StatusBadRequest)
		return
	}
	comment, err := c.services.PostCommentService(reqComment)
	if err != nil {
		http.Error(w, "failed to insert comment", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(comment); err != nil {
		http.Error(w, "failed to encode json", http.StatusInternalServerError)
		return
	}
}
