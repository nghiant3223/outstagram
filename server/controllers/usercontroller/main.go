package usercontroller

import (
	"encoding/json"
	"fmt"
	"log"
	"outstagram/server/db"
	"outstagram/server/models"
	"outstagram/server/services/imgservice"
	"outstagram/server/services/postimgservice"
	"outstagram/server/services/postservice"
	"outstagram/server/services/storybservice"
	"outstagram/server/services/userservice"
	"outstagram/server/services/vwableservice"
)

type Controller struct {
	userService       *userservice.UserService
	storyBoardService *storybservice.StoryBoardService
	postService       *postservice.PostService
	imageService      *imgservice.ImageService
	postImageService  *postimgservice.PostImageService
	viewableService   *vwableservice.ViewableService
}

func New(userService *userservice.UserService,
	storyBoardService *storybservice.StoryBoardService,
	postService *postservice.PostService,
	imageService *imgservice.ImageService,
	postImageService *postimgservice.PostImageService,
	viewableService *vwableservice.ViewableService) *Controller {

	return &Controller{
		userService:       userService,
		storyBoardService: storyBoardService,
		postService:       postService,
		imageService:      imageService,
		postImageService:  postImageService,
		viewableService:   viewableService,
	}
}

func (uc *Controller) InitNewsfeed() {
	redisSupplier, err := db.NewRedisSupplier()
	if err != nil {
		log.Fatal(err.Error())
	}

	users, _ := uc.userService.GetAllUsers()
	for _, user := range users {
		posts := uc.userService.GetPostFeed(user.ID)
		for _, post := range posts {
			fmt.Printf("userID: %v - postID: %v\n", user.ID, post.ID)
			sRedisPost, err := json.Marshal(models.RedisPost{ID: post.ID, OwnerID: post.User.ID})
			if err != nil {
				log.Printf("Cannot push post to user newsfeed: %v\n", err.Error())
				continue
			}

			if err := redisSupplier.LPush(fmt.Sprintf("newsfeed:%v", user.ID), string(sRedisPost)).Err(); err != nil {
				log.Fatal(err.Error())
			}
		}
	}
}
