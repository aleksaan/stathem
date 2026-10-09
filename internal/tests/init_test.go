package tests

import (
	"context"

	"github.com/aleksaan/stathem/internal/auth"
	"github.com/aleksaan/stathem/internal/config"
	"github.com/aleksaan/stathem/internal/database"
)

var pathConfig = "../../config.yml"

var ctx context.Context

func init() {
	config.LoadConfig(pathConfig)
	database.InitDBConnection()
	ctx = context.Background()
	// ctx = context.WithValue(ctx, auth.RolesContextKey, auth.RoleAdmin)
	ctx = context.WithValue(ctx, auth.UserNameContextKey, auth.SystemUser)
}
