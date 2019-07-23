package authcontroller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"log"
	"net/http"
	"outstagram/server/dtos/authdtos"
	"outstagram/server/models"
	"outstagram/server/utils"
	"time"
)

func (ac *Controller) Login(c *gin.Context) {
	var reqBody authdtos.LoginRequest

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		utils.ResponseWithError(c, http.StatusBadRequest, "Some required fields missing", nil)
		return
	}

	user, err := ac.userService.VerifyLogin(reqBody.Username, reqBody.Password)
	if err != nil {
		utils.ResponseWithError(c, http.StatusNotFound, "Login failed", err.Error())
		return
	}

	token, err := utils.SignToken(user)
	if err != nil {
		utils.ResponseWithError(c, http.StatusInternalServerError, "Signing token failed", err.Error())
		return
	}

	user.LastLogin = utils.NewTimePointer(time.Now())
	if err = ac.userService.Save(user); err != nil {
		utils.ResponseWithError(c, http.StatusInternalServerError, "Saving user failed", err.Error())
		return
	}

	utils.ResponseWithSuccess(c, http.StatusOK, "Login successfully", token)
}

func (ac *Controller) Logout(c *gin.Context) {
	userID, ok := utils.RetrieveUserID(c)
	if !ok {
		log.Fatal("This route needs VerifyToken middleware")
	}

	user, err := ac.userService.FindByID(userID)
	if err != nil {
		utils.ResponseWithError(c, http.StatusInternalServerError, "Error while fetching user", err.Error())
		return
	}

	user.LastLogin = utils.NewTimePointer(time.Now())
	if err = ac.userService.Save(user); err != nil {
		utils.ResponseWithError(c, http.StatusInternalServerError, "Saving user failed", err.Error())
		return
	}

	utils.AbortRequestWithSuccess(c, http.StatusOK, "Logout user successfully", nil)
}

func (ac *Controller) Register(c *gin.Context) {
	var reqBody authdtos.RegisterRequest

	if err := c.ShouldBind(&reqBody); err != nil {
		utils.ResponseWithError(c, http.StatusBadRequest, "Some required fields missing", err.Error())
		return
	}

	if ac.userService.CheckExistsByUsername(reqBody.Username) {
		utils.ResponseWithError(c, http.StatusConflict, "Username already used", nil)
		return
	}

	if ac.userService.CheckExistsByEmail(reqBody.Username) {
		utils.ResponseWithError(c, http.StatusConflict, "Email already used", nil)
		return
	}

	newUser := models.User{}
	if err := copier.Copy(&newUser, &reqBody); err != nil {
		utils.ResponseWithError(c, http.StatusInternalServerError, "Error while copying from request body to model", err.Error())
		return
	}

	if err := ac.userService.Save(&newUser); err != nil {
		utils.ResponseWithError(c, http.StatusInternalServerError, "Creating user failed", err.Error())
		return
	}

	form, err := c.MultipartForm()
	if err != nil {
		utils.ResponseWithError(c, http.StatusBadRequest, "Invalid form data", err.Error())
		return
	}

	files := form.File["avatar"]
	if len(files) > 1 {
		utils.ResponseWithError(c, http.StatusBadRequest, "Too many avatars", nil)
		return
	}
	if len(files) == 1 {
		image, err := ac.imageService.Save(files[0], newUser.ID, false)
		if err != nil {
			utils.ResponseWithError(c, http.StatusInternalServerError, "Error while saving user's avatar", err.Error())
			return
		}
		newUser.AvatarImageID = image.ID
		if err := ac.userService.Save(&newUser); err != nil {
			utils.ResponseWithError(c, http.StatusBadRequest, "Invalid form data", err.Error())
			return
		}
	}

	utils.ResponseWithSuccess(c, http.StatusCreated, "Create user successfully", nil)
}
