package postVisibility

type Visibility int

const (
	Public Visibility = iota
	OnlyFollowers
	Private
)
