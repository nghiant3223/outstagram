package cmtableservice

import (
	postVisibility "outstagram/server/enums/postvisibility"
	"outstagram/server/models"
	"outstagram/server/repos/cmtablerepo"
	"outstagram/server/services/rctableservice"
)

type CommentableService struct {
	commentableRepo  *cmtablerepo.CommentableRepo
	reactableService *rctableservice.ReactableService
}

func New(commentableRepo *cmtablerepo.CommentableRepo, reactableService *rctableservice.ReactableService) *CommentableService {
	return &CommentableService{commentableRepo: commentableRepo, reactableService: reactableService}
}

func (s *CommentableService) GetCommentCount(commentableID uint) int {
	return s.commentableRepo.GetCommentCount(commentableID)
}

func (s *CommentableService) GetCommentsWithLimit(id, userID, limit, offset uint) (*models.Commentable, error) {
	commentable, err := s.commentableRepo.GetCommentsWithLimit(id, limit, offset)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(commentable.Comments); i++ {
		comment := &commentable.Comments[i]
		comment.Reactors = s.reactableService.GetReactorsFullname(comment.CommentableID, userID)
		comment.ReactCount = s.reactableService.GetReactCount(comment.CommentableID)
	}

	return commentable, nil
}

func (s *CommentableService) GetComments(id, userID uint) (*models.Commentable, error) {
	commentable, err := s.commentableRepo.GetComments(id)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(commentable.Comments); i++ {
		comment := &commentable.Comments[i]
		comment.Reactors = s.reactableService.GetReactorsFullname(comment.CommentableID, userID)
		comment.ReactCount = s.reactableService.GetReactCount(comment.CommentableID)
	}

	return commentable, nil
}

func (s *CommentableService) GetVisibilityByID(commentableID uint) (postVisibility.Visibility, uint, error) {
	return s.commentableRepo.GetVisibility(commentableID)
}

func (s *CommentableService) CheckHasComment(cmtableID, cmtID uint) bool {
	return s.commentableRepo.HasComment(cmtableID, cmtID)
}
