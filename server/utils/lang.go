package utils

import (
	"fmt"
	"outstagram/server/models"
	"reflect"
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

func NewUintPointer(u uint) *uint {
	return &u
}

func StringToUint(s string) (uint, error) {
	num, err := strconv.ParseUint(s, 10, 32)
	if err != nil {
		return 0, err
	}

	return uint(num), nil
}

func GetImageSize(image *models.Image, field string) string {
	r := reflect.ValueOf(image)
	f := reflect.Indirect(r).FieldByName(field)
	fmt.Println("=>", field)
	
	return string(f.String())
}