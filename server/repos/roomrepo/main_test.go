package roomrepo

import (
	"outstagram/server/db"
	"testing"
)

var dbConn, _ = db.New()
var roomRepo = New(dbConn)

func TestRoomRepo_CheckDualRoomExist(t *testing.T) {
	existed := roomRepo.CheckDualRoomExist(1, 3)
	if existed {
		t.Error("Fail")
	}
}

func TestRoomRepo_CheckDualRoomExist2(t *testing.T) {
	existed := roomRepo.CheckDualRoomExist(1, 2)
	if !existed {
		t.Error("Fail")
	}
}