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

func (fc *Controller) RemoveFollow(c *gin.Context) {
	userID, ok := utils.RetrieveUserID(c)
	if !ok {
		log.Fatal("This route needs verifyToken middleware")
	}

	followingID, err := utils.StringToUint(c.Param("followingID"))
	if err != nil {
		utils.ResponseWithError(c, http.StatusBadRequest, "Invalid parameter", err.Error())
		return
	}

	if err := fc.userService.Unfollow(userID, followingID); err != nil {
		if gorm.IsRecordNotFoundError(err) {
			utils.ResponseWithError(c, http.StatusNotFound, "User not found", err.Error())
			return
		}

		if err.Error() == constants.NotExist {
			utils.ResponseWithError(c, http.StatusNotFound, "Not followed yet", err.Error())
			return
		}

		utils.ResponseWithError(c, http.StatusInternalServerError, "Error while following user", err.Error())
		return
	}

	fc.removeFollowingPostFromNewsfeed(userID, followingID)
	utils.ResponseWithSuccess(c, http.StatusNoContent, "Follow user successfully", nil)
}

func (fc *Controller) removeFollowingPostFromNewsfeed(userID, followingID uint) {
	redisSupplier, _ := db.NewRedisSupplier()
	sRedisPosts, err := redisSupplier.LRange(fmt.Sprintf("newsfeed:%v", userID), 0, 200).Result()
	if err != nil {
		log.Printf("Cannot get user with userID = %v newsfeed\n", userID)
		return
	}

	for _, sRedisPost := range sRedisPosts {
		var post models.RedisPost

		err := json.Unmarshal([]byte(sRedisPost), &post)
		if err != nil {
			log.Printf("Cannot remove post from user with userID = %v newsfeed\n", userID)
			continue
		}

		if post.OwnerID == followingID {
			redisSupplier.LRem(fmt.Sprintf("newsfeed:%v", userID), 0, sRedisPost)
		}
	}
}
