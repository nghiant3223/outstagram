package models

import (
	"github.com/jinzhu/gorm"
	"outstagram/server/dtos/dtomodels"
)

// Story entity
type 	Story struct {
	gorm.Model
	ImageID      uint  `gorm:"not null"`
	Image        Image `gorm:"foreignkey:ImageID"`
	StoryBoardID uint
	ReactableID  uint
	ViewableID   uint
}

func (s *Story) ToDTO() dtomodels.Story {
	return dtomodels.Story{
		ID:          s.ID,
		ViewableID:  s.ViewableID,
		ReactableID: s.ReactableID,
		Small:       s.Image.Small,
		Medium:      s.Image.Medium,
		Huge:        s.Image.Huge,
		Big:         s.Image.Big,
		Origin:      s.Image.Origin,
		Tiny:        s.Image.Tiny}
}
