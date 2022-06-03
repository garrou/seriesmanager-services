package dto

type UserDto struct {
	Email    string `json:"email" binding:"required" validate:"email,max:255"`
	Password string `json:"password" binding:"required" validate:"min:8,max:50"`
}

type UserCreateDto struct {
	Email    string `json:"email" binding:"required" validate:"email,max:255"`
	Password string `json:"password" binding:"required" validate:"min:8,max:50"`
	Confirm  string `json:"confirm" binding:"required" validate:":min:8,max:50"`
}

type UserUpdateDto struct {
	Id       string
	Email    string `json:"email" binding:"required" validate:"email,max:255"`
	Password string `json:"password" validate:"min:8,max:50"`
}
