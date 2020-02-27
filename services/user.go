package services

import "github.com/elipZis/beego-realworld-example-app/repositories"

type UserService struct {
	*JwtService
	*repositories.UserRepository
}

func (service *UserService) Login(email string, password string) {
	// Validate the credentials and log-in the user

	// Create a token, store the uid in the session
}
