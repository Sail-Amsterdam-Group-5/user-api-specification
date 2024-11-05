package models

// ResetPasswordDTO struct represents a user's request to reset their password.
// @Description Reset password request data
type ResetPasswordDTO struct {
	Email string `json:"email" example:"forgetfuluser@email.com"`
	ID string `json:"id" example:"123"`
	Password string `json:"password" example:"newpassword"`

}