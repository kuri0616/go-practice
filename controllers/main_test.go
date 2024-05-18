package controllers_test

import (
	"github.com/yourname/reponame/controllers"
	"github.com/yourname/reponame/controllers/testdata"
	"testing"
)

var aCon *controllers.ArticleController

func TestMain(m *testing.M) {
	ser := testdata.NewServiceMock()
	aCon = controllers.NewArticleController(ser)
	m.Run()
}
