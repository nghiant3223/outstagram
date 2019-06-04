package postservice

import (
	"outstagram/server/models"
	"outstagram/server/repos/postrepo"
)

type PostService struct {
	postRepo *postrepo.PostRepo
}

func New(postRepo *postrepo.PostRepo) *PostService {
	return &PostService{postRepo: postRepo}
}

func (ps *PostService) Save(post *models.Post) error {
	return ps.postRepo.Save(post)
}
