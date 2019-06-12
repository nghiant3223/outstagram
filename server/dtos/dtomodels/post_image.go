package dtomodels

type PostImage struct {
	ID     uint   `json:"id"`
	Tiny   string `json:"tiny"`
	Small  string `json:"small"`
	Medium string `json:"medium"`
	Big    string `json:"big"`
	Huge   string `json:"huge"`
	Origin string `json:"origin"`
}
