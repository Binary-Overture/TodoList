package dao

import (
	"Todolist/repository/db/model"
	"context"
	"gorm.io/gorm"
)

type UserDao struct {
	*gorm.DB
}

func NewUserDao(c context.Context) *UserDao {
	if c == nil {
		c = context.Background()
	}
	return &UserDao{NewDBClient(c)}
}

// FindUserByUserId
// 通过id和uid来获得task
func (dao *UserDao) FindUserByUserId(id uint) (user *model.UserModel, err error) {
	err = dao.Model(&model.UserModel{}).Where("id = ? ", id).First(&user).Error
	return
}
