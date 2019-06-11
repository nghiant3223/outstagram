package reactableType

type Type int

const (
	Post Type = iota + 1
	Comment
	Reply
	Story
)

