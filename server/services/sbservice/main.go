package sbservice

import (
	"outstagram/server/models"
	"outstagram/server/repositories/sbrepo"
)

type StoryBoardService struct {
	storyBoardRepo *sbrepo.StoryBoardRepo
}

func New(storyBoardRepository *sbrepo.StoryBoardRepo) *StoryBoardService {
	return &StoryBoardService{storyBoardRepo: storyBoardRepository}
}

func (sbs *StoryBoardService) Save(storyBoard *models.StoryBoard) error {
	return sbs.storyBoardRepo.Save(storyBoard)
}
