package handlers

import (
	"Musica-Backend/internal/domain/auth"

	"github.com/labstack/echo/v4"
)

func GetMyProfile(c echo.Context) error {
	user := c.Get("currentUser").(*auth.User)
	return c.JSON(200, user)
}
