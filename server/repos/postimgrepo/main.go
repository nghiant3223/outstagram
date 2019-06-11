package postimgrepo

import (
	"github.com/jinzhu/gorm"
	"outstagram/server/models"
)

type PostImageRepo struct {
	db *gorm.DB
}

func New(dbConnection *gorm.DB) *PostImageRepo {
	return &PostImageRepo{db: dbConnection}
}

func (r *PostImageRepo) Save(image *models.PostImage) error {
	commentable := models.Commentable{}
	reactable := models.Reactable{}
	viewable := models.Viewable{}

	r.db.Create(&commentable)
	r.db.Create(&reactable)
	r.db.Create(&viewable)
	image.CommentableID = commentable.ID
	image.ReactableID = reactable.ID
	image.ViewableID = viewable.ID

	return r.db.Save(image).Error
}

func (r *PostImageRepo) FindByID(id uint) (*models.PostImage, error) {
	var image models.PostImage

	if err := r.db.First(&image, id).Error; err != nil {
		return nil, err
	}

	return &image, nil
}
