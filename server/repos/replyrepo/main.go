package replyrepo

import (
	"github.com/jinzhu/gorm"
	"outstagram/server/models"
)

type ReplyRepo struct {
	db *gorm.DB
}

func New(dbConnection *gorm.DB) *ReplyRepo {
	return &ReplyRepo{db: dbConnection}
}
func (r *ReplyRepo) Save(reply *models.Reply) error {
	reactable := models.Reactable{}
	r.db.Create(&reactable)
	reply.ReactableID = reactable.ID
	err := r.db.Create(&reply).Error
	if err != nil {
		return err
	}

	// WARNING: This is for retrieve the information of the reply's owner
	r.db.Model(&reply).Related(&reply.User)
	return nil
}

func (r *ReplyRepo) FindByID(id uint) (*models.Reply, error) {
	var reply models.Reply

	if err := r.db.First(&reply, id).Error; err != nil {
		return nil, err
	}

	return &reply, nil
}
