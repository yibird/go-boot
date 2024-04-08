package sys

import (
	"gorm.io/gorm"
)

type UserService struct {
	DB *gorm.DB
}

func NewUserService(DB *gorm.DB) *UserService {
	return &UserService{DB}
}

// GetList 获取用户列表
func (s *UserService) GetList() {

}

// GetUserByAccount 根据账号查询用户
func (s *UserService) GetUserByAccount(account string) *User {
	var user *User
	err := s.DB.First(&user, "account = ?", account).Error
	if err != nil {
		return nil
	}
	return user
}
