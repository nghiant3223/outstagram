package userdtos

type SearchUserRequest struct {
	Filter    string `form:"filter" binding:"required"`
	IncludeMe *bool   `form:"include_me"`
}
