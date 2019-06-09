package cmtableservice

import (
	"outstagram/server/models"
	"outstagram/server/repos/cmtablerepo"
)

type CommentableService struct {
	commentableRepo *cmtablerepo.CommentableRepo
}

func New(commentableRepo *cmtablerepo.CommentableRepo) *CommentableService {
	return &CommentableService{commentableRepo: commentableRepo}
}

func (s *CommentableService) GetCommentCount(commentableID uint) int {
	return s.commentableRepo.GetCommentCount(commentableID)
}

func (s *CommentableService) GetCommentsWithLimit(id uint, limit uint, offset uint) (*models.Commentable, error) {
	return s.commentableRepo.GetCommentsWithLimit(id, limit, offset)
}