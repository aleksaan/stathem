package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"

	"github.com/aleksaan/stathem/internal/auth"
	"github.com/aleksaan/stathem/internal/config"
	"github.com/aleksaan/stathem/internal/database"
	"github.com/aleksaan/stathem/internal/handlers"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

func run() error {

	e := echo.New()
	e.Use(middleware.RequestLogger())

	config.LoadConfig("./config.yml")
	database.InitDBConnection()

	e.Use(middleware.BodyLimit(config.Config.AppLimitBodySizeInBytes))

	// Переопределяем поведение для всех неправильных URL
	e.RouteNotFound("/*", func(c *echo.Context) error {
		// Здесь вы задаете абсолютно любое новое сообщение и формат
		return c.JSON(http.StatusNotFound, map[string]string{
			"message": "Error: URL is not found",
		})
	})

	// пользователи
	err := auth.LoadUsersFromFile("./users.txt")
	if err != nil {
		log.Fatalf("Ошибка загрузки файла пользователей: %v", err)
	}
	fmt.Printf("Успешно загружено пользователей: %d\n", len(auth.UserDB))

	//restricted group of routes
	r := e.Group("/api")
	r.Use(middleware.BasicAuthWithConfig(middleware.BasicAuthConfig{
		Validator: auth.BasicAuthValidator,
		Realm:     "Restricted Access",
	}))

	e.GET("/", handlers.HandlerHelloWorld)
	r.POST("/createModel", handlers.HandlerModelCreate, auth.RequireRoles("createModel"))
	r.POST("/createProcess", handlers.HandlerProcessCreate, auth.RequireRoles("createProcess"))
	r.POST("/getModel", handlers.HandlerModelGet, auth.RequireRoles("getModel"))
	r.POST("/getModels", handlers.HandlerModelsGet, auth.RequireRoles("getModels"))
	r.POST("/finishProcess", handlers.HandlerProcessFinish, auth.RequireRoles("finishProcess"))
	r.POST("/getProcess", handlers.HandlerProcessGet, auth.RequireRoles("getProcess"))
	r.POST("/getProcesses", handlers.HandlerProcessesGet, auth.RequireRoles("getProcesses"))
	r.POST("/setState", handlers.HandlerStateSet, auth.RequireRoles("setState"))
	r.POST("/getStates", handlers.HandlerStatesGet, auth.RequireRoles("getStates"))

	slog.Info(fmt.Sprintf("Start application on port %s ...", config.Config.AppPort))
	if err := e.Start(":" + config.Config.AppPort); err != nil {
		slog.Error(fmt.Sprintf("failed to start server: %s", err.Error()))
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
