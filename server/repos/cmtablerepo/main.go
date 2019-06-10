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

func (r *CommentableRepo) GetComments(id uint) (*models.Commentable, error) {
	var commentable models.Commentable

	if err := r.db.First(&commentable, id).Error; err != nil {
		return nil, err
	}

	r.db.Model(&commentable).Related(&commentable.Comments)
	for i := 0; i < len(commentable.Comments); i++ {
		r.db.Model(&commentable.Comments[i]).Related(&commentable.Comments[i].User)
		r.db.Model(&commentable.Comments[i]).Related(&commentable.Comments[i].Replies)
	}

	commentable.CommentCount = r.GetCommentCount(id)
	return &commentable, nil
}

func (r *CommentableRepo) GetCommentsWithLimit(id uint, limit uint, offset uint) (*models.Commentable, error) {
	var commentable models.Commentable

	if err := r.db.First(&commentable, id).Error; err != nil {
		return nil, err
	}

	if err := r.db.Where("commentable_id = ?", commentable.ID).
		Offset(offset).
		Limit(limit).
		Find(&commentable.Comments).
		Error; err != nil {
		return nil, err
	}

	comments := commentable.Comments
	for i := 0; i < len(comments); i++ {
		r.db.Model(&comments[i]).Related(&comments[i].User)
		r.db.Model(&comments[i]).Related(&comments[i].Replies)
		comments[i].ReplyCount = len(comments[i].Replies)
	}

	commentable.CommentCount = r.GetCommentCount(id)
	return &commentable, nil
}

func (r *CommentableRepo) GetCommentCount(id uint) int {
	var count int
	r.db.Model(&models.Comment{}).Where("commentable_id = ?", id).Count(&count)
	return count
}
