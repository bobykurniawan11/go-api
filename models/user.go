package models

import (
	"time"

	"github.com/bobykurniawan11/starter-go/db"
	"github.com/bobykurniawan11/starter-go/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	//ID UUID
	ID uuid.UUID `gorm:"size:36" json:"id"`
	//Name string
	Name string `gorm:"size:255" json:"name"`
	//Email string
	Email string `gorm:"size:255,unique" json:"email"`
	//Password string
	Password string `gorm:"size:255" json:"-"`

	//Avatar string
	Avatar string `gorm:"size:255" json:"avatar"`

	//Phone string
	Phone string `gorm:"size:255" json:"phone"`

	//CreatedAt time.Time
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	//UpdatedAt time.Time
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

//createdAt to gmt+0

func (u *User) GetCreatedAtFormatted() time.Time {
	return u.CreatedAt.UTC()
}

// updatedAt to gmt+0
func (u *User) GetUpdatedAtFormatted() time.Time {
	return u.UpdatedAt.UTC()
}

func (user *User) BeforeCreate(_ *gorm.DB) (err error) {
	// UUID version 4
	user.ID = uuid.New()
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword
	return nil
}

// CreateUser func
func (u *User) CreateUser(user *User) (User, error) {
	userModel := db.GetDB().Create(&user)
	if userModel.Error != nil {
		return User{}, userModel.Error
	}

	return *user, nil
}

// GetAll func
func (u *User) GetAll() ([]User, error) {
	var users []User
	user := db.GetDB().Find(&users)
	if user.Error != nil && user.Error != gorm.ErrRecordNotFound {
		return nil, user.Error
	}

	for i := range users {
		users[i].CreatedAt = users[i].GetCreatedAtFormatted()
		users[i].UpdatedAt = users[i].GetUpdatedAtFormatted()
	}

	return users, nil

}

func (u *User) GetUserById(id string) (User, error) {
	var user User
	userModel := db.GetDB().Where("id = ?", id).Find(&user)
	if userModel.Error != nil {
		return User{}, userModel.Error
	}

	user.CreatedAt = user.GetCreatedAtFormatted()
	user.UpdatedAt = user.GetUpdatedAtFormatted()

	return user, nil
}

func (u *User) GetUserByEmail(email string) (User, error) {
	var user User
	userModel := db.GetDB().Where("email = ?", email).Find(&user)
	if userModel.Error != nil {
		return User{}, userModel.Error
	}

	user.CreatedAt = user.GetCreatedAtFormatted()
	user.UpdatedAt = user.GetUpdatedAtFormatted()

	return user, nil
}

func (u *User) DeleteUser(id string) error {
	userModel := db.GetDB().Where("id = ?", id).Delete(&User{})
	if userModel.Error != nil {
		return userModel.Error
	}

	return nil
}

func (u *User) UpdateUser(id string, _user *User) (User, error) {
	// Retrieve the existing user from the database
	existingUser, err := u.GetUserById(id)
	if err != nil {
		return User{}, err
	}

	// Update the user fields based on the provided data
	existingUser.Name = _user.Name
	existingUser.Email = _user.Email
	existingUser.Avatar = _user.Avatar

	// Save the updated user back to the database
	if err := db.GetDB().Save(&existingUser).Error; err != nil {
		return User{}, err
	}

	existingUser.CreatedAt = existingUser.GetCreatedAtFormatted()
	existingUser.UpdatedAt = existingUser.GetUpdatedAtFormatted()

	return existingUser, nil
}
