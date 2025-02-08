package service

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"poc.com/user/entity"
)

// HealthCheck implements UserService.
func (u *userService) HealthCheck(c *fiber.Ctx) (interface{}, error) {

	if err := u.UserRepository.GetAll(); err != nil {
		log.Printf("u.UserRepository.GetAll error: %s", err.Error())
		return nil, err
	}

	return &entity.HealthCheck{
		Code:    http.StatusOK,
		Message: "Success",
	}, nil
}
