package repositories_test

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/yourname/reponame/models"
	"github.com/yourname/reponame/repositories"
	"testing"
)

func TestSelectArticle(t *testing.T) {
	dbUser := "docker"
	dbPassword := "docker"
	dbDatabase := "sampledb"
	dbConn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)

	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	tests := []struct {
		testTitle string
		expected  models.Article
	}{
		{
			testTitle: "subtest1",
			expected: models.Article{
				ID:       1,
				Title:    "Hello",
				Content:  "Hello, world!",
				UserName: "Alice",
				NiceNum:  4,
			},
		}, {
			testTitle: "subtest2",
			expected: models.Article{
				ID:       2,
				Title:    "Goodbye",
				Content:  "Goodbye, world!",
				UserName: "Bob",
				NiceNum:  0,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.testTitle, func(t *testing.T) {
			got, err := repositories.SelectArticle(db, test.expected.ID)
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

	//got, err := repositories.SelectArticle(db, expected.ID)
	//if err != nil {
	//	t.Fatal(err)
	//}
	//
	//if got.ID != expected.ID {
	//	t.Errorf("got %v, expected %v", got.ID, expected.ID)
	//}
	//if got.Title != expected.Title {
	//	t.Errorf("got %v, expected %v", got.Title, expected.Title)
	//}
	//if got.Content != expected.Content {
	//	t.Errorf("got %v, expected %v", got.Content, expected.Content)
	//}
	//if got.UserName != expected.UserName {
	//	t.Errorf("got %v, expected %v", got.UserName, expected.UserName)
	//}
	//if got.NiceNum != expected.NiceNum {
	//	t.Errorf("got %v, expected %v", got.NiceNum, expected.NiceNum)
	//}
}
