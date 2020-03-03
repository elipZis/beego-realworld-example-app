package services

import (
	"errors"
	"github.com/astaxie/beego/orm"
	"github.com/elipZis/beego-realworld-example-app/models"
	"github.com/elipZis/beego-realworld-example-app/repositories"
)

type UserService struct {
	*JwtService
	*repositories.UserRepository
}

// Validate a user, login and store the token
func (service *UserService) Login(user *models.User) (string, error) {
	ormer := orm.NewOrm()

	// Validate the credentials and log-in the user
	err := ormer.Read(&user, "Email", "Password")
	if err != nil {
		return "", errors.New("error.user.wrong_credentials")
	}

	// Create a token, store the uid in the session
	token, err := service.JwtService.GenerateToken(user)
	if err != nil {
		return "", err
	}

	// Add the token to the user object for better referencing and usage
	user.Token = token

	return token, nil
}

// Register a user in the database and login the model
func (service *UserService) Register(user *models.User) error {
	ormer := orm.NewOrm()

	// Insert and check for duplicate emails or other errors
	id, err := ormer.Insert(user)
	if err != nil {
		return err
	}

	// Update the model with the newly inserted Id
	user.Id = id

	// Login and update the model with the correct token
	_, error := service.Login(user)
	return error
}
