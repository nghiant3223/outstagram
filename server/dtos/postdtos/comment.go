package postdtos

import "time"

type Comment struct {
	ID         uint      `json:"id"`
	Content    string    `json:"content"`
	CreatedAt  time.Time `json:"createdAt"`
	Reactors   []string  `json:"reactors"`
	ReactCount int       `json:"reactCount"`
	ReplyCount int       `json:"replyCount"`
	Fullname   string    `json:"ownerFullname"`
	UserID     uint      `json:"ownerID"`
}
