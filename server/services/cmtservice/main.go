package cmtservice

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"net/http"
	"outstagram/server/models"
	"outstagram/server/repos/cmtrepo"
	"outstagram/server/services/postservice"
	"outstagram/server/services/rctableservice"
	"outstagram/server/utils"
)

type CommentService struct {
	commentRepo      *cmtrepo.CommentRepo
	postService      *postservice.PostService
	reactableService *rctableservice.ReactableService
}

func New(commentRepo *cmtrepo.CommentRepo, reactableService *rctableservice.ReactableService) *CommentService {
	return &CommentService{commentRepo: commentRepo, reactableService: reactableService}
}

func (s *CommentService) Save(comment *models.Comment) error {
	return s.commentRepo.Save(comment)
}

func (s *CommentService) GetReplyCount(commentableID uint) int {
	return s.commentRepo.GetReplyCount(commentableID)
}

func (s *CommentService) GetRepliesWithLimit(id, userID, limit, offset uint) (*models.Comment, error) {
	comment, err := s.commentRepo.GetRepliesWithLimit(id, limit, offset)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(comment.Replies); i++ {
		reply := &comment.Replies[i]
		reply.Reactors = s.reactableService.GetReactorsFullname(reply.ReactableID, userID)
		reply.ReactCount = s.reactableService.GetReactCount(reply.ReactableID)
	}

	return comment, nil
}

func (s *CommentService) GetReplies(id, userID uint) (*models.Comment, error) {
	comment, err := s.commentRepo.GetReplies(id)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(comment.Replies); i++ {
		reply := &comment.Replies[i]
		reply.Reactors = s.reactableService.GetReactorsFullname(reply.ReactableID, userID)
		reply.ReactCount = s.reactableService.GetReactCount(reply.ReactableID)

	}

	return comment, nil
}

func (s *CommentService) CheckValidComment(postID, userID, commentID uint) *utils.HttpError {
	post, err := s.postService.GetPostByID(postID, userID)
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return utils.NewHttpError(http.StatusNotFound, "Post not found", err.Error())
		}

		return utils.NewHttpError(http.StatusInternalServerError, "Error while retrieving post", err.Error())
	}

	comment, err := s.commentRepo.FindByID(commentID)
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return utils.NewHttpError(http.StatusNotFound, "Comment not found", err.Error())
		}

		return utils.NewHttpError(http.StatusInternalServerError, "Error while retrieving comment", err.Error())
	}

	if comment.CommentableID != post.CommentableID {
		return utils.NewHttpError(http.StatusConflict, "Comment doesn't belong to post", fmt.Sprintf("commentID %v doesn't belong to postID %v", userID, postID))
	}

	return nil
}
