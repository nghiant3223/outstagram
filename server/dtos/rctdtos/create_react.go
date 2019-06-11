package rctdtos

import "outstagram/server/enums/rctabletype"

type CreateReactRequest struct {
	Type reactableType.Type `form:"type"`
}

