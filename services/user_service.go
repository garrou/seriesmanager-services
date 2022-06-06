package services

import (
	"seriesmanager-services/dto"
	"seriesmanager-services/helpers"
	"seriesmanager-services/models"
	"seriesmanager-services/repositories"
)

type UserService interface {
	Get(id string) interface{}
	UpdateBanner(id, banner string) interface{}
	UpdateProfile(toUpdate dto.UserUpdateProfileDto) interface{}
	UpdatePassword(toUpdate dto.UserUpdatePasswordDto) interface{}
}

type userService struct {
	userRepository repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) UserService {
	return &userService{userRepository: userRepository}
}

func (u *userService) Get(id string) interface{} {
	return u.userRepository.FindById(id)
}

func (u *userService) UpdateBanner(id, banner string) interface{} {
	res := u.userRepository.FindById(id)

	if user, ok := res.(models.User); ok {
		user.Banner = banner
		return u.userRepository.Save(user)
	}
	return false
}

func (u *userService) UpdateProfile(toUpdate dto.UserUpdateProfileDto) interface{} {
	res := u.userRepository.FindById(toUpdate.Id)

	if user, ok := res.(models.User); ok {
		user.Username = toUpdate.Username
		user.Email = toUpdate.Email
		return u.userRepository.Save(user)
	}
	return false
}

func (u *userService) UpdatePassword(toUpdate dto.UserUpdatePasswordDto) interface{} {
	res := u.userRepository.FindById(toUpdate.Id)

	if user, ok := res.(models.User); ok {
		same := helpers.ComparePassword(user.Password, toUpdate.CurrentPassword)

		if same && toUpdate.Password == toUpdate.Confirm {
			user.Password = helpers.HashPassword(toUpdate.Password)
			return u.userRepository.Save(user)
		}
	}
	return false
}
