package models

import "time"

var (
	Comment1 = Comment{
		CommentID: 1,
		ArticleID: 1,
		Message:   "This is the first comment.",
		CreatedAt: time.Now(),
	}
	Comment2 = Comment{
		CommentID: 2,
		ArticleID: 1,
		Message:   "This is the second comment.",
		CreatedAt: time.Now(),
	}
)

var (
	Article1 = Article{
		ID:          1,
		Title:       "first article",
		Content:     "This is the first article.",
		UserName:    "riku",
		NiceNum:     10,
		CreatedAt:   time.Now(),
		CommentList: []Comment{Comment1, Comment2},
	}
	Article2 = Article{
		ID:        2,
		Title:     "second article",
		Content:   "This is the second article.",
		UserName:  "riku",
		NiceNum:   20,
		CreatedAt: time.Now(),
	}
)
