package notifbservice

import (
	"outstagram/server/models"
	"outstagram/server/repos/notifbrepo"
)

type NotifBoardService struct {
	notifBoardRepo *notifbrepo.NotifBoardRepo
}

func New(notifBoardRepo *notifbrepo.NotifBoardRepo) *NotifBoardService {
	return &NotifBoardService{notifBoardRepo: notifBoardRepo}
}

func (s *NotifBoardService) Save(notifBoard *models.NotifBoard) error {
	return s.notifBoardRepo.Save(notifBoard)
}
