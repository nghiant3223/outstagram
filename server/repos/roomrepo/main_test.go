package roomrepo

import (
	"fmt"
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

func TestRoomRepo_GetRoomMessages(t *testing.T) {
	room, err := roomRepo.GetRoomMessages(1)
	if err != nil {
		t.Error("Fail", err.Error())
		return
	}
	if len(room.Messages) < 1 {
		t.Error("Fail", "no messages")
	}
	if len(room.Messages) != 5 {
		t.Error("Fail")
	}
	fmt.Println(room.Messages[0].User)
}

func TestRoomRepo_GetRoomMessagesWithLimit(t *testing.T) {
	room, err := roomRepo.GetRoomMessagesWithLimit(1, 3, 0)
	if err != nil {
		t.Error("Fail", err.Error())
		return
	}
	if len(room.Messages) != 3 {
		t.Error("Fail", "no messages")
	}
	fmt.Println(*room.Messages[0])
	fmt.Println(*room.Messages[1])
	fmt.Println(*room.Messages[2])
}
