package usercontroller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
	"outstagram/server/dtos/dtomodels"
	"outstagram/server/dtos/postdtos"
	"outstagram/server/dtos/userdtos"
	"outstagram/server/models"
	"outstagram/server/utils"
)

func (uc *Controller) SearchUser(c *gin.Context) {
	var req userdtos.SearchUserRequest

	if err := c.ShouldBindQuery(&req); err != nil {
		utils.ResponseWithError(c, http.StatusBadRequest, "Bad request", err.Error())
		return
	}

	options := make(map[string]interface{})
	if req.IncludeMe != nil && !*req.IncludeMe {
		if userID, ok := utils.RetrieveUserID(c); ok {
			fmt.Println("here")
			options["include_me"] = userID
		}
	}

	users, err := uc.userService.Search(req.Filter, options)
	if err != nil {
		utils.ResponseWithError(c, http.StatusInternalServerError, "Error while searching for user", err.Error())
		return
	}

	var dtoUsers []dtomodels.User
	for _, user := range users {
		dtoUsers = append(dtoUsers, user.ToUserDTO())
	}

	utils.ResponseWithSuccess(c, http.StatusOK, "Fetching user successfully", dtoUsers)
}

func (uc *Controller) GetUsersInfo(c *gin.Context) {
	username := c.Param("userID")
	if username == "" {
		utils.ResponseWithError(c, http.StatusBadRequest, "Username must be provided", nil)
		return
	}

	user, err := uc.userService.FindByUsername(username)
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			utils.ResponseWithError(c, http.StatusNotFound, "Not found", err.Error())
			return
		}

		utils.ResponseWithError(c, http.StatusInternalServerError, "Error while retrieving story board", err.Error())
		return
	}

	audienceUserID, ok := utils.RetrieveUserID(c)
	if !ok {
		log.Fatal("This route needs VerifyToken middleware")
	}

	var res userdtos.GetUserResponse

	res.ID = user.ID
	res.CreatedAt = user.CreatedAt
	res.Fullname = user.Fullname
	res.Username = user.Username
	res.FollowerCount = len(uc.userService.GetFollowers(user.ID))
	res.FollowingCount = len(uc.userService.GetFollowings(user.ID))

	posts, _ := uc.postService.GetUserPosts(user.ID)
	res.PostCount = len(posts)

	isMe := audienceUserID == user.ID
	if !isMe {
		ok, err := uc.userService.CheckFollow(audienceUserID, user.ID)
		if err != nil {
			utils.ResponseWithError(c, http.StatusOK, "Error while retrieving user", err.Error())
			return
		}

		res.Followed = utils.NewBoolPointer(ok)
	}

	res.IsMe = isMe
	utils.ResponseWithSuccess(c, http.StatusOK, "Retrieve user's info successfully", res)
}

func (uc *Controller) GetUserStoryBoard(c *gin.Context) {
	userID, err := utils.StringToUint(c.Param("userID"))
	if err != nil {
		utils.ResponseWithError(c, http.StatusBadRequest, "Invalid userID", err.Error())
		return
	}

	var res userdtos.GetStoryBoardResponse

	userStoryBoardDTO, err := uc.storyBoardService.GetUserStoryBoardDTO(userID)
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			utils.ResponseWithError(c, http.StatusNotFound, "Not found", err.Error())
			return
		}

		utils.ResponseWithError(c, http.StatusInternalServerError, "Error while retrieving story board", err.Error())
		return
	}

	userStoryBoardDTO.IsMy = false
	res.StoryBoard = userStoryBoardDTO
	utils.ResponseWithSuccess(c, http.StatusOK, "Get user's storyboard successfully", res)
}

func (uc *Controller) GetUserPosts(c *gin.Context) {
	audienceUserID, ok := utils.RetrieveUserID(c)
	if !ok {
		log.Fatal("This route needs VerifyToken middleware")
	}

	userID, err := utils.StringToUint(c.Param("userID"))
	if err != nil {
		utils.ResponseWithError(c, http.StatusBadRequest, "Invalid userID", err.Error())
		return
	}

	var req postdtos.GetPostRequest
	var res postdtos.GetPostResponse
	var posts []models.Post

	if err := c.ShouldBindQuery(&req); err != nil {
		utils.ResponseWithError(c, http.StatusBadRequest, "Invalid query parameter", err.Error())
		return
	}

	// If limit and offset are not specified
	if req.Offset == 0 && req.Limit == 0 {
		posts, err = uc.postService.GetUserPosts(userID)
	} else {
		posts, err = uc.postService.GetUsersPostsWithLimit(userID, req.Limit, req.Offset)
	}

	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			utils.ResponseWithSuccess(c, http.StatusNoContent, "No posts", nil)
			return
		}

		utils.ResponseWithError(c, http.StatusInternalServerError, "Error while retrieving user's posts", err.Error())
		return
	}

	for _, post := range posts {
		dtoPost, err := uc.postService.GetDTOPost(&post, userID, audienceUserID)
		if err != nil {
			if gorm.IsRecordNotFoundError(err) {
				utils.ResponseWithError(c, http.StatusNotFound, "Post not found", err.Error())
				return
			}

			utils.ResponseWithError(c, http.StatusInternalServerError, "Error while retrieving post", err.Error())
			return
		}

		res.Posts = append(res.Posts, *dtoPost)
	}

	utils.ResponseWithSuccess(c, http.StatusOK, "Fetch user's posts successfully", res)
}
