package controllers

import (
	"net/http"
	"user-api-specification/models"

	"github.com/gin-gonic/gin"
)

// CreateUser creates a new user.
// @Summary Create a new user
// @Description Add a new user to the database
// @Accept json
// @Produce json
// @Param user body models.UserDTO true "User data"
// @Success 200 {object} models.UserDTO
// @Router /user [post]
func CreateUser(c *gin.Context) {
	var user models.UserDTO
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Here you would add code to save the user to a database.
	c.JSON(http.StatusOK, user)
}

// RegisterUser registers a new user in the microservice's database.
// @Summary Register a new user
// @Description Register a new user and save their information in the database
// @Accept json
// @Produce json
// @Param user body models.UserDTO true "User registration data"
// @Success 201 {object} models.UserDTO
// @Router /register [post]
func RegisterUser(c *gin.Context) {
	var user models.UserDTO
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Simulate registering the user
	user.ID = "newly-generated-id" // Simulate ID generation for the new user
	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully", "user": user})
}

// GetUser retrieves a user by ID.
// @Summary Get user by ID
// @Description Get a user by their ID
// @Param id path string true "User ID"
// @Success 200 {object} models.UserDTO
// @Router /user/{id} [get]
func GetUser(c *gin.Context) {
	id := c.Param("id")
	// Here you would fetch the user from the database.
	c.JSON(http.StatusOK, models.UserDTO{ID: id, Username: "Example User", Email: "example@example.com"})
}

// GetUserForKeycloak retrieves a user's info specifically for Keycloak's use.
// @Summary Get user info for Keycloak
// @Description Retrieve minimal user information for Keycloak authentication purposes
// @Param id path string true "User ID"
// @Success 200 {object} models.UserKeycloakDTO
// @Router /user/keycloak/{id} [get]
func GetUserForKeycloak(c *gin.Context) {
	id := c.Param("id")

	// Fetch minimal user data for Keycloak authentication
	// In a real implementation, you would fetch from the database
	user := models.UserKeycloakDTO{
		ID:       id,
		Username: "example_username",
		Roles:    []string{"user", "admin"}, // Example roles
		GroupID:  "group_id"}

	c.JSON(http.StatusOK, user)
}

// UpdateUser updates an existing user by ID.
// @Summary Update user by ID
// @Description Update user details
// @Param id path string true "User ID"
// @Param user body models.UserDTO true "User data"
// @Success 200 {object} models.UserDTO
// @Router /user/{id} [put]
func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var user models.UserDTO
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user.ID = id
	// Here you would update the user in the database.
	c.JSON(http.StatusOK, user)
}

// ResetPassword handles the password reset process for a user.
// @Summary Reset a user's password
// @Description Initiates a password reset process for a user by their email or ID
// @Accept json
// @Produce json
// @Param resetRequest body models.ResetPasswordDTO true "Password reset request data"
// @Success 200 {string} string "Password reset initiated"
// @Router /resetpassword [post]
func ResetPassword(c *gin.Context) {
	var resetRequest models.ResetPasswordDTO
	if err := c.BindJSON(&resetRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Simulate resetting the password
	c.JSON(http.StatusOK, gin.H{"message": "Password reset successful for user"})
}

// DeleteUser deletes a user by ID.
// @Summary Delete user by ID
// @Description Delete user by their ID
// @Param id path string true "User ID"
// @Success 200 {string} string "Deleted successfully"
// @Router /user/{id} [delete]
func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	// Here you would delete the user from the database using the id.
	// For example: db.DeleteUserByID(id)
	c.JSON(http.StatusOK, gin.H{"message": "User with ID " + id + " deleted successfully"})
}
