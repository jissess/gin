package service

import (
	"myGin/dao"
	"myGin/model"
)

type UserService struct {
}

func (us *UserService) Register(user model.User) (*model.User, error) {
	userDao := dao.NewUserDao()
	return userDao.Insert(user)
}

func (us *UserService) IsPhoneExists(telephone string) bool {
	userDao := dao.NewUserDao()
	user := userDao.QueryByTelephone(telephone)
	if user != nil {
		return true
	}
	return false
}
