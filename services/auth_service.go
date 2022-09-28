package services

import (
	"github.com/google/uuid"
	"seriesmanager-services/dto"
	"seriesmanager-services/entities"
	"seriesmanager-services/helpers"
	"seriesmanager-services/repositories"
	"time"
)

func Register(user dto.UserCreateDto) dto.UserDto {
	toCreate := entities.User{
		ID:       uuid.New().String(),
		Username: user.Username,
		Email:    user.Email,
		Password: helpers.HashPassword(user.Password),
		JoinedAt: time.Now(),
	}
	created := repositories.SaveUser(toCreate)

	return dto.UserDto{
		Username: created.Username,
		Email:    created.Email,
		JoinedAt: created.JoinedAt,
		Banner:   created.Banner,
	}
}

func Login(email, password string) interface{} {
	res := repositories.FindUserByEmail(email)

	if user, ok := res.(entities.User); ok {
		same := helpers.ComparePassword(user.Password, password)

		if user.Email == email && same {
			return res
		}
		return false
	}
	return false
}

func IsDuplicateEmail(email string) bool {
	res := repositories.UserExists(email)
	return res.Error == nil
}
