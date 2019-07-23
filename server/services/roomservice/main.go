package roomservice

import (
	"errors"
	"outstagram/server/models"
	"outstagram/server/repos/roomrepo"
	"outstagram/server/services/userservice"
)

type RoomService struct {
	roomRepo    *roomrepo.RoomRepo
	userService *userservice.UserService
}

func New(roomRepo *roomrepo.RoomRepo,
	userService *userservice.UserService) *RoomService {
	return &RoomService{
		roomRepo:    roomRepo,
		userService: userService,
	}
}

func (s *RoomService) CreateDualRoom(userAID, userBID uint) error {
	var room models.Room

	exist := s.roomRepo.CheckDualRoomExist(userAID, userBID)
	if exist {
		return errors.New("room existed")
	}

	for _, id := range []uint{userAID, userBID} {
		user, err := s.userService.FindByID(id)
		if err != nil {
			return err
		}
		room.Members = append(room.Members, user)
	}

	room.Type = false
	return s.roomRepo.Save(&room)
}

func (s *RoomService) GetRoomByID(id uint) (*models.Room, error) {
	return s.roomRepo.FindByID(id)
}

func (s *RoomService) GetRoomByPartnerUsername(userID uint, partnerUsername string) (*models.Room, error) {
	partner, err := s.userService.FindByUsername(partnerUsername)
	if err != nil {
		return nil, err
	}
	room := s.roomRepo.FindByPartnerID(userID, partner.ID)
	return room, nil
}

func (s *RoomService) CreateMessage(roomID uint, message *models.Message) error {
	message.RoomID = roomID
	return s.roomRepo.SaveMessage(message)
}

func (s *RoomService) GetRecentRooms(userID uint) ([]*models.Room, error) {
	return s.roomRepo.RecentRooms(userID)
}

func (s *RoomService) GetRoomMessages(roomID uint) (*models.Room, error) {
	return s.roomRepo.GetRoomMessages(roomID)
}

func (s *RoomService) GetRoomMessagesWithLimit(roomID uint, limit, offset uint) (*models.Room, error) {
	return s.roomRepo.GetRoomMessagesWithLimit(roomID, limit, offset)
}
