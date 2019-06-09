package postdtos

import "outstagram/server/enums/postenums"

type Post struct {
	ID           uint                 `json:"id"`
	Images       []PostImage          `json:"images"`
	ImageCount   int                  `json:"imageCount"`
	Comments     []Comment            `json:"comments"`
	CommentCount int                  `json:"commentCount"`
	Visibility   postenums.Visibility `json:"visibility"`
	Content      *string              `json:"content"`
	NumRead      int                  `json:"numRead"`
	Reactors     []string             `json:"reactors"`
	ReactCount   int                  `json:"reactCount"`
}
