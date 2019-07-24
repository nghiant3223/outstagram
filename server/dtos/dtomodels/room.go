package dtomodels

type Room struct {
	ID            uint          `json:"id"`
	Name          string        `json:"name,omitempty"`
	Type          bool          `json:"type"`
	Members       []*SimpleUser `json:"members"`
	Partner       *SimpleUser   `json:"partner,omitempty"`
	Messages      []*Message    `json:"messages"`
	LatestMessage *Message      `json:"lastMessage"`
	ImageID       uint          `json:"imageID,omitempty"`
}
