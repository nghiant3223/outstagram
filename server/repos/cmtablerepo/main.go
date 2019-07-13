package cmtablerepo

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	postVisibility "outstagram/server/enums/postprivacy"
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

func (r *CommentableRepo) GetCommentsWithLimit(id, limit, offset uint) (*models.Commentable, error) {
	var commentable models.Commentable

	if err := r.db.First(&commentable, id).Error; err != nil {
		return nil, err
	}

	if err := r.db.Raw(`
	SELECT * 
		FROM (SELECT * FROM comment WHERE commentable_id = ? ORDER BY created_at DESC LIMIT ? OFFSET ?) AS reversed
	ORDER BY created_at ASC
	`, id, limit, offset).Scan(&commentable.Comments).Error; err != nil {
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

// GetVisibility returns visibility for an commentable
func (r *CommentableRepo) GetVisibility(commentableID uint) (postVisibility.Privacy, uint, error) {
	var commentable models.Commentable
	var post models.Post
	var postImage models.PostImage

	if err := r.db.First(&commentable, commentableID).Error; err != nil {
		return 0, 0, nil
	}

	r.db.Model(&commentable).Related(&post)
	if post.ID != 0 {
		return post.Privacy, post.UserID, nil
	}

	r.db.Model(&commentable).Related(&postImage)
	if postImage.ID != 0 {
		r.db.Model(&postImage).Related(&postImage.Post)
		return postImage.Post.Privacy, postImage.Post.UserID, nil
	}

	return 0, 0, errors.New(fmt.Sprintf("Database error, invalid use of commentable_id = %v", commentableID))
}

func (r *CommentableRepo) HasComment(cmtableID, cmtID uint) bool {
	var count int
	r.db.Raw("SELECT 1 FROM comment WHERE commentable_id = ? AND id = ?", cmtableID, cmtID).Count(&count)
	return count > 0
}
