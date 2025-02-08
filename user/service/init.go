package service

import (
	"github.com/gofiber/fiber/v2"
	"poc.com/user/repository"
)

//go:generate /Users/babe/go/bin/mockgen -source=init.go -destination=mocks/user.go

type UserService interface {
	HealthCheck(c *fiber.Ctx) (interface{}, error)
}

type userService struct {
	UserRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{
		UserRepository: userRepository,
	}
}
