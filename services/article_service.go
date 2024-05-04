package services

import (
	"github.com/yourname/reponame/models"
	"github.com/yourname/reponame/repositories"
)

func (s *MyAppService) GetArticleService(articleID int) (models.Article, error) {
	article, err := repositories.SelectArticle(s.db, articleID)
	if err != nil {
		return models.Article{}, err
	}
	commentList, err := repositories.SelectCommentList(s.db, articleID)
	if err != nil {
		return models.Article{}, err
	}
	article.CommentList = append(article.CommentList, commentList...)
	return article, nil
}

func (s *MyAppService) PostArticleService(article models.Article) (models.Article, error) {
	newArticle, err := repositories.InsertArticle(s.db, article)
	if err != nil {
		return models.Article{}, err
	}
	return newArticle, nil
}

func (s *MyAppService) ArticleListHandler(page int) ([]models.Article, error) {
	articleList, err := repositories.SelectArticleList(s.db, page)
	if err != nil {
		return nil, err
	}
	return articleList, nil
}

func (s *MyAppService) PostNiceService(articleID int) (models.Article, error) {
	article, err := repositories.UpdateNice(s.db, articleID)
	if err != nil {
		return models.Article{}, err
	}
	return article, nil
}
