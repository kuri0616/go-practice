package testdata

import "github.com/yourname/reponame/models"

var articleTestData = []models.Article{
	models.Article{
		ID:          1,
		Title:       "Title 1",
		Content:     "Content 1",
		UserName:    "User 1",
		NiceNum:     1,
		CommentList: commentTestData,
	},
	models.Article{
		ID:       2,
		Title:    "Title 2",
		Content:  "Content 2",
		UserName: "User 2",
		NiceNum:  2,
	},
}

var commentTestData = []models.Comment{
	models.Comment{
		CommentID: 1,
		ArticleID: 1,
		Message:   "Comment 1",
	},
	models.Comment{
		CommentID: 2,
		ArticleID: 1,
		Message:   "Comment 2",
	},
}
