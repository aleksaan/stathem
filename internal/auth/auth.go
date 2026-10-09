package auth

import (
	"bufio"
	"context"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v5"
	"golang.org/x/crypto/bcrypt"
)

// Константы для ролей и контекста
const (
	SystemUser         = "_system_"
	RolesContextKey    = "user_roles"
	UserNameContextKey = "user_name"
)

// User хранит список ролей в виде слайса строк
type User struct {
	Username string
	Password string
	Roles    []string // Теперь здесь массив (слайс)
	jwt.RegisteredClaims
}

// type MyCustomClaims struct {
// 	UserID   int    `json:"user_id"`
// 	Role     string `json:"role"`
// 	jwt.RegisteredClaims
// }

// Кэш пользователей в памяти
var UserDB = make(map[string]User)

// loadUsersFromFile парсит текстовый файл с поддержкой нескольких ролей
func LoadUsersFromFile(filepath string) error {
	file, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.Split(line, ",")
		if len(parts) != 3 {
			log.Printf("Пропущена некорректная строка: %s", line)
			continue
		}

		username := strings.TrimSpace(parts[0])
		password := strings.TrimSpace(parts[1])

		// Разбиваем строку ролей (например, "admin:manager") по разделителю ":"
		rawRoles := strings.Split(strings.TrimSpace(parts[2]), ":")
		var roles []string
		for _, r := range rawRoles {
			trimmedRole := strings.TrimSpace(r)
			if trimmedRole != "" {
				roles = append(roles, trimmedRole)
			}
		}

		UserDB[username] = User{
			Username: username,
			Password: password,
			Roles:    roles,
		}
	}

	return scanner.Err()
}

// Basic Auth Validator Middleware
func BasicAuthValidator(c *echo.Context, username, password string) (bool, error) {
	u, exists := UserDB[username]
	if !exists {
		return false, nil
	}

	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		return false, nil
	}

	//passwordMatch := subtle.ConstantTimeCompare([]byte(password), []byte(u.Password)) == 1

	// Записываем массив ролей в контекст запроса
	ctx := context.WithValue(c.Request().Context(), RolesContextKey, u.Roles)
	ctx = context.WithValue(ctx, UserNameContextKey, u.Username)

	// Мутируем текущий HTTP-запрос, добавив в него новый контекст
	c.SetRequest(c.Request().WithContext(ctx))

	return true, nil
}

// RequireRoles Middleware проверяет, есть ли у пользователя ХОТЯ БЫ ОДНА из разрешенных ролей
func RequireRoles(allowedRoles ...string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {
			// Получаем массив ролей пользователя из контекста
			userRoles, ok := c.Request().Context().Value(RolesContextKey).([]string)
			if !ok {
				return echo.NewHTTPError(http.StatusForbidden, "Доступ запрещен: отсутствует контекст ролей")
			}

			// Проверяем пересечение ролей пользователя и разрешенных ролей для эндпоинта
			for _, allowedRouteRole := range allowedRoles {
				for _, userRole := range userRoles {
					if userRole == allowedRouteRole {
						return next(c) // Роль совпала, пускаем дальше
					}
				}
			}

			return echo.NewHTTPError(http.StatusForbidden, "Доступ запрещен: недостаточно прав")
		}
	}
}
