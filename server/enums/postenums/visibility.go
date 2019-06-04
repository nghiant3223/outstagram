package postenums

type Visibility int

const (
	PUBLIC Visibility = iota
	ONLY_FOLLOWERS
	PRIVATE
)
