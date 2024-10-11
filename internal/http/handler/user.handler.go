package handler

import (
	"net/http"

	"github.com/Sandhya-Pratama/weather-app/entity"
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

func (h *UserHandler) CreateUser(ctx echo.Context) error {
	var input struct{
			Name string `json:"name" validate:"required"`
	}
	if err := ctx.Bind(&input); err != nil{
		return ctx.JSON(http.StatusBadRequest, err)
	}
	user := entity.NewUser(input.Name) 
	err := h.userService.Create(ctx.Request().Context(), user)

	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}

	return ctx.JSON(http.StatusCreated, "User created successfuly")
}