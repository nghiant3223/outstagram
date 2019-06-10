package postdtos

import "time"

type Reply struct {
	ID            uint      `json:"id"`
	Content       *string   `json:"content"`
	OwnerID       uint      `json:"ownerID"`
	OwnerFullname string    `json:"ownerFullname"`
	CreatedAt     time.Time `json:"createdAt"`
}
