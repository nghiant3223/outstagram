package userrepo

import (
	"outstagram/server/models"

	"github.com/jinzhu/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func New(dbConnection *gorm.DB) *UserRepo {
	return &UserRepo{db: dbConnection}
}

func (r *UserRepo) FindAll() ([]models.User, error) {
	var users []models.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}

	for _, user := range users {
		r.db.Model(&user).Related(&user.NotifBoard, "NotifBoard")
		r.db.Model(&user).Related(&user.StoryBoard, "StoryBoard")
	}

	return users, nil
}

func (r *UserRepo) FindByUsername(username string) (*models.User, error) {
	var user models.User
	if err := r.db.Find(&user, "username = ?", username).Error; err != nil {
		return nil, err
	}

	r.db.Model(&user).Related(&user.NotifBoard, "NotifBoard")
	r.db.Model(&user).Related(&user.StoryBoard, "StoryBoard")
	return &user, nil
}

func (r *UserRepo) FindByID(id uint) (*models.User, error) {
	var user models.User
	if err := r.db.Find(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}

	r.db.Model(&user).Related(&user.NotifBoard, "NotifBoard")
	r.db.Model(&user).Related(&user.StoryBoard, "StoryBoard")
	return &user, nil
}

func (r *UserRepo) Save(user *models.User) error {
	if err := r.db.Save(user).Error; err != nil {
		return err
	}

	return nil
}

func (r *UserRepo) Create(user *models.User) error {
	if err := r.db.Create(user).Error; err != nil {
		return err
	}

	r.db.Create(&models.NotifBoard{UserID: user.ID})
	r.db.Create(&models.StoryBoard{UserID: user.ID})
	return nil
}

func (r *UserRepo) SaveAll(users []*models.User) error {
	for _, user := range users {
		err := r.db.Save(&user).Error
		if err != nil {
			return err
		}

		r.db.Create(&models.NotifBoard{UserID: user.ID})
		r.db.Create(&models.StoryBoard{UserID: user.ID})
	}

	return nil
}

func (r *UserRepo) DeleteByID(id uint) error {
	var user models.User
	if err := r.db.Where("id = ?", id).Find(&user).Error; err != nil {
		return err
	}

	r.db.Delete(&user)
	return nil
}

func (r *UserRepo) DeleteAll(ids []uint) error {
	for _, id := range ids {
		var user models.User
		if err := r.db.Where("id = ?", id).Find(&user).Error; err != nil {
			return err
		}
		r.db.Delete(&user)
	}

	return nil
}

func (r *UserRepo) ExistsByID(id uint) bool {
	var user models.User
	return !r.db.First(&user, id).RecordNotFound()
}

func (r *UserRepo) ExistsByUsername(username string) bool {
	var user models.User
	return !r.db.Where("username = ?", username).First(&user).RecordNotFound()
}

func (r *UserRepo) ExistsByEmail(email string) bool {
	var user models.User
	return !r.db.Where("email = ?", email).First(&user).RecordNotFound()
}

func (r *UserRepo) Follow(following, followID uint) error {
	return r.db.Exec("INSERT INTO follows(user_follow_id, user_followed_id) VALUES (?, ?)", following, followID).Error
}

func (r *UserRepo) Unfollow(following, followID uint) error {
	return r.db.Exec("DELETE FROM follows WHERE user_follow_id = ? AND user_followed_id = ?", following, followID).Error
}

func (r *UserRepo) GetFollowers(userID uint) []models.User {
	var users []models.User
	r.db.Raw("SELECT user.* FROM user INNER JOIN follows ON user_follow_id = user.id WHERE follows.user_followed_id = ?", userID).Scan(&users)
	return users
}

func (r *UserRepo) GetFollowings(userID uint) []*models.User {
	var users []*models.User
	r.db.Raw("SELECT user.* FROM user INNER JOIN follows ON user_followed_id = user.id WHERE follows.user_follow_id = ?", userID).Scan(&users)
	return users
}

