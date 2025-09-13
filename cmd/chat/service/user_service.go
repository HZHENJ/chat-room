package service

import (
	"chatting-room/cmd/chat/dal/db"
	"chatting-room/pkg/errno"
	"context"
)

type UserService struct {
	ctx context.Context
}

func NewUserService(ctx context.Context) *UserService {
	return &UserService{ctx: ctx}
}

// GetUserById get a user by user id
func (s *UserService) GetUserById(userId int64) (*db.User, error) {
	user, err := db.GetUserById(s.ctx, userId)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errno.UserIsNotExistErr
	}
	return user, nil
}

// CreateUser create a new user
func (s *UserService) CreateUser(userName, avatar string) (*db.User, error) {
	// check if user exists
	exist, err := db.GetUserByUserName(s.ctx, userName)
	if err != nil {
		return nil, err
	}
	if exist != nil {
		return nil, errno.UserAlreadyExistErr
	}

	user := &db.User{
		UserName: userName,
		Avatar:   avatar,
	}
	_, err = db.CreateUser(s.ctx, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// CheckUserIsExist check if a user exists by user id
func (s *UserService) CheckUserIsExist(userId int64) (bool, error) {
	return db.CheckUserById(s.ctx, userId)
}
