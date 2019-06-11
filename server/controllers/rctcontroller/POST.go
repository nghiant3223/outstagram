package rctcontroller

import (
	"github.com/gin-gonic/gin"
)

func (rc *Controller) CreateReaction(c *gin.Context) {
	//userID, ok := utils.RetrieveUserID(c)
	//if !ok {
	//	log.Fatal("This route needs verifyToken middleware")
	//}
	//
	//var reqBody rctdtos.CreateReactRequest
	//var reactableID uint
	//var err error
	//
	//if err := c.ShouldBindQuery(&reqBody); err != nil {
	//	utils.ResponseWithError(c, http.StatusBadRequest, "Invalid query parameter", err.Error())
	//	return
	//}
	//
	//reactEntityID, err := utils.StringToUint(c.Param("rctableID"))
	//if err != nil {
	//	utils.ResponseWithError(c, http.StatusBadRequest, "Invalid parameter", err.Error())
	//	return
	//}
	//
	//switch reqBody.Type {
	//case reactableType.Post:
	//	post, err := rc.postService.GetPostByID(reactEntityID, userID)
	//	if err != nil {
	//		if gorm.IsRecordNotFoundError(err) {
	//			utils.ResponseWithError(c, http.StatusNotFound, "Post not found", err.Error())
	//			return
	//		}
	//
	//		utils.ResponseWithError(c, http.StatusInternalServerError, "Error while retrieving post", err.Error())
	//		return
	//	}
	//	reactableID = post.ReactableID
	//
	//case reactableType.Comment:
	//	//post, err := rc.commentService.GetPostByID(reactEntityID, userID)
	//	if err != nil {
	//		if gorm.IsRecordNotFoundError(err) {
	//			utils.ResponseWithError(c, http.StatusNotFound, "Post not found", err.Error())
	//			return
	//		}
	//
	//		utils.ResponseWithError(c, http.StatusInternalServerError, "Error while retrieving post", err.Error())
	//		return
	//	}
	//	//reactableID = post.ReactableID
	//
	//default:
	//	utils.ResponseWithError(c, http.StatusBadRequest, "Invalid reactable type", nil)
	//	return
	//}
	//
	//if err := rc.reactService.Save(userID, reactableID); err != nil {
	//	utils.ResponseWithError(c, http.StatusInternalServerError, "Error while saving react", err.Error())
	//	return
	//}
	//
	//utils.ResponseWithError(c, http.StatusCreated, "Save post successfully", nil)
}
