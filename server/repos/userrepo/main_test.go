package userrepo

import (
	"fmt"
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
		Username: "52332",
		Password: "3422334",
		Email:    "me22232ee@gmail.com",
		Phone:    utils.NewStringPointer("la230j")}
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
		Username: "1asdd",
		Password: "3422334",
		Email:    "mde23e@gmail.com",
		Phone:    utils.NewStringPointer("033dd9")}
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
	user := models.User{
		Username: "1asd",
		Password: "3422334",
		Email:    "me23e@gmail.com",
		Phone:    utils.NewStringPointer("123")}
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

func TestUserRepository_GetFollowers(t *testing.T) {
	users := userRepo.GetFollowers(42)
	if len(users) != 2 {
		t.Error("Length users < 2")
	}
	if !((users[0].ID == 43 || users[0].ID == 44) && (users[1].ID == 44 || users[1].ID == 43)) {
		t.Error("Wrong followers")
	}
}

func TestUserRepository_GetFollowings(t *testing.T) {
	users := userRepo.GetFollowings(43)
	if len(users) != 2 {
		t.Error("Length users < 2")
	}
	if !((users[0].ID == 44 || users[0].ID == 42) && (users[1].ID == 44 || users[1].ID == 42)) {
		t.Error("Wrong followers")
	}
}

func TestUserRepository_GetFollowers2(t *testing.T) {
	users := userRepo.GetFollowers(43)
	if len(users) != 0 {
		t.Error("Length users < 2")
	}
}

func TestUserRepo_GetFollowingsWithAffinity(t *testing.T) {
	users := userRepo.GetFollowingsWithAffinity(1)
	for _, user := range users {
		if user.StoryBoard.ID == 0 {
			t.Error("No storyboard id")
		}
		fmt.Println(user.StoryBoard.ID)
	}
	fmt.Println(users)
}
