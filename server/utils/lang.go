package utils

import (
	"strconv"
	"time"
)

func NewStringPointer(str string) *string {
	return &str
}

func NewTimePointer(t time.Time) *time.Time {
	return &t
}

func NewBoolPointer(b bool) *bool {
	return &b
}

func StringToUint(s string) (uint, error) {
	num, err := strconv.ParseUint(s, 10, 32)
	if err != nil {
		return 0, err
	}

	return uint(num), nil
}
