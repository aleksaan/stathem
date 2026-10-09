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

func HandlerProcessFinish(c *echo.Context) error {
	p := &requests.ProcessFinishRequest{}

	// Bind the request body to the User struct
	if err := c.Bind(p); err != nil {
		// Return a 400 Bad Request error if binding fails
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	ctx, cancel := context.WithTimeout(c.Request().Context(), config.Config.DbQueryTimeout)
	defer cancel()

	err := database.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) (err1 error) {
		err1 = actions.ProcessFinish(ctx, tx, p)
		return err1
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.NoContent(http.StatusOK)

}
