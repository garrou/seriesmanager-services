package repositories

import (
	"services-series-manager/models"

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
	res := u.db.Where("email = ?", email).Take(&user)

	if res.Error == nil {
		return user
	}
	return nil
}

func (u *userRepository) FindById(id string) interface{} {
	var user models.User
	res := u.db.Where("id = ?", id).Take(&user)

	if res.Error == nil {
		return user
	}
	return nil
}

func (u *userRepository) Exists(email string) *gorm.DB {
	var user models.User
	return u.db.Where("email = ?", email).Take(&user)
}
