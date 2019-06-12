package postPrivacy

type Privacy int

const (
	Public Privacy = iota + 1
	OnlyFollowers
	Private
)
