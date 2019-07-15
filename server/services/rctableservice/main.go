package rctableservice

import (
	"outstagram/server/constants"
	"outstagram/server/dtos/dtomodels"
	postVisibility "outstagram/server/enums/postprivacy"
	"outstagram/server/models"
	"outstagram/server/repos/rctablerepo"
	"outstagram/server/services/userservice"
	"outstagram/server/utils"
)

type ReactableService struct {
	reactableRepo *rctablerepo.ReactableRepo
	userService   *userservice.UserService
}

func New(reactableRepo *rctablerepo.ReactableRepo,
	userService *userservice.UserService) *ReactableService {
	return &ReactableService{
		reactableRepo: reactableRepo,
		userService:   userService,
	}
}

func (s *ReactableService) GetReactorsOrderByQuality(reactableID, userID uint, limit, offset int) []models.User {
	return s.reactableRepo.GetReactorsOrderByQuality(reactableID, userID, limit, offset)
}

func (s *ReactableService) GetReactCount(reactableID uint) int {
	return s.reactableRepo.GetReactCount(reactableID)
}

func (s *ReactableService) GetVisibilityByID(reactableID uint) (postVisibility.Privacy, uint, error) {
	return s.reactableRepo.GetVisibility(reactableID)
}

func (s *ReactableService) GetReactors(reactableID, userID uint, limit, offset int) []models.User {
	return s.reactableRepo.GetReactorsOrderByQuality(reactableID, userID, limit, offset)
}

func (s *ReactableService) GetReactorDTOs(reactableID, userID uint, limit, offset int) []dtomodels.SimpleUser {
	var reactors []dtomodels.SimpleUser
	for _, user := range s.reactableRepo.GetReactorsOrderByQuality(reactableID, userID, limit, offset) {
		dtoReactor := user.ToSimpleUser()
		followed, _ := s.userService.CheckFollow(userID, user.ID)
		dtoReactor.Followed = utils.NewBoolPointer(followed)
		reactors = append(reactors, dtoReactor)
	}

	return reactors
}
func (s *ReactableService) GetReactorsFullname(reactableID, userID uint) []string {
	var users []string

	for _, user := range s.GetReactorsOrderByQuality(reactableID, userID, constants.ReactorCount, 0) {
		users = append(users, user.Fullname)
	}

	return users
}

func (s *ReactableService) CheckUserReaction(userID, reactableID uint) bool {
	return s.reactableRepo.CheckUserReaction(userID, reactableID)
}
