package rctservice

import (
	"github.com/jinzhu/gorm"
	"outstagram/server/models"
	"outstagram/server/repos/rctrepo"
)

type ReactService struct {
	reactRepo *rctrepo.ReactRepo
}

func New(reactRepo *rctrepo.ReactRepo) *ReactService {
	return &ReactService{reactRepo: reactRepo}
}

func (s *ReactService) Save(react *models.React) error {
	return s.reactRepo.Save(react)
}

func (s *ReactService) Remove(userID, reactableID uint) error {
	reacts, err := s.reactRepo.Find(map[string]interface{}{"user_id": userID, "reactable_id": reactableID})
	if err != nil {
		return err
	}

	if len(reacts) < 1 {
		return gorm.ErrRecordNotFound
	}

	if err := s.reactRepo.Delete(reacts[0]); err != nil {
		return err
	}

	return nil
}
