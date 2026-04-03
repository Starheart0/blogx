package user_service

import "blogx_server/models"

type UserService struct {
	userModel models.UserModel
}

var UserServiceApp = new(UserService)

func NewUserService(user models.UserModel) *UserService {
	return &UserService{
		userModel: user,
	}
}
