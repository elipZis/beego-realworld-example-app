package services

import "github.com/elipZis/beego-realworld-example-app/repositories"

type UserService struct {
	*JwtService
	*repositories.UserRepository
}
