package flcontroller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
	"outstagram/server/constants"
	"outstagram/server/db"
	"outstagram/server/models"
	"outstagram/server/utils"
)

func (fc *Controller) CreateFollow(c *gin.Context) {
	userID, ok := utils.RetrieveUserID(c)
	if !ok {
		log.Fatal("This route needs verifyToken middleware")
	}

	followingID, err := utils.StringToUint(c.Param("followingID"))
	if err != nil {
		utils.ResponseWithError(c, http.StatusBadRequest, "Invalid parameter", err.Error())
		return
	}

	if userID == followingID {
		utils.ResponseWithError(c, http.StatusConflict, "Cannot follow yourself", nil)
		return
	}

	if err := fc.userService.Follow(userID, followingID); err != nil {
		if gorm.IsRecordNotFoundError(err) {
			utils.ResponseWithError(c, http.StatusNotFound, "User not found", err.Error())
			return
		}

		if err.Error() == constants.AlreadyExist {
			utils.ResponseWithError(c, http.StatusNotFound, "Already followed", err.Error())
			return
		}

		utils.ResponseWithError(c, http.StatusInternalServerError, "Error while following user", err.Error())
		return
	}

	posts, _ := fc.postService.GetUsersPostsWithLimit(followingID, 10, 0)
	fc.appendFollowingPostToNewsfeed(userID, posts)
	utils.ResponseWithSuccess(c, http.StatusCreated, "Follow user successfully", nil)
}

func (fc *Controller) appendFollowingPostToNewsfeed(userID uint, posts []models.Post) {
	redisSupplier, _ := db.NewRedisSupplier()
	for _, post := range posts {
		sRedisPost, err := json.Marshal(models.RedisPost{ID: post.ID, OwnerID: post.User.ID})
		if err != nil {
			log.Printf("Cannot append post to user with userID = %v newsfeed\n", userID)
			continue
		}

		redisSupplier.LPush(fmt.Sprintf("newsfeed:%v", userID), string(sRedisPost))
	}
}
