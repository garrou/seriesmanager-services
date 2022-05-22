package services

import (
	"seriesmanager-services/dto"
	"seriesmanager-services/helpers"
	"seriesmanager-services/models"
	"seriesmanager-services/repositories"

	"github.com/google/uuid"
)

type AuthService interface {
	Register(user dto.UserCreateDto) models.User
	Login(email, password string) interface{}
	IsDuplicateEmail(email string) bool
}

type authService struct {
	userRepository repositories.UserRepository
}

func NewAuthService(userRepository repositories.UserRepository) AuthService {
	return &authService{
		userRepository: userRepository,
	}
}

func (a *authService) Register(user dto.UserCreateDto) models.User {
	toCreate := models.User{
		Id:       uuid.New().String(),
		Email:    user.Email,
		Password: helpers.HashPassword(user.Password),
	}
	return a.userRepository.Save(toCreate)
}

func (a *authService) Login(email, password string) interface{} {
	res := a.userRepository.FindByEmail(email)

	if user, ok := res.(models.User); ok {
		same := helpers.ComparePassword(user.Password, password)

		if user.Email == email && same {
			return res
		}
		return false
	}
	return false
}

func (a *authService) IsDuplicateEmail(email string) bool {
	res := a.userRepository.Exists(email)
	return res.Error == nil
}
