package handlers

import (
	"context"
	"net/http"

	"github.com/aleksaan/stathem/internal/actions"
	"github.com/aleksaan/stathem/internal/config"
	"github.com/aleksaan/stathem/internal/database"
	"github.com/aleksaan/stathem/internal/entities/processes"
	"github.com/aleksaan/stathem/internal/requests"
	"github.com/aleksaan/stathem/internal/utils"
	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

func HandlerProcessCreate(c *echo.Context) error {
	pr := &requests.ProcessCreateRequest{}

	// Bind the request body to the User struct
	if err := c.Bind(pr); err != nil {
		// Return a 400 Bad Request error if binding fails
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	ctx, cancel := context.WithTimeout(c.Request().Context(), config.Config.DbQueryTimeout)
	defer cancel()

	p := &processes.Process{}
	err := database.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) (err1 error) {
		p, err1 = actions.ProcessCreate(ctx, tx, pr)
		return err1
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	json := utils.ConvertStructToJson(actions.ProcessCreateResponse(p))

	return c.JSONBlob(http.StatusOK, json)

}
