package http

import "poc.com/user/service"

type Handler struct {
	UserService service.UserService
}

func NewHTTPHandler(userService service.UserService) ServerInterface {
	return &Handler{
		UserService: userService,
	}
}
