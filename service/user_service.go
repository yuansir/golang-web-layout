package service

import (
	"golang-echo-layout/model"
	"golang-echo-layout/repository"
)

type UserService interface {
	GetUserByOpenId(openId string) *model.User
}

func NewUserService() UserService {
	return &userService{
		userRepo: repository.NewUserRepository(),
	}
}

type userService struct {
	userRepo repository.UserRepository
}

func (u *userService) GetUserByOpenId(openId string) *model.User {
	return u.userRepo.GetUserByOpenId(openId)
}
