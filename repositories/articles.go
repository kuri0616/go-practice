package repositories

import (
	"database/sql"
	"github.com/yourname/reponame/models"
)

const (
	perPage = 5
)

func InsertArticle(db *sql.DB, article models.Article) (models.Article, error) {
	const sqlStr = `
insert into articles (title, contents, username, nice, created_at)
values (?, ?, ?, 0, now());
`
	var newArticle models.Article
	result, err := db.Exec(sqlStr, article.Title, article.Content, article.UserName)
	if err != nil {
		return newArticle, err
	}
	id, _ := result.LastInsertId()
	newArticle.ID = int(id)
	newArticle.Title, newArticle.Content, newArticle.UserName = article.Title, article.Content, article.UserName

	return newArticle, nil
}

func SelectArticleList(db *sql.DB, page int) ([]models.Article, error) {
	const sqlStr = `
select article_id,title,contents,username,nice,created_at
from articles
limit ? offset ?;
`
	rows, err := db.Query(sqlStr, perPage, (page-1)*perPage)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	getArticleList := make([]models.Article, 0)

	for rows.Next() {
		var article models.Article
		var createdTime sql.NullTime
		err := rows.Scan(&article.ID, &article.Title, &article.Content, &article.UserName, &article.NiceNum, createdTime)
		if createdTime.Valid {
			article.CreatedAt = createdTime.Time
		}
		if err != nil {
			return []models.Article{}, err
		}
		getArticleList = append(getArticleList, article)
	}
	return getArticleList, nil
}

func SelectArticle(db *sql.DB, articleID int) (models.Article, error) {
	const sqlStr = `
select article_id,title,contents,username,nice,created_at
from articles
where article_id = ?;
`
	row := db.QueryRow(sqlStr, articleID)
	var article models.Article
	var createdTime sql.NullTime
	err := row.Err()
	if err != nil {
		return models.Article{}, err
	}
	err = row.Scan(&article.ID, &article.Title, &article.Content, &article.UserName, &article.NiceNum, createdTime)
	if createdTime.Valid {
		article.CreatedAt = createdTime.Time
	}
	if err != nil {
		return models.Article{}, err
	}
	return article, nil
}

func UpdateNice(db *sql.DB, articleID int) error {
	const getArticleSql = `
select nice from articles where article_id = ?;
`
	const updateNiceSql = `
update articles
set nice = ?
where article_id = ?;
`

	tx, err := db.Begin()
	if err != nil {
		return err
	}
	row := tx.QueryRow(getArticleSql, articleID)

	if err := row.Err(); err != nil {
		tx.Rollback
		return err
	}
	var nicenum int

	err = row.Scan(&nicenum)
	if err != nil {
		tx.Rollback
		return err
	}
	_, err = tx.Exec(updateNiceSql, nicenum+1, articleID)
	if err != nil {
		tx.Rollback
		return err
	}
	if err := tx.Commit(); err != nil {
		tx.Rollback
		return err
	}
	return nil
}
