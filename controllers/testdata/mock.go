package testdata

import "github.com/yourname/reponame/models"

type serviceMock struct{}

func NewServiceMock() *serviceMock {
	return &serviceMock{}
}

func (s *serviceMock) PostArticleService(article models.Article) (models.Article, error) {
	return articleTestData[1], nil
}

func (s *serviceMock) GetArticleService(id int) (models.Article, error) {
	return articleTestData[0], nil
}

func (s *serviceMock) ArticleListService(page int) ([]models.Article, error) {
	return articleTestData, nil
}

func (s *serviceMock) PostNiceService(id int) (models.Article, error) {
	return articleTestData[0], nil
}
