package handler

import (
	"net/http"

	"github.com/Sandhya-Pratama/weather-app/internal/service"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService service.UserUseCase
}

func NewUserHandler(userService service.UserUseCase) *UserHandler   {
	return &UserHandler{userService}
}

func (h *UserHandler) GetAllUsers(ctx echo.Context) error {
	users, err := h.userService.FindAll(ctx.Request().Context())

	if err != nil{
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}
	return ctx.JSON(http.StatusOK, users)
}