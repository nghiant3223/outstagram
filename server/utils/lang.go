package utils

import (
	"time"
)

func NewStringPointer(str string) *string {
	return &str
}

func NewTimePointer(t time.Time) *time.Time {
	return &t
}
