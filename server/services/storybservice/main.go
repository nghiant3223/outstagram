package storybservice

import (
	"outstagram/server/dtos/dtomodels"
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

// GetUserStoryBoardDTO return story of userA towards userB
func (s *StoryBoardService) GetUserStoryBoardDTO(userAID uint, userB *models.User) (*dtomodels.StoryBoard, error) {
	stories, err := s.GetStories(userB.StoryBoard.ID)
	if err != nil {
		return nil, err
	}

	dtoStoryBoard := dtomodels.StoryBoard{
		UserID:     userB.ID,
		Fullname:   userB.Fullname,
		AvatarURL:  userB.AvatarURL,
		StoryCount: len(stories)}

	hasNewStoryFlag := false
	for _, story := range stories {
		dtoStory := story.ToDTO()
		seen, err := s.CheckUserViewedStory(userAID, story.ID)
		if err != nil {
			return nil, err
		}

		dtoStory.Seen = seen
		if !hasNewStoryFlag && !seen {
			dtoStoryBoard.HasNewStory = true
			hasNewStoryFlag = true
		}
		dtoStoryBoard.Stories = append(dtoStoryBoard.Stories, dtoStory)
	}

	return &dtoStoryBoard, nil
}
