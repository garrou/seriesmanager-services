package services

import (
	"seriesmanager-services/dto"
	"seriesmanager-services/entities"
	"seriesmanager-services/helpers"
	"seriesmanager-services/repositories"
)

func GetUser(id string) interface{} {
	return repositories.FindUserById(id)
}

func UpdateBanner(id, banner string) interface{} {
	res := repositories.FindUserById(id)

	if user, ok := res.(entities.User); ok {
		user.Banner = banner
		return repositories.SaveUser(user)
	}
	return false
}

func UpdateProfile(toUpdate dto.UserUpdateProfileDto) interface{} {
	res := repositories.FindUserById(toUpdate.Id)

	if user, ok := res.(entities.User); ok {
		user.Username = toUpdate.Username
		user.Email = toUpdate.Email
		return repositories.SaveUser(user)
	}
	return false
}

func UpdatePassword(toUpdate dto.UserUpdatePasswordDto) interface{} {
	res := repositories.FindUserById(toUpdate.Id)

	if user, ok := res.(entities.User); ok {
		same := helpers.ComparePassword(user.Password, toUpdate.CurrentPassword)

		if same && toUpdate.Password == toUpdate.Confirm {
			user.Password = helpers.HashPassword(toUpdate.Password)
			return repositories.SaveUser(user)
		}
	}
	return false
}
