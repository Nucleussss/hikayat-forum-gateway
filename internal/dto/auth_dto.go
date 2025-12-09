package dto

// POST /register
type RegisterRequest struct {
	Name     string `json:"name" binding:"required,min=2,max=50"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8,max=128"`
}

// POST /login
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8,max=128"`
}

type GetUserRequest struct {
	Id string `json:"id" binding:"required"`
}

// PATCH /profile
type UpdateUserProfileRequest struct {
	Name string `json:"name" binding:"required,min=2,max=50"`
	Id   string `json:"id" binding:"required"`
}

// PATCH /profile/email
type ChangeUserEmailRequest struct {
	Email string `json:"email" binding:"required,email"`
	Id    string `json:"id" binding:"required"`
}

// PATCH /profile/password
type ChangeUserPasswordRequest struct {
	Id              string `json:"id" binding:"required"`
	CurrentPassword string `json:"current_password" binding:"required,min=8,max=128"`
	NewPassword     string `json:"new_password" binding:"required,min=8,max=128,nefield=CurrentPassword"`
}
