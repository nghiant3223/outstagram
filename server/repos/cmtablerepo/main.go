package cmtablerepo

import (
	"github.com/jinzhu/gorm"
	"outstagram/server/models"
)

type CommentableRepo struct {
	db *gorm.DB
}

func New(dbConnection *gorm.DB) *CommentableRepo {
	return &CommentableRepo{db: dbConnection}
}

func (r *CommentableRepo) GetCommentsByID(id uint) ([]models.Comment, error) {
	var commentable models.Commentable
	var comments []models.Comment

	if err := r.db.First(&commentable, id).Error; err != nil {
		return nil, err
	}

	r.db.Model(&commentable).Related(&comments)
	for i := 0; i < len(comments); i++ {
		r.db.Model(&comments[i]).Related(&comments[i].User)
		r.db.Model(&comments[i]).Related(&comments[i].Replies)
	}

	return comments, nil
}