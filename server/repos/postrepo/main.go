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

	r.db.Model(&post).Related(&post.User)
	r.db.Model(&post).Related(&post.Images)
	for j := 0; j < len(post.Images); j++ {
		r.db.Model(&post.Images[j]).Related(&post.Images[j].Image)
	}

	return &post, nil
}

func (r *PostRepo) GetPostsByUserIDWithLimit(userID, limit, offset uint) ([]models.Post, error) {
	var posts []models.Post
	if err := r.db.Where("user_id = ?", userID).
		Offset(offset).
		Limit(limit).
		Find(&posts).
		Error; err != nil {
		return nil, err
	}

	for i := 0; i < len(posts); i++ {
		r.db.Model(&posts[i]).Related(&posts[i].Images)
		for j := 0; j < len(posts[i].Images); j++ {
			r.db.Model(&posts[i].Images[j]).Related(&posts[i].Images[j].Image)
			r.db.Model(&posts[i]).Related(&posts[i].User)
		}
	}

	return posts, nil
}

func (r *PostRepo) GetPostsByUserID(userID uint) ([]models.Post, error) {
	var posts []models.Post
	if err := r.db.Where("user_id = ?", userID).
		Find(&posts).
		Error; err != nil {
		return nil, err
	}

	for i := 0; i < len(posts); i++ {
		r.db.Model(&posts[i]).Related(&posts[i].Images)
	}

	return posts, nil
}
