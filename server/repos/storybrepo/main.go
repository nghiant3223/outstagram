package storybrepo

import (
	"github.com/jinzhu/gorm"
	"outstagram/server/models"
	"sort"
)

type StoryBoardRepo struct {
	db *gorm.DB
}

func New(dbConnection *gorm.DB) *StoryBoardRepo {
	return &StoryBoardRepo{db: dbConnection}
}

func (r *StoryBoardRepo) Save(board *models.StoryBoard) error {
	return r.db.Create(board).Error
}

func (r *StoryBoardRepo) SaveStory(story *models.Story) error {
	reactable := models.Reactable{}
	viewable := models.Viewable{}

	r.db.Create(&reactable)
	r.db.Create(&viewable)
	story.ReactableID = reactable.ID
	story.ViewableID = viewable.ID

	return r.db.Create(&story).Error
}

// Get stories of a specific user by his storyBoardID
func (r *StoryBoardRepo) GetStories(storyBoardID uint) ([]models.Story, error) {
	var board models.StoryBoard

	if err := r.db.First(&board, storyBoardID).Error; err != nil {
		return nil, err
	}

	r.db.Model(&board).Related(&board.Stories)
	for i := 0; i < len(board.Stories); i++ {
		storyPtr := &board.Stories[i]
		r.db.Model(&storyPtr).Related(&storyPtr.Image)
	}
	sort.Sort(byCreatedAt(board.Stories))
	return board.Stories, nil
}

func (r *StoryBoardRepo) FindStoryByID(storyID uint) (*models.Story, error) {
	var story models.Story

	if err := r.db.First(&story, storyID).Error; err != nil {
		return nil, err
	}

	r.db.Model(&story).Related(&story.Image)
	return &story, nil
}

// Check if user has viewed the story
func (r *StoryBoardRepo) CheckUserViewedStory(userID, storyID uint) (bool, error) {
	var story models.Story

	if err := r.db.First(&story, storyID).Error; err != nil {
		return false, err
	}

	rows, err := r.db.Raw("SELECT 1 FROM views WHERE user_id = ? AND viewable_id = ?", userID, story.ViewableID).Rows()
	if err != nil {
		return false, err
	}

	defer rows.Close()
	return rows.Next(), nil
}
