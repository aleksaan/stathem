package handlers

import (
	"net/http"

	"github.com/labstack/echo/v5"
)

func HandlerHelloWorld(c *echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
