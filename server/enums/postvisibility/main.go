package postVisibility

type Visibility int

const (
	Public Visibility = iota + 1
	OnlyFollowers
	Private
)