func (r *UserRepo) GetFollowingsWithAffinity(userID uint) []*models.User {
	var users []*models.User
	r.db.Raw("SELECT user.* FROM user INNER JOIN follows ON user_followed_id = user.id WHERE follows.user_follow_id = ? ORDER BY quality DESC", userID).Scan(&users)
	for i := 0; i < len(users); i++ {
		user := users[i]
		r.db.Model(&user).Related(&user.StoryBoard)
	}

	return users
}

func (r *UserRepo) CheckFollow(follow, followed uint) (bool, error) {
	rows, err := r.db.Raw("SELECT 1 FROM follows WHERE user_follow_id = ? AND user_followed_id = ?", follow, followed).Rows()
	if err != nil {
		return false, err
	}

	return rows.Next(), nil
}

func (r *UserRepo) GetPostFeed(userID uint) []models.Post {
	var posts []models.Post
	query := `
	SELECT candidate_post.id
	FROM (
		SELECT p.id
		FROM (
			SELECT *
				FROM views
				INNER JOIN viewable 
				ON views.viewable_id = viewable.id
				WHERE user_id = ?
			) AS user_views_post
			INNER JOIN post AS p 
			ON p.viewable_id = user_views_post.viewable_id
		) AS seen_post
	
		RIGHT JOIN
	
		(
		SELECT p.*, f.quality
		FROM post AS p
		INNER JOIN follows AS f 
		ON p.user_id = f.user_followed_id
		WHERE f.user_follow_id = ? AND quality IS NOT NULL
		) AS candidate_post 
	
		ON seen_post.id = candidate_post.id
		ORDER BY 
			CASE WHEN seen_post.id IS NULL
				 THEN candidate_post.popularity * 0.25 + candidate_post.quality * 0.75 END DESC, candidate_post.created_at DESC, candidate_post.created_at DESC,
			CASE WHEN seen_post.id IS NOT NULL
				 THEN candidate_post.popularity * 0.25 + candidate_post.quality * 0.75 END DESC, candidate_post.created_at DESC;`
	r.db.Raw(query, userID, userID).Scan(&posts)
	return posts
}

func (r *UserRepo) Search(text string, options ...map[string]interface{}) ([]*models.User, error) {
	var users []*models.User

	query := r.db

	if text != "" {
		regex := "%" + text + "%"
		query = query.
			Or("username LIKE ?", regex).
			Or("fullname LIKE ?", regex).
			Or("email LIKE ?", regex)
	}

	if err := query.Find(&users).Error; err != nil {
		return nil, err
	}

	if len(options) > 0 {
		options := options[0]
		if meUserID, ok := options["include_me"]; ok {
			if id, _ok := meUserID.(uint); _ok {
				foundIdx := -1
				for i, user := range users {
					if user.ID == id {
						foundIdx = i
					}
				}
				if foundIdx != -1 {
					users = append(users[:foundIdx], users[foundIdx+1:]...)
				}
			}
		}
	}

	return users, nil
}

func (r *UserRepo) GetUserRoomIDs(userID uint) ([]uint, error) {
	user, err := r.FindByID(userID)
	if err != nil {
		return nil, err
	}

	var roomIDs []uint
	r.db.Model(&user).Association("Rooms").Find(&user.Rooms)
	for _, room := range user.Rooms {
		roomIDs = append(roomIDs, room.ID)
	}
	return roomIDs, nil
}

func (r *UserRepo) GetFollowSuggestions(userID uint) []*models.User {
	var allUsers []*models.User
	var candidateUsers []*models.User

	r.db.Where("").Find(&allUsers)
	for _, user := range allUsers {
		if user.ID == userID {
			continue
		}
		followed, err := r.CheckFollow(userID, user.ID)
		if err != nil {
			continue
		}
		if !followed {
			candidateUsers = append(candidateUsers, user)
		}
		if len(candidateUsers) >= 20 {
			return candidateUsers
		}
	}

	return candidateUsers
}
