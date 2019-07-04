package storybservice

import (
	"outstagram/server/constants"
	"outstagram/server/dtos/dtomodels"
	"outstagram/server/models"
	"outstagram/server/repos/storybrepo"
	"outstagram/server/services/rctableservice"
	"outstagram/server/services/userservice"
	"outstagram/server/utils"
)

type StoryBoardService struct {
	storyBoardRepo   *storybrepo.StoryBoardRepo
	userService      *userservice.UserService
	reactableService *rctableservice.ReactableService
}

func New(storyBoardRepo *storybrepo.StoryBoardRepo, userService *userservice.UserService, reactableService *rctableservice.ReactableService) *StoryBoardService {
	return &StoryBoardService{storyBoardRepo: storyBoardRepo, userService: userService, reactableService: reactableService}
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

// GetFollowingStoryBoardDTO return story of userA towards userB, whom userA follows
func (s *StoryBoardService) GetFollowingStoryBoardDTO(userAID uint, userB *models.User) (*dtomodels.StoryBoard, error) {
	stories, err := s.GetStories(userB.StoryBoard.ID)
	if err != nil {
		return nil, err
	}

	if len(stories) < 1 {
		return nil, nil
	}

	dtoStoryBoard := dtomodels.StoryBoard{
		StoryBoardID: userB.StoryBoard.ID,
		UserID:       userB.ID,
		Fullname:     userB.Fullname,
		StoryCount:   len(stories)}

	hasNewStoryFlag := false
	for _, story := range stories {
		dtoStory := story.ToDTO()

		// Check story has been seen by userA
		seen, err := s.CheckUserViewedStory(userAID, story.ID)
		if err != nil {
			return nil, err
		}
		dtoStory.Seen = utils.NewBoolPointer(seen)
		if !hasNewStoryFlag && !seen {
			dtoStoryBoard.HasNewStory = true
			hasNewStoryFlag = true
		}

		// Check story has been reacted by userA
		reacted := s.reactableService.CheckUserReaction(userAID, story.ReactableID)
		dtoStory.Reacted = utils.NewBoolPointer(reacted)

		dtoStoryBoard.Stories = append(dtoStoryBoard.Stories, dtoStory)
	}

	return &dtoStoryBoard, nil
}

// GetUserStoryBoardDTO return story of a user
func (s *StoryBoardService) GetUserStoryBoardDTO(userID uint) (*dtomodels.StoryBoard, error) {
	user, err := s.userService.FindByID(userID)
	if err != nil {
		return nil, err
	}

	stories, err := s.GetStories(user.StoryBoard.ID)
	if err != nil {
		return nil, err
	}

	dtoStoryBoard := dtomodels.StoryBoard{
		StoryBoardID: user.StoryBoard.ID,
		UserID:       user.ID,
		Fullname:     user.Fullname,
		StoryCount:   len(stories),
		HasNewStory:  len(stories) > 0}

	for _, story := range stories {
		dtoStory := story.ToDTO()
		storyReactors := s.reactableService.GetReactorsOrderByQuality(story.ReactableID, userID, constants.StoryReactorCount)
		for _, reactor := range storyReactors {
			dtoStory.Reactors = append(dtoStory.Reactors, reactor.ToBasicUserDTO())
		}
		dtoStoryBoard.Stories = append(dtoStoryBoard.Stories, dtoStory)
	}

	return &dtoStoryBoard, nil
}
