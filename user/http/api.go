package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

type ServerInterface interface {
	HealthCheck(c *fiber.Ctx) error
}

type HTTPWrapper struct {
	Handler ServerInterface
}

func (w *HTTPWrapper) HealthCheck(c *fiber.Ctx) error {

	if err := w.Handler.HealthCheck(c); err != nil {
		log.Warnf("w.Handler.HealthCheck error occurred: %s", err.Error())
		return err
	}

	return nil
}
func UserHandler(app *fiber.App, si ServerInterface) {
	wrapper := HTTPWrapper{
		Handler: si,
	}

	route := app.Group("")
	route.Get("/healthcheck", wrapper.HealthCheck)
}
