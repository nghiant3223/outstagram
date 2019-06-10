package postdtos

type GetPostComments struct {
	Limit  uint `form:"limit"`
	Offset uint `form:"offset"`
}
