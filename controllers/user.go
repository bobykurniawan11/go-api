package controllers

import (
	"net/http"

	"github.com/bobykurniawan11/starter-go/models"
	"github.com/gin-gonic/gin"
)

type UserController struct{}
type InputUser struct {
	Name            string `json:"name" binding:"required"`
	Email           string `json:"email" binding:"required,email"`
	Password        string `json:"password" binding:"required,min=8"`
	PasswordConfirm string `json:"password_confirm" binding:"required,eqfield=Password"`
}

type UpdateUser struct {
	//NAme Optional
	Name string `json:"name"`
	//Email Optional

	Phone string `json:"phone"`
}

// GetAll retrieves all users from the database and returns them as a JSON response.
// It handles the GET request to the "/users" endpoint.
// If an error occurs during the retrieval process, it returns a JSON response with the error message and a 500 status code.
// If the retrieval is successful, it returns a JSON response with the list of users and a 200 status code.
func (u UserController) GetAll(c *gin.Context) {
	var userModel = new(models.User)

	users, err := userModel.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"users": users})
}

// CreateUser creates a new user based on the input provided in the request body.
// It binds the request body to the InputUser struct, validates the input,
// and then calls the CreateUser method of the userModel to create the user.
// If successful, it returns the created user in the response body.
// If there is an error during the process, it returns the appropriate error response.
func (u UserController) CreateUser(c *gin.Context) {
	var userModel = new(models.User)

	var form InputUser
	// This will infer what binder to use depending on the content-type header.
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, error := userModel.CreateUser(&models.User{
		Name:     form.Name,
		Email:    form.Email,
		Password: form.Password,
	})
	if error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

// GetUserById retrieves a user by their ID.
// It takes a gin.Context object as a parameter and uses the ID parameter from the request URL to fetch the user from the database.
// If the user is found, it returns the user details as a JSON response with HTTP status code 200.
// If there is an error during the retrieval process, it returns an error message as a JSON response with HTTP status code 500.
func (u UserController) GetUserById(c *gin.Context) {
	var userModel = new(models.User)

	id := c.Param("id")
	user, err := userModel.GetUserById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

// UpdateUser updates the user information based on the provided data.
// It retrieves the existing user from the database, updates the specified fields,
// and saves the updated user back to the database.
// If any error occurs during the process, it returns the corresponding error response.
func (u UserController) UpdateUser(c *gin.Context) {
	var userModel = new(models.User)

	id := c.Param("id")

	// Define a struct to hold the fields you want to update
	var updateFields UpdateUser

	// This will infer what binder to use depending on the content-type header.
	if err := c.ShouldBind(&updateFields); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Retrieve the existing user from the database
	_user, err := userModel.GetUserById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Update the user fields based on the provided data
	_user.Name = updateFields.Name
	_user.Phone = updateFields.Phone

	// Save the updated user back to the database
	user, err := userModel.UpdateUser(id, &models.User{
		Name:     _user.Name,
		Password: _user.Password,
		Email:    _user.Email,
		ID:       _user.ID,
		Phone:    _user.Phone,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})

}

// DeleteUser deletes a user based on the provided ID.
// It retrieves the user by ID, deletes it from the database, and returns a JSON response with the deleted user's information.
// If any error occurs during the process, an error JSON response is returned.
func (u UserController) DeleteUser(c *gin.Context) {
	var userModel = new(models.User)

	id := c.Param("id")
	User, err := userModel.GetUserById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = userModel.DeleteUser(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success delete user", "user": User})
}
