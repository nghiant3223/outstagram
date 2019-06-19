package mecontroller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
	"outstagram/server/dtos/dtomodels"
	"outstagram/server/dtos/medtos"
	"outstagram/server/dtos/storydtos"
	"outstagram/server/models"
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
		fmt.Println((dtoPost).ImageCount)
	}

	utils.ResponseWithSuccess(c, http.StatusOK, "Get post successfully", getNewsfeedResponse)
}

func (mc *Controller) GetStoryFeed(c *gin.Context) {
	userID, ok := utils.RetrieveUserID(c)
	if !ok {
		log.Fatal("This route needs verifyToken middleware")
	}

	var storyBoardResponse storydtos.GetStoryFeedResponse

	followings := mc.userService.GetFollowingsWithAffinity(userID)
	users, boundaryIndex, err := mc.groupStoryBoardByActive(userID, followings)

	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			utils.ResponseWithError(c, http.StatusNotFound, "Not found", err.Error())
			return
		}

		utils.ResponseWithError(c, http.StatusInternalServerError, "Error while sorting story board", err.Error())
		return
	}

	for i, user := range users {
		stories, err := mc.storyBoardService.GetStories(user.StoryBoard.ID)
		if err != nil {
			if gorm.IsRecordNotFoundError(err) {
				utils.ResponseWithError(c, http.StatusNotFound, "Not found", err.Error())
				return
			}

			utils.ResponseWithError(c, http.StatusInternalServerError, "Error while sorting story board", err.Error())
			return
		}

		if len(stories) < 1 {
			continue
		}

		dtoStoryBoard := dtomodels.StoryBoard{
			UserID:      user.ID,
			Fullname:    user.Fullname,
			AvatarURL:   user.AvatarURL,
			HasNewStory: i < boundaryIndex}

		dtoStoryBoard.StoryCount = len(stories)
		for _, story := range stories {
			dtoStoryBoard.Stories = append(dtoStoryBoard.Stories, story.ToDTO())
		}

		storyBoardResponse.StoryBoards = append(storyBoardResponse.StoryBoards, dtoStoryBoard)
	}

	utils.ResponseWithSuccess(c, http.StatusOK, "Retrieve story feed successfully", storyBoardResponse)
}

// groupStoryBoardByActive puts user that has new story at the front, user that has old story at the back
// Each user group is also sorted by the affinity between the user and their followings
func (mc *Controller) groupStoryBoardByActive(userID uint, followings []*models.User) ([]*models.User, int, error) {
	var viewed []*models.User
	var notViewed []*models.User

	for _, user := range followings {
		isNew, err := mc.storyBoardService.IsActiveStoryBoard(userID, user.ID)
		if err != nil {
			return nil, -1, err
		}

		if isNew {
			notViewed = append(notViewed, user)
		} else {
			viewed = append(viewed, user)
		}
	}

	return append(notViewed, viewed...), len(notViewed), nil
}
