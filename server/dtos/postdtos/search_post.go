package postdtos

type SearchPostRequest struct {
	Filter string `form:"filter" binding:"required"`
}
