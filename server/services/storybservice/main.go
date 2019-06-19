package storybservice

import (
	"outstagram/server/models"
	"outstagram/server/repos/storybrepo"
	"outstagram/server/services/userservice"
)

type StoryBoardService struct {
	storyBoardRepo *storybrepo.StoryBoardRepo
	userService    *userservice.UserService
}

func New(storyBoardRepo *storybrepo.StoryBoardRepo, userService *userservice.UserService) *StoryBoardService {
	return &StoryBoardService{storyBoardRepo: storyBoardRepo, userService: userService}
}

func (s *StoryBoardService) Save(storyBoard *models.StoryBoard) error {
	return s.storyBoardRepo.Save(storyBoard)
}
func (s *StoryBoardService) SaveStory(story *models.Story) error {
	return s.storyBoardRepo.SaveStory(story)
}

// GetStories returns stories of a specific user by his storyBoardID
func (s *StoryBoardService) GetStories(storyBoardID uint) ([]models.Story, error) {
	return s.storyBoardRepo.GetStories(storyBoardID)
}

func (s *StoryBoardService) GetStoryByID(storyID uint) (*models.Story, error) {
	return s.storyBoardRepo.FindStoryByID(storyID)
}

func (s *StoryBoardService) CheckUserViewedStory(userID, storyID uint) (bool, error) {
	return s.storyBoardRepo.CheckUserViewedStory(userID, storyID)
}

// IsActiveStoryBoard returns true if user A has not see some of user B's story
func (s *StoryBoardService) IsActiveStoryBoard(userAID, userBID uint) (bool, error) {
	userA, err := s.userService.FindByID(userAID)
	if err != nil {
		return false, err
	}

	userB, err := s.userService.FindByID(userBID)
	if err != nil {
		return false, err
	}

	stories, err := s.GetStories(userB.StoryBoard.ID)
	if err != nil {
		return false, err
	}

	for _, story := range stories {
		viewed, err := s.CheckUserViewedStory(userA.ID, story.ID)
		if err != nil {
			return false, err
		}

		if !viewed {
			return true, nil
		}
	}

	return false, nil
}
