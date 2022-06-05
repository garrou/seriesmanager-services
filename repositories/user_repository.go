package repositories

import (
	"seriesmanager-services/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	Save(user models.User) models.User
	FindByEmail(email string) interface{}
	FindById(id string) interface{}
	Exists(email string) (tx *gorm.DB)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (u *userRepository) Save(user models.User) models.User {
	u.db.Save(&user)
	return user
}

func (u *userRepository) FindByEmail(email string) interface{} {
	var user models.User
	res := u.db.Find(&user, "email = ?", email)

	if res.Error == nil {
		return user
	}
	return nil
}

func (u *userRepository) FindById(id string) interface{} {
	var user models.User
	res := u.db.Find(&user, "id = ?", id)

	if res.Error == nil {
		return user
	}
	return nil
}

func (u *userRepository) Exists(email string) *gorm.DB {
	var user models.User
	return u.db.Take(&user, "email = ?", email)
}
