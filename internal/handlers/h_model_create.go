package handlers

import (
	"context"
	"net/http"

	"github.com/aleksaan/stathem/internal/actions"
	"github.com/aleksaan/stathem/internal/config"
	"github.com/aleksaan/stathem/internal/database"
	"github.com/aleksaan/stathem/internal/requests"
	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

func HandlerModelCreate(c *echo.Context) error {
	m := &requests.ModelCreateRequest{}

	// Bind the request body to the User struct
	if err := c.Bind(m); err != nil {
		// Return a 400 Bad Request error if binding fails
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	// It's good practice to close the body after reading
	defer c.Request().Body.Close()

	ctx, cancel := context.WithTimeout(c.Request().Context(), config.Config.DbQueryTimeout)
	defer cancel()

	err := database.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := actions.ModelCreate(ctx, tx, m)
		return err
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.NoContent(http.StatusOK)
}
