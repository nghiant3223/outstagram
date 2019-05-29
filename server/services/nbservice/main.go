package nbservice

import (
	"outstagram/server/models"
	"outstagram/server/repositories/nbrepo"
)

type NotifBoardService struct {
	notifBoardRepo *nbrepo.NotifBoardRepo
}

func New(notifBoardRepository *nbrepo.NotifBoardRepo) *NotifBoardService {
	return &NotifBoardService{notifBoardRepo: notifBoardRepository}
}

func (nbs *NotifBoardService) Save(notifBoard *models.NotifBoard) error {
	return nbs.notifBoardRepo.Save(notifBoard)
}
