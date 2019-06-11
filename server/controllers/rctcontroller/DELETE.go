package rctcontroller

import (
	"github.com/gin-gonic/gin"
)

func (rc *Controller) RemoveReaction(c *gin.Context) {
	//userID, ok := utils.RetrieveUserID(c)
	//if !ok {
	//	log.Fatal("This route needs verifyToken middleware")
	//}
	//
	//postID, err := utils.StringToUint(c.Param("postID"))
	//if err != nil {
	//	utils.ResponseWithError(c, http.StatusBadRequest, "Invalid parameter", err.Error())
	//	return
	//}
	//
	//post, err := rc.postService.GetPostByID(postID, userID)
	//if err != nil {
	//	if gorm.IsRecordNotFoundError(err) {
	//		utils.ResponseWithError(c, http.StatusNotFound, "Post not found", err.Error())
	//		return
	//	}
	//
	//	utils.ResponseWithError(c, http.StatusInternalServerError, "Error while retrieving post", err.Error())
	//	return
	//}
	//
	//if err := rc.reactService.Remove(userID, post.ReactableID); err != nil {
	//	utils.ResponseWithError(c, http.StatusInternalServerError, "Error while removing react", err.Error())
	//	return
	//}
	//
	//utils.ResponseWithError(c, http.StatusCreated, "Remove post successfully", nil)
}
