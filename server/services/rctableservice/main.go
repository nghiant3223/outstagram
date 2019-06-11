package rctableservice

import (
	"outstagram/server/models"
	"outstagram/server/repos/rctablerepo"
)

type ReactableService struct {
	reactableRepo *rctablerepo.ReactableRepo
}

func New(reactableRepo *rctablerepo.ReactableRepo) *ReactableService {
	return &ReactableService{reactableRepo: reactableRepo}
}

func (s *ReactableService) GetReactsWithLimit(id uint, limit uint, offset uint) (*models.Reactable, error) {
	return s.reactableRepo.GetReactsWithLimit(id, limit, offset)
}

func (s *ReactableService) GetReactors(id uint) []string {
	// TODO: Return related reactors
	return []string{"go", "gin", "gorm"}
}

func (s *ReactableService) GetReactCount(id uint) int {
	return s.reactableRepo.GetReactCount(id)
}
