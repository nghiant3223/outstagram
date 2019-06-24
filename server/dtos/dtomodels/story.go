package dtomodels

import "time"

type Story struct {
	ID          uint      `json:"id"`
	CreatedAt   time.Time `json:"createdAt"`
	Duration    uint      `json:"duration"`
	Tiny        string    `json:"tiny"`
	Small       string    `json:"small"`
	Medium      string    `json:"medium"`
	Big         string    `json:"big"`
	Huge        string    `json:"huge"`
	Origin      string    `json:"origin"`
	ReactableID uint      `json:"reactableID"`
	ViewableID  uint      `json:"viewableID"`
	Seen        *bool     `json:"seen,omitempty"`
}
