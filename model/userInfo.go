package model

type userInfo struct {
	ID int `json:"id"`
	Email string `json:"email"`
	UserRole string `json:"user_role"`
}
