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

func (u UserController) GetAll(c *gin.Context) {
	var userModel = new(models.User)

	users, err := userModel.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"users": users})
}

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
