package rctablerepo

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	postVisibility "outstagram/server/enums/postvisibility"
	"outstagram/server/models"
	"outstagram/server/repos/cmtablerepo"
)

type ReactableRepo struct {
	db              *gorm.DB
	commentableRepo *cmtablerepo.CommentableRepo
}

func New(dbConnection *gorm.DB, commentableRepo *cmtablerepo.CommentableRepo) *ReactableRepo {
	return &ReactableRepo{db: dbConnection, commentableRepo: commentableRepo}
}

func (r *ReactableRepo) GetReacts(id uint) ([]models.React, error) {
	var reactable models.Reactable
	var reacts []models.React

	if err := r.db.First(&reactable, id).Error; err != nil {
		return nil, err
	}

	r.db.Model(&reactable).Related(&reacts)
	for i := 0; i < len(reacts); i++ {
		r.db.Model(&reacts[i]).Related(&reacts[i].User)
	}

	return reacts, nil
}

func (r *ReactableRepo) GetReactsWithLimit(id uint, limit uint, offset uint) (*models.Reactable, error) {
	var reactable models.Reactable

	if err := r.db.First(&reactable, id).Error; err != nil {
		return nil, err
	}

	if err := r.db.Where("reactable_id = ?", reactable.ID).
		Offset(offset).
		Limit(limit).
		Find(&reactable.Reacts).
		Error; err != nil {
		return nil, err
	}

	reacts := reactable.Reacts
	for i := 0; i < len(reacts); i++ {
		r.db.Model(&reacts[i]).Related(&reacts[i].User)
	}

	return &reactable, nil
}

func (r *ReactableRepo) GetReactCount(reactableID uint) int {
	var count int
	r.db.Model(&models.React{}).Where("reactable_id = ?", reactableID).Count(&count)

	return count
}

func (r *ReactableRepo) GetVisibility(reactableID uint) (postVisibility.Visibility, uint, error) {
	var reactable models.Reactable
	var post models.Post
	var postImage models.PostImage
	var comment models.Comment
	var reply models.Reply

	if err := r.db.First(&reactable, reactableID).Error; err != nil {
		return 0, 0, nil
	}

	r.db.Model(&reactable).Related(&post)
	if post.ID != 0 {
		return r.commentableRepo.GetVisibility(post.CommentableID)
	}

	r.db.Model(&reactable).Related(&postImage)
	if postImage.ID != 0 {
		return r.commentableRepo.GetVisibility(postImage.CommentableID)
	}

	r.db.Model(&reactable).Related(&comment)
	if comment.ID != 0 {
		return r.commentableRepo.GetVisibility(comment.CommentableID)
	}

	r.db.Model(&reactable).Related(&reply)
	if reply.ID != 0 {
		r.db.Model(&reply).Related(&reply.Comment)
		return r.commentableRepo.GetVisibility(reply.Comment.CommentableID)
	}

	return 0, 0, errors.New(fmt.Sprintf("Database error, invalid use of reactable_id = %v", reactableID))
}

func (r *ReactableRepo) GetReactors(reactableID, userID uint, limit int) []models.User {
	var users []models.User
	query := `
SELECT reactors.*
FROM 
	(SELECT user.* FROM user INNER JOIN react on user.id = react.user_id WHERE reactable_id = ?) reactors
	LEFT JOIN
	(SELECT user.id as following_id, level_of_interest, avatar_url as avatarURL FROM user INNER JOIN follows ON user.id = user_followed_id WHERE user_follow_id = ?) followings
	ON reactors.id = following_id
ORDER BY level_of_interest DESC 
LIMIT ?
`
	r.db.Raw(query, reactableID, userID, limit).Scan(&users)

	return users
}
