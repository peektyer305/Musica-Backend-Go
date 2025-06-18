package handlers

import (
	usecase "Musica-Backend/internal/application/user"
	valueobject "Musica-Backend/internal/domain/value_object"
	"Musica-Backend/internal/presentation/rest/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	UserUseCase *usecase.UserUseCase
}

func (u *UserHandler) FindUserById(ctx echo.Context) (response.UserResponse, error) {
	id := ctx.Param("id")
	userId, err := valueobject.NewUserId(id)
	if err != nil {
		return response.UserResponse{}, echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	user, err := u.UserUseCase.FindUserById(ctx.Request().Context(), userId)
	if err != nil {
		return response.UserResponse{}, echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return response.DomainToResponseUser(user), nil
}
