package handlers

import (
	"context"
	"net/http"

	"github.com/aleksaan/stathem/internal/actions"
	"github.com/aleksaan/stathem/internal/config"
	"github.com/aleksaan/stathem/internal/database"
	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

func HandlerModelsGet(c *echo.Context) error {

	ctx, cancel := context.WithTimeout(c.Request().Context(), config.Config.DbQueryTimeout)
	defer cancel()

	var json []byte
	err := database.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) (err1 error) {
		json, err1 = actions.ModelsGet(ctx, tx)
		return err1
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSONBlob(http.StatusOK, json)
}
