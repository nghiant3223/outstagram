package cmtrepo

import (
	"github.com/jinzhu/gorm"
	"outstagram/server/models"
)

type CommentRepo struct {
	db *gorm.DB
}

func New(dbConnection *gorm.DB) *CommentRepo {
	return &CommentRepo{db: dbConnection}
}

func (r *CommentRepo) GetRepliesByID(id uint) ([]models.Reply, error) {
	var comment models.Comment
	var replies []models.Reply

	if err := r.db.First(&comment, id).Error; err != nil {
		return nil, err
	}

	r.db.Model(&comment).Related(&replies)
	for i := 0; i < len(replies); i++ {
		r.db.Model(&replies[i]).Related(&replies[i].User)
	}

	return replies, nil
}

func (r *CommentRepo) GetRepliesByIDWithLimit(id uint, limit int, offset int) ([]models.Reply, error) {
	var comment models.Comment
	var replies []models.Reply

	if err := r.db.First(&comment, id).Error; err != nil {
		return nil, err
	}

	if err := r.db.Where("comment_id = ?", comment.ID).
		Offset(offset).
		Limit(limit).
		Find(&replies).
		Error; err != nil {
		return nil, err
	}

	for i := 0; i < len(replies); i++ {
		r.db.Model(&replies[i]).Related(&replies[i].User)
	}

	return replies, nil
}