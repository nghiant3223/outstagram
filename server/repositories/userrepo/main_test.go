package userrepo

import (
	"github.com/jinzhu/gorm"
	"outstagram/server/db"
	"outstagram/server/models"
	"outstagram/server/utils"
	"testing"
)

var dbConn, _ = db.New()
var userRepo = New(dbConn)

func TestUserRepository_Save(t *testing.T) {
	notifBoard := models.NotifBoard{}
	storyBoard := models.StoryBoard{}

	dbConn.Save(&notifBoard)
	dbConn.Save(&storyBoard)
	user := models.User{
		Username:     "52332",
		Password:     "3422334",
		Email:        "me22232ee@gmail.com",
		NotifBoardID: notifBoard.ID,
		StoryBoardID: storyBoard.ID,
		Phone:        utils.StringPointer("la230j")}
	userRepo.Save(&user)

	if dbConn.Where("username = ?", "52332").RecordNotFound() {
		t.Error("Fail at TestUserRepository_Save")
	}
}

func TestUserRepository_ExistsById(t *testing.T) {
	notifBoard := models.NotifBoard{}
	storyBoard := models.StoryBoard{}

	dbConn.Save(&notifBoard)
	dbConn.Save(&storyBoard)
	user := models.User{
		Username:     "1asdd",
		Password:     "3422334",
		Email:        "mde23e@gmail.com",
		NotifBoardID: notifBoard.ID,
		StoryBoardID: storyBoard.ID,
		Phone:        utils.StringPointer("033dd9")}
	userRepo.Save(&user)

	if !userRepo.ExistsByID(user.ID) {
		t.Error("Fail at TestUserRepository_ExistsById")
	}
}

func TestUserRepository_FindById(t *testing.T) {
	user, err := userRepo.FindByID(14)
	if err != nil {
		t.Error("Fail at TestUserRepository_FindById")
		return
	}
	if user.Username != "1asd" {
		t.Error("Fail at TestUserRepository_FindById")
	}
}

func TestUserRepository_FindById2(t *testing.T) {
	_, err := userRepo.FindByID(3939393)
	if err == nil {
		t.Error("Fail at TestUserRepository_FindById2")
	}
}

func TestUserRepository_DeleteByID(t *testing.T) {
	err := userRepo.DeleteByID(13)
	if err != nil {
		t.Error("Fail at TestUserRepository_DeleteByID")
	}
	user, err := userRepo.FindByID(3)
	if err != gorm.ErrRecordNotFound {
		t.Error("Unexpected error at TestUserRepository_DeleteByID")
		return
	}
	if user != nil {
		t.Error("Fail at TestUserRepository_DeleteByID")
	}
}

func TestUserRepository_FindByUsername(t *testing.T) {
	user, err := userRepo.FindByUsername("123")
	if err != nil {
		t.Error("Fail at TestUserRepository_FindByUsername: ", err.Error())
		return
	}
	if user.ID != 3 && user.Password != "456" {
		t.Error("Fail at TestUserRepository_FindByUsername")
	}
}

func TestUserRepository_Save2(t *testing.T) {
	notifBoard := models.NotifBoard{}
	storyBoard := models.StoryBoard{}

	user := models.User{
		Username:     "1asd",
		Password:     "3422334",
		Email:        "me23e@gmail.com",
		NotifBoardID: notifBoard.ID,
		StoryBoardID: storyBoard.ID,
		Phone:        utils.StringPointer("123")}
	err := userRepo.Save(&user)
	if err == nil {
		t.Error("Fail at TestUserRepository_Save2")
	}
}

func TestUserRepository_DeleteByID2(t *testing.T) {
	err := userRepo.DeleteByID(20202)
	if err == nil {
		t.Error("Fail at TestUserRepository_DeleteByID2")
	}
}

func TestUserRepository_ExistsById2(t *testing.T) {
	if !userRepo.ExistsByID(14) {
		t.Error("Fail at TestUserRepository_ExistsById2")
	}
}
