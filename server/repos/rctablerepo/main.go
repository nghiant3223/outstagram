package rctablerepo

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	privacyLevel "outstagram/server/enums/postprivacy"
	"outstagram/server/models"
	"outstagram/server/repos/cmtablerepo"
	"strings"
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

func (r *ReactableRepo) GetReactsWithLimit(id uint, limit uint, offset uint) ([]models.React, error) {
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

	return reactable.Reacts, nil
}

func (r *ReactableRepo) GetReactCount(reactableID uint) int {
	var count int
	r.db.Model(&models.React{}).Where("reactable_id = ?", reactableID).Count(&count)

	return count
}

func (r *ReactableRepo) GetVisibility(reactableID uint) (privacyLevel.Privacy, uint, error) {
	var reactable models.Reactable
	var post models.Post
	var postImage models.PostImage
	var comment models.Comment
	var reply models.Reply
	var story models.Story

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

	r.db.Model(&reactable).Related(&story)
	if story.ID != 0 {
		// TODO: Change story visibility here
		r.db.Model(&story).Related(&story.User)
		return privacyLevel.Public, story.User.ID, nil
	}

	return 0, 0, errors.New(fmt.Sprintf("Database error, invalid use of reactable_id = %v", reactableID))
}

func (r *ReactableRepo) GetReactorsOrderByQuality(reactableID, userID uint, limit int) []models.User {
	var users []models.User
	var currentUserIndex int

	query := `
SELECT reactors.*
FROM 
	(SELECT user.* FROM user INNER JOIN react ON user.id = react.user_id WHERE reactable_id = ?) reactors
	LEFT JOIN
	(SELECT user.id as following_id, quality FROM user INNER JOIN follows ON user.id = user_followed_id WHERE user_follow_id = ?) followings
	ON reactors.id = following_id
ORDER BY quality DESC 
LIMIT ?
`
	r.db.Raw(query, reactableID, userID, limit).Scan(&users)

	for i := range users {
		if users[i].ID == userID {
			currentUserIndex = i
			break
		}
	}

	// Put current user to the start of `users` array
	if currentUserIndex != 0 {
		sortedUsers := append([]models.User{}, users[currentUserIndex])
		sortedUsers = append(sortedUsers, users[:currentUserIndex]...)
		sortedUsers = append(sortedUsers, users[currentUserIndex+1:]...)
		return sortedUsers
	}

	return users
}

func (r *ReactableRepo) Find(projector map[string]interface{}) ([]*models.Reactable, error) {
	var fields []string
	var values []interface{}
	var reactables []*models.Reactable

	for key, v := range projector {
		fields = append(fields, fmt.Sprintf("%v = ?", key))
		values = append(values, v)
	}

	projections := strings.Join(fields, " AND ")
	if err := r.db.Where(projections, values...).Find(&reactables).Error; err != nil {
		return nil, err
	}

	return reactables, nil
}

func (r *ReactableRepo) CheckUserReaction(userID, reactableID uint) bool {
	var count int
	r.db.Model(&models.React{}).Where("user_id = ? AND reactable_id = ?", userID, reactableID).Count(&count)
	return count > 0
}
