package cmtablerepo

import (
	"outstagram/server/db"
	"testing"
)

var dbConn, _ = db.New()
var cr = New(dbConn)

func TestCommentableRepo_GetCommentsByID(t *testing.T) {
	cmts, err := cr.GetComments(1)
	if err != nil {
		t.Error(err.Error())
	}
	if len(cmts) < 1 {
		t.Error("No cmt found")
	}
}
