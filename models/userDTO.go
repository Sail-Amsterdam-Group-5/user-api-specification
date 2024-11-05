package models

// User struct represents a user in the system.
// @Description User model
type UserDTO struct {
	ID       string `json:"id" example:"123"`
	Username string `json:"username" example:"John Doe"`
	Email    string `json:"email" example:"john@example.com"`
	Role     string `json:"role" example:"volunteer"`
	Phone    string `json:"phone" example:"1234567890"`
	GroupID  string `json:"groupID" example:"2"`
}
