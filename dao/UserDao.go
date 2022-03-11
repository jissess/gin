package dao

import (
	"myGin/common"
	"myGin/model"
)

type UserDao struct {
	*common.Orm
}

func NewUserDao() *UserDao {
	return &UserDao{common.DbEngine}
}

func (ud *UserDao) Insert(user model.User) (*model.User, error) {
	if err := ud.DB.Create(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (ud *UserDao) QueryByTelephone(telephone string) *model.User {
	var user model.User
	ud.DB.Where("telephone = ?", telephone).First(&user)
	if user.Id != 0 {
		return &user
	}
	return nil
}
