package http

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

// HealthCheck implements ServerInterface.
func (h *Handler) HealthCheck(c *fiber.Ctx) error {
	resp, err := h.UserService.HealthCheck(c)
	if err != nil {
		log.Printf("h.UserService.HealthCheck Error Occurred: %s", err.Error())

	}

	c.JSON(resp)
	return nil
}
