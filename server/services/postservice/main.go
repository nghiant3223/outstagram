package postservice

import (
	"github.com/jinzhu/gorm"
	"outstagram/server/dtos/dtomodels"
	"outstagram/server/enums/postprivacy"
	"outstagram/server/models"
	"outstagram/server/repos/postrepo"
	"outstagram/server/services/cmtableservice"
	"outstagram/server/services/rctableservice"
	"outstagram/server/services/userservice"
)

type PostService struct {
	postRepo           *postrepo.PostRepo
	userService        *userservice.UserService
	reactableService   *rctableservice.ReactableService
	commentableService *cmtableservice.CommentableService
}

func New(postRepo *postrepo.PostRepo, userService *userservice.UserService, reactableService *rctableservice.ReactableService, commentableService *cmtableservice.CommentableService) *PostService {
	return &PostService{postRepo: postRepo, userService: userService, reactableService:reactableService, commentableService:commentableService}
}

func (s *PostService) Save(post *models.Post) error {
	return s.postRepo.Save(post)
}

// GetUserPosts return array of all posts that user has
func (s *PostService) GetUserPosts(userID uint) ([]models.Post, error) {
	posts, err := s.postRepo.GetPostsByUserID(userID)
	if err != nil {
		return nil, err
	}

	return posts, nil
}

// GetUsersPostsWithLimit returns array of posts with their basic info
func (s *PostService) GetUsersPostsWithLimit(userID, limit, offset uint) ([]models.Post, error) {
	posts, err := s.postRepo.GetPostsByUserIDWithLimit(userID, limit, offset)
	if err != nil {
		return nil, err
	}

	return posts, nil
}

// GetPostByID lets user get the post that has the postID specified in parameter
// User may be restricted to view the post due to its visibility. In such case, ErrRecordNotFound is returned.
// `userID` is the id of user who wants to view the post
func (s *PostService) GetPostByID(postID, userID uint) (*models.Post, error) {
	post, err := s.postRepo.FindByID(postID)

	if err != nil {
		return nil, err
	}

	if userID == post.UserID {
		return post, nil
	}

	if post.Privacy == postPrivacy.Public {
		return post, nil
	}

	if post.Privacy == postPrivacy.Private {
		return nil, gorm.ErrRecordNotFound
	}

	if userID == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	// If post.Privacy is OnlyFollowers
	ok, err := s.userService.CheckFollow(userID, post.UserID)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, gorm.ErrRecordNotFound
	}

	return post, nil
}

// getDTOPost maps post, including post's images, post's comments into a DTO object
func (s *PostService) GetDTOPost(post *models.Post, userID uint) (*dtomodels.Post, error) {
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
		ReactCount:    s.reactableService.GetReactCount(post.ReactableID),
		Reactors:      s.reactableService.GetReactorsFullname(post.ReactableID, userID)}

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
	commentable, err := s.commentableService.GetCommentsWithLimit(post.CommentableID, userID, 5, 0)
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
