package dto

import (
	"strings"
	"time"
)

// UserLoginDto represents a user during login
type UserLoginDto struct {
	Email    string `json:"email" binding:"required" validate:"email,max:255"`
	Password string `json:"password" binding:"required" validate:"min:8,max:50"`
}

// UserCreateDto represents a user during register
type UserCreateDto struct {
	Email    string `json:"email" binding:"required" validate:"email,min:8,max:255"`
	Username string `json:"username" binding:"required" validate:"min:3,max:50"`
	Password string `json:"password" binding:"required" validate:"min:8,max:50"`
	Confirm  string `json:"confirm" binding:"required" validate:":min:8,max:50"`
}

// TrimSpace spaces on user information
func (u *UserCreateDto) TrimSpace() {
	u.Email = strings.TrimSpace(u.Email)
	u.Username = strings.TrimSpace(u.Username)
	u.Password = strings.TrimSpace(u.Password)
	u.Confirm = strings.TrimSpace(u.Confirm)
}

// IsValid checks if user information is valid
func (u *UserCreateDto) IsValid() bool {
	return len(u.Password) >= 8 && u.Password == u.Confirm && len(u.Username) >= 3
}

// UserDto represents a user profile
type UserDto struct {
	Username string    `json:"username"`
	Email    string    `json:"email"`
	JoinedAt time.Time `json:"joinedAt"`
	Banner   string    `json:"banner"`
}

// UserUpdateProfileDto information send to update user profile
type UserUpdateProfileDto struct {
	Id       string
	Username string `json:"username" binding:"required" validate:"min:3"`
	Email    string `json:"email" binding:"required" validate:"email,max:255"`
}

// UserUpdatePasswordDto information send to update user password
type UserUpdatePasswordDto struct {
	Id              string
	CurrentPassword string `json:"current" binding:"required" validate:"min:8,max:50"`
	Password        string `json:"password" binding:"required" validate:"min:8,max:50"`
	Confirm         string `json:"confirm" binding:"required" validate:"min:8,max:50"`
}
