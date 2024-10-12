package handler

import (
	"net/http"

	"github.com/Sandhya-Pratama/weather-app/entity"
	"github.com/Sandhya-Pratama/weather-app/internal/http/validator"
	"github.com/Sandhya-Pratama/weather-app/internal/service"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService service.UserUseCase
}

func NewUserHandler(userService service.UserUseCase) *UserHandler {
	return &UserHandler{userService}
}

func (h *UserHandler) GetAllUsers(ctx echo.Context) error {
	users, err := h.userService.FindAll(ctx.Request().Context())

	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}
	return ctx.JSON(http.StatusOK, users)
}

func (h *UserHandler) CreateUser(ctx echo.Context) error {
	var input struct {
		Name     string `json:"name" validate:"required"`
		Email    string `json:"email" validate:"required"`
		Password string `json:"password" validate:"required"`
		Role     string `json:"role" validate:"required" oneof:"Administrator,Editor"`
	}

	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}

	user := entity.NewUser(input.Name, input.Email, input.Password, input.Role)
	err := h.userService.Create(ctx.Request().Context(), user)

	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}

	return ctx.JSON(http.StatusCreated, "User created successfuly")
}

func (h *UserHandler) UpdateUser(ctx echo.Context) error {
	var input struct {
		ID       int64  `param:"id" validate:"required"`
		Name     string `json:"name" validate:"required"`
		Email    string `json:"email" validate:"required"`
		Password string `json:"password" validate:"required"`
		Role     string `json:"role" validate:"required" oneof:"Administrator,Editor"`
	}

	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}
	user := entity.UpdateUser(input.ID, input.Name, input.Email, input.Password, input.Role)
	err := h.userService.Update(ctx.Request().Context(), user)

	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}
	return ctx.JSON(http.StatusOK, user)
}

func (h *UserHandler) DeleteUser(ctx echo.Context) error {
	var input struct {
		ID int64 `param:"id" validate:"required"`
	}
	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}

	err := h.userService.Delete(ctx.Request().Context(), input.ID)

	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}
	return ctx.NoContent(http.StatusNoContent)
}

func (h *UserHandler) GetUserByID(ctx echo.Context) error {
	var input struct {
		ID int64 `param:"id" validate:"required"`
	}
	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}
	user, err := h.userService.FindByID(ctx.Request().Context(), input.ID)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}
	return ctx.JSON(http.StatusOK, user)
}
