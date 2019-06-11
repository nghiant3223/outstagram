package postrepo

import (
	"outstagram/server/db"
	"outstagram/server/enums/postvisibility"
	"outstagram/server/models"
	"outstagram/server/utils"
	"testing"
)

var dbConn, _ = db.New()
var prRepo = New(dbConn)

func TestPostRepository_Create(t *testing.T) {
	post := models.Post{
		Content:    utils.NewStringPointer("lmao"),
		Visibility: postVisibility.OnlyFollowers,
	}

	prRepo.Save(&post)

	if post.ID == 0 && post.Visibility != 1 {
		t.Error("Post not created")
	}
}

func TestPostRepository_Save(t *testing.T) {
	var post models.Post
	prRepo.db.First(&post, 4)
	post.Content = utils.NewStringPointer("1")
	prRepo.Save(&post)

	var post2 models.Post
	prRepo.db.First(&post2, 4)
	if *post2.Content != "1" {
		t.Error("Fail to save")
	}
}

func TestPostRepo_FindByID(t *testing.T) {
	post, err := prRepo.FindByID(7)
	if err != nil {
		t.Error(err.Error())
		return
	}
	if post.ID == 0 {
		t.Error()
	}
}