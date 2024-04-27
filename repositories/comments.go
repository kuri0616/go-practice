package repositories

import (
	"database/sql"
	"github.com/yourname/reponame/models"
)

func InsertComment(db *sql.DB, comment models.Comment) (models.Comment, error) {
	const sqlStr = `
insert into comments (article_id, message, created_at)
values (?, ?, now());
`
	var newComment models.Comment
	result, err := db.Exec(sqlStr, comment.ArticleID, comment.Message)
	if err != nil {
		return newComment, err
	}
	id, _ := result.LastInsertId()
	newComment.CommentID = int(id)
	newComment.ArticleID, newComment.Message = comment.ArticleID, comment.Message

	return newComment, nil
}

func SelectCommentList(db *sql.DB, articleID int) ([]models.Comment, error) {
	const sqlStr = `
select comment_id, article_id, message, created_at
from comments
where article_id = ?;
`
	rows, err := db.Query(sqlStr, articleID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	getCommentList := make([]models.Comment, 0)

	for rows.Next() {
		var comment models.Comment
		var createdTime sql.NullTime
		err := rows.Scan(&comment.CommentID, &comment.ArticleID, &comment.Message, &createdTime)
		if createdTime.Valid {
			comment.CreatedAt = createdTime.Time
		}
		if err != nil {
			return []models.Comment{}, err
		}
		getCommentList = append(getCommentList, comment)
	}
	return getCommentList, nil
}
