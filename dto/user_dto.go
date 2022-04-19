package dto

type UserDto struct {
	Email    string `json:"email" form:"email" binding:"required" validate:"email,max:255"`
	Password string `json:"password" form:"password" binding:"required" validate:"min:8,max:255"`
}

type UserCreateDto struct {
	Email    string `json:"email" form:"email" binding:"required" validate:"email,max:255"`
	Password string `json:"password" form:"password" binding:"required" validate:"min:8,max:255"`
	Confirm  string `json:"confirm" form:"confirm" binding:"required" validate:":min:8,max:255"`
}

type UserUpdateDto struct {
	Id       string `json:"omitempty"`
	Email    string `json:"email" form:"email" binding:"required" validate:"email,max:255"`
	Password string `json:"password" form:"password" validate:"min:8,max:255"`
}
