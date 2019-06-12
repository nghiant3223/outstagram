package rctableservice

import (
	postVisibility "outstagram/server/enums/postvisibility"
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

func (s *ReactableService) GetReactCount(reactableID uint) int {
	return s.reactableRepo.GetReactCount(reactableID)
}

func (s *ReactableService) GetVisibilityByID(reactableID uint) (postVisibility.Visibility, uint, error) {
	return s.reactableRepo.GetVisibility(reactableID)
}