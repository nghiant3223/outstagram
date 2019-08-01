package mecontroller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"outstagram/server/db"
	"outstagram/server/utils"
)

func (mc *Controller) UpdateUser(c *gin.Context) {
	userID, ok := utils.RetrieveUserID(c)
	if !ok {
		log.Fatal("This route needs VerifyToken middleware")
	}

	form, err := c.MultipartForm()
	if err != nil {
		utils.ResponseWithError(c, http.StatusBadRequest, "Invalid form data", err.Error())
		return
	}

	user, err := mc.userService.FindByID(userID)
	if err != nil {
		utils.ResponseWithError(c, http.StatusBadRequest, "Invalid form data", err.Error())
		return
	}

	// Update avatar
	avatar := form.File["avatar"]
	if len(avatar) > 0 {
		image, err := mc.imageService.Save(avatar[0], userID, false)
		if err != nil {
			utils.ResponseWithError(c, http.StatusInternalServerError, "Error while saving user's avatar", err.Error())
			return
		}

		user.AvatarImageID = image.ID
		if err := mc.userService.Save(user); err != nil {
			utils.ResponseWithError(c, http.StatusBadRequest, "Invalid form data", err.Error())
			return
		}

		redisClient, _ := db.NewRedisSupplier()
		redisClient.Set(fmt.Sprintf("avatarImageID:%v", userID), image.ID, 0)
	}

	// Update cover
	cover := form.File["cover"]
	if len(cover) > 0 {

	}

	utils.ResponseWithSuccess(c, http.StatusOK, "Fetch user successfully", user.ToMeDTO())
}
