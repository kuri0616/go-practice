package repositories_test

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/yourname/reponame/models"
	"github.com/yourname/reponame/repositories"
	"github.com/yourname/reponame/repositories/testdata"
	"testing"
)

func TestSelectArticle(t *testing.T) {
	tests := []struct {
		testTitle string
		expected  models.Article
	}{
		{
			testTitle: "testTitle1",
			expected:  testdata.ArticleTestData[0],
		},
		{
			testTitle: "testTitle2",
			expected:  testdata.ArticleTestData[1],
		},
	}

	for _, test := range tests {
		t.Run(test.testTitle, func(t *testing.T) {
			got, err := repositories.SelectArticle(testDB, test.expected.ID)
			if err != nil {
				t.Fatal(err)
			}
			if got.ID != test.expected.ID {
				t.Errorf("get %v, but want %v", got.ID, test.expected.ID)
			}
			if got.Title != test.expected.Title {
				t.Errorf("get %v, but want %v", got.Title, test.expected.Title)
			}
			if got.Content != test.expected.Content {
				t.Errorf("get %v, but want %v", got.Content, test.expected.Content)
			}
			if got.UserName != test.expected.UserName {
				t.Errorf("get %v, but want %v", got.UserName, test.expected.UserName)
			}
		})
	}
}

func TestSelectArticleList(t *testing.T) {
	expectedNum := len(testdata.ArticleTestData)
	got, err := repositories.SelectArticleList(testDB, 1)
	if err != nil {
		t.Error(err)
	}
	if len(got) != expectedNum {
		t.Errorf("len(got) = %v, want %v", len(got), expectedNum)
	}
}

func TestInsertArticle(t *testing.T) {
	article := models.Article{
		Title:    "yeah",
		Content:  "yeah, world!",
		UserName: "Charlie",
	}

	expectedArticleNum := 3
	newArticle, err := repositories.InsertArticle(testDB, article)
	if err != nil {
		t.Error(err)
	}
	if newArticle.ID != expectedArticleNum {
		t.Errorf("newArticle.ID = %v, want %v", newArticle.ID, expectedArticleNum)
	}

	t.Cleanup(func() {
		const sqlStr = `
		delete from articles
		where article_id = ?;
	`
		_, err := testDB.Exec(sqlStr, newArticle.ID)

		if err != nil {
			t.Error(err)
		}
	})
}

func TestUpdateNice(t *testing.T) {
	expectedNiceNum := testdata.ArticleTestData[0].NiceNum + 1
	articleID := testdata.ArticleTestData[0].ID
	err := repositories.UpdateNice(testDB, articleID)
	if err != nil {
		t.Error(err)
	}

	got, _ := repositories.SelectArticle(testDB, articleID)
	if expectedNiceNum != got.NiceNum {
		t.Errorf("expectedNiceNum = %v, want %v", expectedNiceNum, got.NiceNum)
	}

	t.Cleanup(func() {
		beforeNiceNum := testdata.ArticleTestData[0].NiceNum - 1
		const sqlStr = `
		update articles
		set nice = ?
		where article_id = ?;
	`
		_, err := testDB.Exec(sqlStr, beforeNiceNum, articleID)

		if err != nil {
			t.Error(err)
		}
	})
}

func TestInsertComment(t *testing.T) {
	comment := models.Comment{
		ArticleID: 1,
		Message:   "Hello, world!",
	}
	expectedCommentNum := 3

	newComment, err := repositories.InsertComment(testDB, comment)
	if err != nil {
		t.Error(err)
	}
	if newComment.CommentID != expectedCommentNum {
		t.Errorf("newComment.CommentID = %v, want %v", newComment.CommentID, expectedCommentNum)
	}
	t.Cleanup(func() {
		const sqlStr = `
		delete from comments
		where comment_id = ?;
	`
		_, err := testDB.Exec(sqlStr, newComment.CommentID)
		if err != nil {
			t.Error(err)
		}
	})
}
