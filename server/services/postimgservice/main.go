package postimgservice

import (
	"outstagram/server/constants"
	"outstagram/server/dtos/dtomodels"
	"outstagram/server/models"
	"outstagram/server/repos/postimgrepo"
	"outstagram/server/services/cmtableservice"
	"outstagram/server/services/rctableservice"
)

type PostImageService struct {
	postImageRepo      *postimgrepo.PostImageRepo
	reactableService   *rctableservice.ReactableService
	commentableService *cmtableservice.CommentableService
}

func New(postImageRepo *postimgrepo.PostImageRepo,
	reactableService *rctableservice.ReactableService,
	commentableService *cmtableservice.CommentableService) *PostImageService {
	return &PostImageService{
		postImageRepo:      postImageRepo,
		reactableService:   reactableService,
		commentableService: commentableService,
	}
}

func (s *PostImageService) Save(postImage *models.PostImage) error {
	return s.postImageRepo.Save(postImage)
}

func (s *PostImageService) FindByID(id uint) (*models.PostImage, error) {
	return s.postImageRepo.FindByID(id)
}

func (s *PostImageService) Update(post *models.Post, values map[string]interface{}) error {
	return s.postImageRepo.Update(post, values)
}

func (s *PostImageService) GetDTOPostImage(postImage *models.PostImage, userID, audienceUserID uint) (*dtomodels.PostImage, error) {
	dtoPostImage := dtomodels.PostImage{
		ID:            postImage.ID,
		ImageID:       postImage.ImageID,
		OwnerID:       postImage.Post.UserID,
		Content:       postImage.Content,
		CreatedAt:     postImage.CreatedAt,
		OwnerFullname: postImage.Post.User.Fullname,
		OwnerUsername: postImage.Post.User.Username,
		ViewableID:    postImage.ViewableID,
		CommentableID: postImage.CommentableID,
		ReactableID:   postImage.ReactableID,
		ReactCount:    s.reactableService.GetReactCount(postImage.ReactableID),
		Reacted:       s.reactableService.CheckUserReaction(audienceUserID, postImage.ReactableID),
		Reactors:      s.reactableService.GetReactorDTOs(postImage.ReactableID, audienceUserID, 5, 0),
	}

	// Get post's comments
	commentable, err := s.commentableService.GetCommentsWithLimit(postImage.CommentableID, userID, constants.PostCommentCount, 0)
	if err != nil {
		return nil, err
	}

	// Map to dto's comments
	dtoPostImage.CommentCount = commentable.CommentCount
	for _, comment := range commentable.Comments {
		dtoComment := comment.ToDTO()
		dtoComment.Reacted = s.reactableService.CheckUserReaction(audienceUserID, comment.ReactableID)
		dtoPostImage.Comments = append(dtoPostImage.Comments, dtoComment)
	}

	return &dtoPostImage, nil
}
