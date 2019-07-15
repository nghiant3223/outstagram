package postcontroller

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"outstagram/server/dtos/postdtos"
	"outstagram/server/utils"
)

func (pc *Controller) UpdatePost(c *gin.Context) {
	var reqBody postdtos.UpdatePostRequest

	if err := c.ShouldBind(&reqBody); err != nil {
		utils.ResponseWithError(c, http.StatusBadRequest, "Invalid query parameter", err.Error())
		return
	}

	postID, err := utils.StringToUint(c.Param("postID"))
	if err != nil {
		utils.ResponseWithError(c, http.StatusBadRequest, "Invalid parameter", err.Error())
		return
	}

	audienceUserID, ok := utils.RetrieveUserID(c)
	if !ok {
		log.Fatal("This route needs verifyToken middleware")
	}

	post, err := pc.postService.GetPostByID(postID, audienceUserID)
	if err != nil {
		utils.ResponseWithError(c, http.StatusInternalServerError, "Error while updating post", err.Error())
		return
	}

	if err := pc.postService.Update(post, map[string]interface{}{"content": reqBody.Content}); err != nil {
		utils.ResponseWithError(c, http.StatusInternalServerError, "Error while updating post", err.Error())
		return
	}

	if post.UserID != audienceUserID {
		utils.ResponseWithError(c, http.StatusForbidden, "Cannot update post that's not your", nil)
		return
	}

	utils.ResponseWithSuccess(c, http.StatusOK, "Update post successfully", reqBody)
}

func (pc *Controller) UpdatePostImage(c *gin.Context) {

}
