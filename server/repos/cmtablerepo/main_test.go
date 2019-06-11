package cmtablerepo

import (
	"outstagram/server/db"
)

var dbConn, _ = db.New()
var r = New(dbConn)