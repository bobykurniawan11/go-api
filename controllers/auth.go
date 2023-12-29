package controllers

import (
	"log"
	"net/http"

	"github.com/bobykurniawan11/starter-go/models"
	"github.com/bobykurniawan11/starter-go/utils"

	"os"

	"github.com/gin-gonic/gin"
)

// LoginUser represents the structure of the user login request.
// It contains the user's email and password.
type LoginUser struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

// RegisterUser represents the data structure for registering a new user.
// It contains the user's name, email, password, and password confirmation.
// The fields are tagged with JSON tags for serialization and binding tags for validation.
type RegisterUser struct {
	Name            string `json:"name" binding:"required"`
	Email           string `json:"email" binding:"required,email"`
	Password        string `json:"password" binding:"required,min=8"`
	PasswordConfirm string `json:"password_confirm" binding:"required,eqfield=Password"`
}

type UploadAvatar struct {
	Avatar string `json:"avatar" binding:"required"`
}

type AuthController struct{}

// SignIn handles the sign-in functionality.
// It receives a POST request with user credentials and validates them.
// If the credentials are valid, it generates a token and returns it in the response.
// If there is any error during the sign-in process, it returns an appropriate error response.
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

// SignUp handles the sign-up request and creates a new user.
// It binds the request body to the RegisterUser struct and validates the input.
// If the input is valid, it creates a new user using the CreateUser method of the User model.
// It then generates a token for the user using the GenerateToken function from the utils package.
// Finally, it returns the generated token in the response.
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

// Me is a handler function for the "/me" endpoint.
// It retrieves the user information based on the token ID extracted from the request context.
// If the token ID is valid, it returns the user information in the response.
// If there is an error during the process, it returns an error response with the corresponding error message.
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

// UploadAvatar handles the file upload for the user's avatar.
// It expects a single file with the key "file" in the request form data.
// If the file is missing, it returns a JSON response with a "File is required" error.
// If there is an error while receiving the file, it returns a JSON response with the error message.
// The uploaded file is saved to the "./uploads" directory with the original filename.
// It then extracts the user ID from the request token using the utils.ExtractTokenID function.
// It checks if the uploaded file is a valid image using the utils.CheckImage function.
// If the image is invalid, it removes the uploaded file, returns a JSON response with an "Invalid image" error.
// If there is an error while extracting the user ID or checking the image, it returns a JSON response with the error message.
// It retrieves the user by ID using the models.User.GetUserById method.
// It updates the user's avatar filename with the uploaded file's filename.
// Finally, it returns a JSON response with the updated user object.
func (u AuthController) UploadAvatar(c *gin.Context) {
	// single file
	file, err := c.FormFile("file")
	switch err {
	case nil:
		// do nothing
	case http.ErrMissingFile:
		c.JSON(http.StatusBadRequest, gin.H{"error": "File is required"})
		return
	default:
		log.Printf("Error while receiving the file: %s\n", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	// Upload the file to specific dst.
	c.SaveUploadedFile(file, "./uploads/"+file.Filename)

	id, err := utils.ExtractTokenID(c)

	imageStatus, errImage := utils.CheckImage("./uploads/" + file.Filename)

	if errImage != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errImage.Error()})
		return
	}

	if !imageStatus {
		os.Remove("./uploads/" + file.Filename)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid image"})
		return

	}

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
