package storybservice

import (
	"outstagram/server/models"
	"outstagram/server/repos/storybrepo"
)

type StoryBoardService struct {
	storyBoardRepo *storybrepo.StoryBoardRepo
}

func New(storyBoardRepo *storybrepo.StoryBoardRepo) *StoryBoardService {
	return &StoryBoardService{storyBoardRepo: storyBoardRepo}
}

func (s *StoryBoardService) Save(storyBoard *models.StoryBoard) error {
	return s.storyBoardRepo.Save(storyBoard)
}
