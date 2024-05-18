package services

import (
	"database/sql"
	"errors"
	"github.com/yourname/reponame/apperrors"
	"github.com/yourname/reponame/models"
	"github.com/yourname/reponame/repositories"
)

func (s *MyAppService) GetArticleService(articleID int) (models.Article, error) {
	article, err := repositories.SelectArticle(s.db, articleID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = apperrors.NoData.Wrap(err, "article not found")
			return models.Article{}, err
		}
		err = apperrors.GetDataFailed.Wrap(err, "failed to select article")
		return models.Article{}, err
	}
	commentList, err := repositories.SelectCommentList(s.db, articleID)
	if err != nil {
		err = apperrors.GetDataFailed.Wrap(err, "failed to select comment list")
		return models.Article{}, err
	}
	article.CommentList = append(article.CommentList, commentList...)
	return article, nil
}

func (s *MyAppService) PostArticleService(article models.Article) (models.Article, error) {
	newArticle, err := repositories.InsertArticle(s.db, article)
	if err != nil {
		err = apperrors.InsertDataFailed.Wrap(err, "failed to insert article")
		return models.Article{}, err
	}
	return newArticle, nil
}

func (s *MyAppService) ArticleListService(page int) ([]models.Article, error) {
	articleList, err := repositories.SelectArticleList(s.db, page)
	if err != nil {
		err = apperrors.GetDataFailed.Wrap(err, "failed to select article list")
		return nil, err
	}
	if len(articleList) == 0 {
		err = apperrors.NoData.Wrap(ErrNoData, "no data")
	}
	return articleList, nil
}

func (s *MyAppService) PostNiceService(articleID int) (models.Article, error) {
	article, err := repositories.UpdateNice(s.db, articleID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = apperrors.NoTargetData.Wrap(err, "does not exist article")
			return models.Article{}, err
		}
		err = apperrors.UpdateDataFailed.Wrap(err, "failed to update nice")
		return models.Article{}, err
	}
	return article, nil
}
