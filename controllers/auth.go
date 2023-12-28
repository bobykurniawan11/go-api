package controllers

import (
	"net/http"

	"github.com/bobykurniawan11/starter-go/models"
	"github.com/bobykurniawan11/starter-go/utils"

	"github.com/gin-gonic/gin"
)

type LoginUser struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type RegisterUser struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type UploadAvatar struct {
	Avatar string `json:"avatar" binding:"required"`
}

type AuthController struct{}

func (u AuthController) SignIn(c *gin.Context) {
	var userModel = new(models.User)

	var form LoginUser

	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, error := userModel.GetUserByEmail(form.Email)

	if error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": error.Error()})
		return
	}

	if !utils.CheckPasswordHash(form.Password, user.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid password"})
		return
	}

	token, err := utils.GenerateToken(
		user.ID,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})

}

func (u AuthController) SignUp(c *gin.Context) {
	var userModel = new(models.User)

	var form RegisterUser
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
	token, err := utils.GenerateToken(
		user.ID,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (u AuthController) Me(c *gin.Context) {
	var userModel = new(models.User)

	//make as uuid tokenString

	id, err := utils.ExtractTokenID(c)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	user, err := userModel.GetUserById(id.String())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

func (u AuthController) UploadAvatar(c *gin.Context) {
	// single file
	file, _ := c.FormFile("file")

	// Upload the file to specific dst.
	c.SaveUploadedFile(file, "./uploads/"+file.Filename)

	id, err := utils.ExtractTokenID(c)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var userModel = new(models.User)

	user, err := userModel.GetUserById(id.String())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	user.Avatar = file.Filename

	userModel.UpdateUser(id.String(), &user)

	c.JSON(http.StatusOK, gin.H{"user": user})
}

func (u AuthController) LogoutUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "logout"})
}
