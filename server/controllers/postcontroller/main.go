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
func (pc *Controller) getDTOPost(post *models.Post, userID uint) (*dtomodels.Post, error) {
	// Set basic post's info
	dtoPost := dtomodels.Post{
		ID:            post.ID,
		ViewableID:    post.ViewableID,
		CommentableID: post.CommentableID,
		ReactableID:   post.ReactableID,
		CreatedAt:     post.CreatedAt,
		Content:       post.Content,
		Visibility:    post.Privacy,
		ImageCount:    len(post.Images),
		OwnerID:       post.UserID,
		OwnerFullname: post.User.Fullname,
		ReactCount:    pc.reactableService.GetReactCount(post.ReactableID),
		Reactors:      pc.reactableService.GetReactorsFullname(post.ReactableID, userID)}

	// Map post's images to DTO
	for _, postImage := range post.Images {
		image := postImage.Image
		dtoPostImage := dtomodels.PostImage{
			ID:            postImage.ID,
			Tiny:          image.Tiny,
			Origin:        image.Origin,
			Huge:          image.Huge,
			Big:           image.Huge,
			Medium:        image.Medium,
			Small:         image.Small,
			ReactableID:   postImage.ReactableID,
			CommentableID: postImage.CommentableID,
			ViewableID:    postImage.ViewableID}
		dtoPost.Images = append(dtoPost.Images, dtoPostImage)
	}

	// Get post's comments
	commentable, err := pc.commentableService.GetCommentsWithLimit(post.CommentableID, userID, 5, 0)
	if err != nil {
		return nil, err
	}

	dtoPost.CommentCount = commentable.CommentCount
	for _, comment := range commentable.Comments {
		dtoComment := comment.ToDTO()
		dtoPost.Comments = append(dtoPost.Comments, dtoComment)
	}

	return &dtoPost, nil
}
