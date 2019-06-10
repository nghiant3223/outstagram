package postservice

import (
	"github.com/jinzhu/gorm"
	"outstagram/server/enums/postenums"
	"outstagram/server/models"
	"outstagram/server/repos/cmtablerepo"
	"outstagram/server/repos/postrepo"
	"outstagram/server/repos/userrepo"
)

type PostService struct {
	postRepo        *postrepo.PostRepo
	userRepo        *userrepo.UserRepo
	commentableRepo *cmtablerepo.CommentableRepo
}

func New(postRepo *postrepo.PostRepo, userRepo *userrepo.UserRepo) *PostService {
	return &PostService{postRepo: postRepo, userRepo: userRepo}
}

func (s *PostService) Save(post *models.Post) error {
	return s.postRepo.Save(post)
}

// GetUserPosts return array of all posts that user has
func (s *PostService) GetUserPosts(userID uint) ([]models.Post, error) {
	posts, err := s.postRepo.GetPostsByUserID(userID)
	if err != nil {
		return nil, err
	}

	return posts, nil
}

// GetUsersPostsWithLimit returns array of posts with their basic info
func (s *PostService) GetUsersPostsWithLimit(userID uint, limit uint, offset uint) ([]models.Post, error) {
	posts, err := s.postRepo.GetPostsByUserIDWithLimit(userID, limit, offset)
	if err != nil {
		return nil, err
	}

	return posts, nil
}

// GetPostByID lets user get the post that has the postID specified in parameter
// User may be restricted to view the post due to its visibility. In such case, ErrRecordNotFound is returned.
// `userID` is the id of user who wants to view the post
func (s *PostService) GetPostByID(userID, postID uint) (*models.Post, error) {
	post, err := s.postRepo.FindByID(postID)
	
	if err != nil {
		return nil, err
	}

	if userID == post.UserID {
		return post, nil
	}

	if post.Visibility == postenums.Public {
		return post, nil
	}

	if post.Visibility == postenums.Private {
		return nil, gorm.ErrRecordNotFound
	}

	// If post.Visibility is OnlyFollowers
	ok, err := s.userRepo.CheckFollow(userID, post.UserID)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, gorm.ErrRecordNotFound
	}

	return post, nil
}
