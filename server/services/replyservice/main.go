package replyservice

import (
	"outstagram/server/models"
	"outstagram/server/repos/replyrepo"
)

type ReplyService struct {
	replyRepo *replyrepo.ReplyRepo
}

func New(replyRepo *replyrepo.ReplyRepo) *ReplyService {
	return &ReplyService{replyRepo: replyRepo}
}

func (s *ReplyService) Save(reply *models.Reply) error {
	return s.replyRepo.Save(reply)
}
