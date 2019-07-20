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

	redisSupplier, err := db.NewRedisSupplier()
	if err != nil {
		log.Fatal(err.Error())
	}

	users, _ := userService.GetAllUsers()
	for _, user := range users {
		posts := userService.GetPostFeed(user.ID)
		for _, post := range posts {
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

	return &Controller{
		userService:       userService,
		storyBoardService: storyBoardService,
		postService:       postService,
		imageService:      imageService,
		postImageService:  postImageService,
		viewableService:   viewableService,
	}
}
