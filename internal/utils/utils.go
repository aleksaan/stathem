package utils

import (
	"context"
	"encoding/json"

	"github.com/aleksaan/stathem/internal/auth"
	apperrors "github.com/aleksaan/stathem/internal/errors"
)

func ConvertJsonToStruct[T any](js []byte) (*T, error) {
	var result *T

	err := json.Unmarshal(js, &result)
	if err != nil {
		return nil, apperrors.ErrWrongJsonFormat
	}
	return result, nil
}

func ConvertStructToJson(T any) []byte {

	json, _ := json.Marshal(&T)

	return json
}

func GetUserNameFromContext(ctx context.Context) string {
	userName, _ := ctx.Value(auth.UserNameContextKey).(string)

	return userName
}

func GetUserRolesFromContext(ctx context.Context) []string {

	userRoles, _ := ctx.Value(auth.RolesContextKey).([]string)

	return userRoles
}

func InitSystemContext(ctx context.Context) context.Context {
	ctx = context.WithValue(ctx, auth.UserNameContextKey, auth.SystemUser)
	return ctx
}
