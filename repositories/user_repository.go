package repositories

import (
	"gorm.io/gorm"
	"seriesmanager-services/database"
	"seriesmanager-services/entities"
)

func SaveUser(user entities.User) entities.User {
	database.Db.Save(&user)
	return user
}

func FindUserByEmail(email string) interface{} {
	var user entities.User
	database.Db.Find(&user, "email = ?", email)
	return user
}

func FindUserById(id string) interface{} {
	var user entities.User
	database.Db.Find(&user, "id = ?", id)
	return user
}

func UserExists(email string) *gorm.DB {
	return database.Db.Take(&entities.User{}, "email = ?", email)
}
