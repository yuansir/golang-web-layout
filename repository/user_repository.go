package repository

import (
	"golang-echo-layout/database"
	"golang-echo-layout/model"
)

type UserRepository interface {
	FirstOrCreate(openId string, unionId string, sessionKey string) *model.User
	GetUserByOpenId(openId string) *model.User
}

func NewUserRepository() UserRepository {
	return &userDBRepository{}
}

type userDBRepository struct {
}

func (u *userDBRepository) FirstOrCreate(openId string, unionId string, sessionKey string) *model.User {
	user := new(model.User)
	database.Mysql.Where(model.User{OpenId: openId}).Attrs(model.User{UnionId: unionId, SessionKey: sessionKey}).FirstOrCreate(&user)
	return user
}

func (u *userDBRepository) GetUserByOpenId(openId string) *model.User {
	user := new(model.User)
	database.Mysql.Where("open_id = ?", openId).First(&user)
	return user
}
