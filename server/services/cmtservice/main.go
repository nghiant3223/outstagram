package cmtservice

import (
	"outstagram/server/models"
	"outstagram/server/repos/cmtrepo"
)

type CommentService struct {
	commentRepo *cmtrepo.CommentRepo
}

func New(commentRepo *cmtrepo.CommentRepo) *CommentService {
	return &CommentService{commentRepo: commentRepo}
}

func (s *CommentService) GetReplyCount(commentableID uint) int {
	return s.commentRepo.GetReplyCount(commentableID)
}

func (s *CommentService) GetRepliesWithLimit(id uint, limit uint, offset uint) (*models.Comment, error) {
	return s.commentRepo.GetRepliesWithLimit(id, limit, offset)
}

func (s *CommentService) Save(comment *models.Comment) error {
	return s.commentRepo.Save(comment)
}

func (s *CommentService) SaveReply(reply *models.Reply) error {
	return s.commentRepo.SaveReply(reply)
}

func (s *CommentService) GetReplies(id uint) (*models.Comment, error) {
	return s.commentRepo.GetReplies(id)
}