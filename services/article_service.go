package services

import (
	"github.com/yourname/reponame/models"
	"github.com/yourname/reponame/repositories"
)

func GetArticleService(articleID int) (models.Article, error) {
	db, err := connectDB()
	if err != nil {
		return models.Article{}, err
	}
	defer db.Close()

	article, err := repositories.SelectArticle(db, articleID)
	if err != nil {
		return models.Article{}, err
	}
	commentList, err := repositories.SelectCommentList(db, articleID)
	if err != nil {
		return models.Article{}, err
	}
	article.CommentList = append(article.CommentList, commentList...)
	return article, nil
}

func PostArticleSerrvice(article models.Article) (models.Article, error) {
	db, err := connectDB()
	if err != nil {
		return models.Article{}, err
	}
	defer db.Close()

	newArticle, err := repositories.InsertArticle(db, article)
	if err != nil {
		return models.Article{}, err
	}
	return newArticle, nil
}

func ArticleListHandler(page int) ([]models.Article, error) {
	db, err := connectDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	articleList, err := repositories.SelectArticleList(db, page)
	if err != nil {
		return nil, err
	}
	return articleList, nil
}

func PostNiceService(articleID int) error {
	db, err := connectDB()
	if err != nil {
		return err
	}
	defer db.Close()

	err = repositories.UpdateNice(db, article)
	if err != nil {
		return err
	}
	return nil
}
