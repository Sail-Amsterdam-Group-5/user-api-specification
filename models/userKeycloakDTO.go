package models

// UserKeycloakDTO struct to pass info to Keycloak
// @Description User model for Keycloak

type UserKeycloakDTO struct {
	ID       string   `json:"id" example:"123"`
	Username string   `json:"username" example:"John Doe"`
	Roles    []string `json:"roles" example:"volunteer"`
	GroupID string   `json:"groupID" example:"2"`
}