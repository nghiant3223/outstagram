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

func (r *PostRepo) Save(post *models.Post) error {
	commentable := models.Commentable{}
	reactable := models.Reactable{}
	viewable := models.Viewable{}

	r.db.Create(&commentable)
	r.db.Create(&reactable)
	r.db.Create(&viewable)
	post.CommentableID = commentable.ID
	post.ReactableID = reactable.ID
	post.ViewableID = viewable.ID

	return r.db.Create(post).Error
}

func (r *PostRepo) FindByID(id uint) (*models.Post, error) {
	var post models.Post
	err := r.db.First(&post, id).Error
	if err != nil {
		return nil, err
	}

	r.db.Model(&post).Related(&post.PostImages, "PostImages")
	return &post, nil
}