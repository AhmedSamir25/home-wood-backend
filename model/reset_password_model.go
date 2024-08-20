package model

type PasswordReset struct {
	Id         uint   `json:"id"`
	Email      string `json:"email"`
	ResetToken int    `json:"token"`
}
