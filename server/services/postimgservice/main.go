package postimgservice

import (
	"outstagram/server/models"
	"outstagram/server/repos/postimgrepo"
)

type PostImageService struct {
	postImageRepo *postimgrepo.PostImageRepo
}

func New(postImageRepo *postimgrepo.PostImageRepo) *PostImageService {
	return &PostImageService{postImageRepo: postImageRepo}
}

func (s *PostImageService) Save(postImage *models.PostImage) error {
	return s.postImageRepo.Save(postImage)
}