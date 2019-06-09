package postservice

import (
	"outstagram/server/models"
	"outstagram/server/repos/cmtablerepo"
	"outstagram/server/repos/postrepo"
)

type PostService struct {
	postRepo        *postrepo.PostRepo
	commentableRepo *cmtablerepo.CommentableRepo
}

func New(postRepo *postrepo.PostRepo) *PostService {
	return &PostService{postRepo: postRepo}
}

func (s *PostService) Save(post *models.Post) error {
	return s.postRepo.Save(post)
}

// GetUserPosts return array of all posts that user has
func (s *PostService) GetUserPosts(userID uint) ([]models.Post, error) {
	posts, err  := s.postRepo.GetPostsByUserID(userID)
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
