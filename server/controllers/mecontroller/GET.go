package mecontroller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
	"outstagram/server/constants"
	"outstagram/server/db"
	"outstagram/server/dtos/dtomodels"
	"outstagram/server/dtos/medtos"
	"outstagram/server/dtos/storydtos"
	"outstagram/server/models"
	"outstagram/server/utils"
)

func (mc *Controller) GetMe(c *gin.Context) {
	userID, ok := utils.RetrieveUserID(c)
	if !ok {
		log.Fatal("This route needs VerifyToken middleware")
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
		log.Fatal("This route needs VerifyToken middleware")
	}

	var res medtos.GetNewsFeedResponse
	var req medtos.GetNewsFeedRequest

	if err := c.ShouldBindQuery(&req); err != nil {
		utils.ResponseWithError(c, http.StatusBadRequest, "Invalid parameter", err.Error())
		return
	}

	var posts []string
	var err error
	var start int64
	var stop int64

	redisSupplier, _ := db.NewRedisSupplier()
	key := fmt.Sprintf("newsfeed:%v", userID)

	// Get start & stop for fetching posts from redis. Next NextSinceID equals to the id of post at index 'stop'
	if !req.Pagination {
		start = 0
		stop = constants.NewsfeedPaginationMax
	} else if req.SinceID == 0 {
		start = 0
		stop = constants.NewsfeedPaginationCount

		sRedisPosts, err := redisSupplier.LRange(key, 0, constants.NewsfeedPaginationCount).Result()
		if err != nil {
			utils.ResponseWithError(c, http.StatusInternalServerError, "Error while retrieving post", err.Error())
			return
		}

		// If there is NextSinceID
		if len(sRedisPosts) > int(constants.NewsfeedPaginationCount) {
			var rPost models.RedisPost
			if err := json.Unmarshal([]byte(sRedisPosts[len(sRedisPosts)-1 ]), &rPost); err != nil {
				log.Println(err.Error())
			}
			res.NextSinceID = rPost.ID
		} else {
			res.NextSinceID = 0
		}
	} else {
		sRedisPosts, err := redisSupplier.LRange(key, 0, constants.NewsfeedPaginationMax).Result()
		if err != nil {
			utils.ResponseWithError(c, http.StatusInternalServerError, "Error while retrieving post", err.Error())
			return
		}

		// Iterate to get right SinceID
		for i, sRedisPost := range sRedisPosts {
			var rPost models.RedisPost
			if err := json.Unmarshal([]byte(sRedisPost), &rPost); err != nil {
				log.Println(err.Error())
				continue
			}

			if rPost.ID == req.SinceID {
				start = int64(i)
				stop = start + constants.NewsfeedPaginationCount

				if availablePostCount := int64(len(sRedisPosts)); stop >= availablePostCount {
					// If there is no NextSinceID
					// Example: 0 1 2 3 4 5, start = 3, NewsfeedPaginationCount = 3
					// stop = 5, there is no NextSinceID
					stop = availablePostCount
					res.NextSinceID = 0
				} else {
					// If there is no NextSinceID
					// Example: 0 1 2 3 4 5 6, start = 3, NewsfeedPaginationCount = 3
					// stop = 6, NextSinceID equals id of post whose id = 6
					var nextRPost models.RedisPost
					if err := json.Unmarshal([]byte(sRedisPosts[stop]), &nextRPost); err != nil {
						log.Println(err.Error())
					}
					res.NextSinceID = nextRPost.ID
				}
				break
			}
		}
	}

	posts, err = redisSupplier.LRange(key, start, stop-1).Result()
	if err != nil {
		utils.ResponseWithError(c, http.StatusInternalServerError, "Error while retrieving post", err.Error())
		return
	}

	for _, sRedisPost := range posts {
		var rPost models.RedisPost

		if err := json.Unmarshal([]byte(sRedisPost), &rPost); err != nil {
			log.Println(err.Error())
			continue
		}

		post, err := mc.postService.GetPostByID(rPost.ID, userID)
		if err != nil {
			if gorm.IsRecordNotFoundError(err) {
				continue
			}

			utils.ResponseWithError(c, http.StatusInternalServerError, "Error while retrieving post", err.Error())
			return
		}

		dtoPost, err := mc.postService.GetDTOPost(post, userID, userID)
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

	utils.ResponseWithSuccess(c, http.StatusOK, "Get post successfully", res)
}

func (mc *Controller) GetStoryFeed(c *gin.Context) {
	userID, ok := utils.RetrieveUserID(c)
	if !ok {
		log.Fatal("This route needs VerifyToken middleware")
	}

	var storyBoardResponse storydtos.GetStoryFeedResponse
	var activeStoryBoard []*dtomodels.StoryBoard
	var inactiveStoryBoard []*dtomodels.StoryBoard

	// Get user's own storyboard
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

	// Get storyboard of people whom user follows
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
