package mecontroller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
	"outstagram/server/dtos/dtomodels"
	"outstagram/server/dtos/medtos"
	"outstagram/server/dtos/storydtos"
	"outstagram/server/utils"
)

func (mc *Controller) GetMe(c *gin.Context) {
	userID, ok := utils.RetrieveUserID(c)
	if !ok {
		log.Fatal("This route needs verifyToken middleware")
	}

	user, err := mc.userService.FindByID(userID)
	if gorm.IsRecordNotFoundError(err) {
		utils.ResponseWithError(c, http.StatusNotFound, "User not found", nil)
		return
	}

	utils.ResponseWithSuccess(c, http.StatusOK, "Fetch user successfully", user.ToMeDTO())
}

func (mc *Controller) GetNewsFeed(c *gin.Context) {
	userID, ok := utils.RetrieveUserID(c)
	if !ok {
		log.Fatal("This route needs verifyToken middleware")
	}

	var getNewsfeedResponse medtos.GetNewsFeedResponse

	ids := mc.userService.GetPostFeed(userID)
	for _, postID := range ids {
		post, err := mc.postService.GetPostByID(postID, userID)
		if err != nil {
			if gorm.IsRecordNotFoundError(err) {
				continue
			}

			utils.ResponseWithError(c, http.StatusInternalServerError, "Error while retrieving post", err.Error())
			return
		}

		dtoPost, err := mc.postService.GetDTOPost(post, userID)
		if err != nil {
			if gorm.IsRecordNotFoundError(err) {
				utils.ResponseWithError(c, http.StatusNotFound, "Post not found", err.Error())
				return
			}

			utils.ResponseWithError(c, http.StatusInternalServerError, "Error while retrieving post", err.Error())
			return
		}

		getNewsfeedResponse.Posts = append(getNewsfeedResponse.Posts, *dtoPost)
	}

	utils.ResponseWithSuccess(c, http.StatusOK, "Get post successfully", getNewsfeedResponse)
}

func (mc *Controller) GetStoryFeed(c *gin.Context) {
	userID, ok := utils.RetrieveUserID(c)
	if !ok {
		log.Fatal("This route needs verifyToken middleware")
	}

	var storyBoardResponse storydtos.GetStoryFeedResponse
	var activeStoryBoard []*dtomodels.StoryBoard
	var inactiveStoryBoard []*dtomodels.StoryBoard

	userStoryBoardDTO, err := mc.storyBoardService.GetUserStoryBoardDTO(userID)
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			utils.ResponseWithError(c, http.StatusNotFound, "Not found", err.Error())
			return
		}

		utils.ResponseWithError(c, http.StatusInternalServerError, "Error while retrieving story board", err.Error())
		return
	}

	userStoryBoardDTO.IsMy = true
	storyBoardResponse.StoryBoards = append(storyBoardResponse.StoryBoards, userStoryBoardDTO)
	followings := mc.userService.GetFollowingsWithAffinity(userID)
	for _, following := range followings {
		storyBoardDTO, err := mc.storyBoardService.GetFollowingStoryBoardDTO(userID, following)

		if err != nil {
			if gorm.IsRecordNotFoundError(err) {
				utils.ResponseWithError(c, http.StatusNotFound, "Not found", err.Error())
				return
			}

			utils.ResponseWithError(c, http.StatusInternalServerError, "Error while retrieving story board", err.Error())
			return
		}

		if storyBoardDTO == nil {
			continue
		}

		if storyBoardDTO.HasNewStory {
			activeStoryBoard = append(activeStoryBoard, storyBoardDTO)
		} else {
			inactiveStoryBoard = append(inactiveStoryBoard, storyBoardDTO)
		}
	}

	storyBoardResponse.StoryBoards = append(storyBoardResponse.StoryBoards, activeStoryBoard...)
	storyBoardResponse.StoryBoards = append(storyBoardResponse.StoryBoards, inactiveStoryBoard...)
	utils.ResponseWithSuccess(c, http.StatusOK, "Retrieve story feed successfully", storyBoardResponse)
}
