package postrepo

import (
	"github.com/jinzhu/gorm"
	"outstagram/server/models"
)

type PostRepo struct {
	db *gorm.DB
}

func New(dbConnection *gorm.DB) *PostRepo {
	return &PostRepo{db: dbConnection}
}

func (pr *PostRepo) Save(post *models.Post) error {
	commentable := models.Commentable{}
	reactable := models.Reactable{}
	viewable := models.Viewable{}

	pr.db.Create(&commentable)
	pr.db.Create(&reactable)
	pr.db.Create(&viewable)
	post.CommentableID = commentable.ID
	post.ReactableID = reactable.ID
	post.ViewableID = viewable.ID

	return pr.db.Create(post).Error
}
