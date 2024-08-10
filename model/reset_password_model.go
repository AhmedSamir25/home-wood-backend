package model

type PasswordReset struct {
	Id         uint   `json:"id"`
	Email      string `json:"user_email"`
	ResetToken int    `json:"token"`
}
