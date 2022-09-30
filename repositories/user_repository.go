package repositories

import (
	"gorm.io/gorm"
	"seriesmanager-services/entities"
)

type UserRepository interface {
	Save(user entities.User) entities.User
	FindByEmail(email string) interface{}
	FindById(id string) interface{}
	Exists(email string) *gorm.DB
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (u *userRepository) Save(user entities.User) entities.User {
	u.db.Save(&user)
	return user
}

func (u *userRepository) FindByEmail(email string) interface{} {
	var user entities.User
	u.db.Find(&user, "email = ?", email)
	return user
}

func (u *userRepository) FindById(id string) interface{} {
	var user entities.User
	u.db.Find(&user, "id = ?", id)
	return user
}

func (u *userRepository) Exists(email string) *gorm.DB {
	return u.db.Take(&entities.User{}, "email = ?", email)
}
