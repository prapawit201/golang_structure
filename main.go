package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"poc.com/user/http"
	"poc.com/user/repository"
	"poc.com/user/service"
)

var (
	app = fiber.New()
)

func main() {

	//init repository
	userRepository := repository.NewUserRepository()

	//init service
	userService := service.NewUserService(userRepository)

	//init http
	httpHandler := http.NewHTTPHandler(userService)
	http.UserHandler(app, httpHandler)

	log.Fatal(app.Listen(":3000"))
}
