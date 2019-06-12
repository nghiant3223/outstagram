package postcontroller

import (
	"outstagram/server/dtos/dtomodels"
	"outstagram/server/models"
	"outstagram/server/services/cmtableservice"
	"outstagram/server/services/imgservice"
	"outstagram/server/services/postimgservice"
	"outstagram/server/services/postservice"
	"outstagram/server/services/rctableservice"
	"outstagram/server/services/vwableservice"
)

type Controller struct {
	postService        *postservice.PostService
	imageService       *imgservice.ImageService
	postImageService   *postimgservice.PostImageService
	commentableService *cmtableservice.CommentableService
	reactableService   *rctableservice.ReactableService
	viewableService    *vwableservice.ViewableService

}

func New(postService *postservice.PostService, imageService *imgservice.ImageService, postImageService *postimgservice.PostImageService, commentableService *cmtableservice.CommentableService, reactableService *rctableservice.ReactableService, viewableService *vwableservice.ViewableService) *Controller {
	return &Controller{postService: postService, imageService: imageService, postImageService: postImageService, commentableService: commentableService, reactableService: reactableService, viewableService: viewableService}
}

// getDTOPost maps post, including post's images, post's comments into a DTO object
func (pc *Controller) getDTOPost(post *models.Post) (*dtomodels.Post, error) {
	// Set basic post's info
	dtoPost := dtomodels.Post{
		ID:            post.ID,
		CreatedAt:     post.CreatedAt,
		Content:       post.Content,
		Visibility:    post.Visibility,
		ImageCount:    len(post.Images),
		NumViewed:     post.NumViewed,
		OwnerID:       post.UserID,
		OwnerFullname: post.User.Fullname,
		ReactCount:    pc.reactableService.GetReactCount(post.ReactableID),
		Reactors:      pc.reactableService.GetReactors(post.ReactableID)}

	// Map post's images to DTO
	for _, postImage := range post.Images {
		image := postImage.Image
		dtoPostImage := dtomodels.PostImage{
			ID:     postImage.ID,
			Tiny:   image.Tiny,
			Origin: image.Origin,
			Huge:   image.Huge, Big: image.Huge,
			Medium: image.Medium,
			Small:  image.Small,
		}
		dtoPost.Images = append(dtoPost.Images, dtoPostImage)
	}

	// Get post's comments
	commentable, err := pc.commentableService.GetCommentsWithLimit(post.CommentableID, 5, 0)
	if err != nil {
		return nil, err
	}

	// Mapping post's comments to DTO
	dtoPost.CommentCount = commentable.CommentCount
	for _, comment := range commentable.Comments {
		dtoComment := pc.getDTOComment(&comment)
		dtoPost.Comments = append(dtoPost.Comments, dtoComment)
	}

	return &dtoPost, nil
}

//getDTOComment maps comment into a DTO object
func (pc *Controller) getDTOComment(comment *models.Comment) dtomodels.Comment {
	return dtomodels.Comment{
		ID:            comment.ID,
		Content:       comment.Content,
		ReplyCount:    comment.ReplyCount,
		CreatedAt:     comment.CreatedAt,
		OwnerFullname: comment.User.Fullname,
		OwnerID:       comment.UserID,
		ReactCount:    pc.reactableService.GetReactCount(comment.ReactableID),
		Reactors:      pc.reactableService.GetReactors(comment.ReactableID)}
}

// getDTOReply maps a reply into a DTO object
func (pc *Controller) getDTOReply(reply *models.Reply) dtomodels.Reply {
	return dtomodels.Reply{
		ID:            reply.ID,
		Content:       reply.Content,
		CreatedAt:     reply.CreatedAt,
		OwnerID:       reply.UserID,
		OwnerFullname: reply.User.Fullname}
}

// checkValidComment checks if:
// 1. Comment, post exist
// 2. Comment belongs to post
// 3. User has the authorization to view the post, to comment the post
//func (pc *Controller) checkValidComment(postID, userID, commentID uint) *utils.HttpError {
//	post, err := pc.postService.GetPostByID(postID, userID)
//	if err != nil {
//		if gorm.IsRecordNotFoundError(err) {
//			return utils.NewHttpError(http.StatusNotFound, "Post not found", err.Error())
//		}
//
//		return utils.NewHttpError(http.StatusInternalServerError, "Error while retrieving post", err.Error())
//	}
//
//	comment, err := pc.commentService.GetPostByID(commentID)
//	if err != nil {
//		if gorm.IsRecordNotFoundError(err) {
//			return utils.NewHttpError(http.StatusNotFound, "Comment not found", err.Error())
//		}
//
//		return utils.NewHttpError(http.StatusInternalServerError, "Error while retrieving comment", err.Error())
//	}
//
//	if comment.CommentableID != post.CommentableID {
//		return utils.NewHttpError(http.StatusConflict, "Comment doesn't belong to post", fmt.Sprintf("commentID %v doesn't belong to postID %v", userID, postID))
//	}
//
//	return nil
//}
