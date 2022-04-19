package services

import (
	"services-series-manager/dto"
	"services-series-manager/helpers"
	"services-series-manager/models"
	"services-series-manager/repositories"
)

type UserService interface {
	Update(user dto.UserUpdateDto) interface{}
}

type userService struct {
	userRepository repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) UserService {
	return &userService{userRepository: userRepository}
}

func (u *userService) Update(toUpdate dto.UserUpdateDto) interface{} {
	res := u.userRepository.FindById(toUpdate.Id)

	if user, ok := res.(models.User); ok {
		if toUpdate.Password != "" {
			user.Password = helpers.HashPassword(toUpdate.Password)
		}
		user.Email = toUpdate.Email
		return u.userRepository.Save(user)
	}
	return false
}
