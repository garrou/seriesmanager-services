package services

import (
	"github.com/google/uuid"
	"seriesmanager-services/dto"
	"seriesmanager-services/entities"
	"seriesmanager-services/helpers"
	"seriesmanager-services/repositories"
	"time"
)

type AuthService interface {
	Register(user dto.UserCreateDto) dto.UserDto
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

func (a *authService) Register(user dto.UserCreateDto) dto.UserDto {
	toCreate := entities.User{
		ID:       uuid.New().String(),
		Username: user.Username,
		Email:    user.Email,
		Password: helpers.HashPassword(user.Password),
		JoinedAt: time.Now(),
	}
	created := a.userRepository.Save(toCreate)

	return dto.UserDto{
		Username: created.Username,
		Email:    created.Email,
		JoinedAt: created.JoinedAt,
		Banner:   created.Banner,
	}
}

func (a *authService) Login(email, password string) interface{} {
	res := a.userRepository.FindByEmail(email)

	if user, ok := res.(entities.User); ok {
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
