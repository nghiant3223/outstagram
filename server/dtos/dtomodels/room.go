package dtomodels

type Room struct {
	ID            uint       `json:"id"`
	Name          string     `json:"name,omitempty"`
	Type          bool       `json:"type"`
	Members       []*User    `json:"members"`
	Partner       *User      `json:"partner,omitempty"`
	Messages      []*Message `json:"messages"`
	LatestMessage *Message   `json:"lastMessage"`
	ImageID       uint       `json:"imageID,omitempty"`
}
