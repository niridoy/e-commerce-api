package handler

import (
	"net/http"
	"strconv"

	"echo-go-api/models"
	"echo-go-api/usecase"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	usecase usecase.IUserUsecase
}

func NewUserHandler(u usecase.IUserUsecase) *UserHandler {
	return &UserHandler{usecase: u}
}

func (h *UserHandler) GetUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid ID")
	}
	user, err := h.usecase.GetUser(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, "User not found")
	}
	return c.JSON(http.StatusOK, user)
}

func (h *UserHandler) CreateUser(c echo.Context) error {
	var user models.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request body")
	}
	if err := h.usecase.CreateUser(&user); err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to create user"+err.Error())
	}
	return c.JSON(http.StatusCreated, "User created successfully")
}

func (h *UserHandler) GetUsers(c echo.Context) error {
	users, err := h.usecase.GetUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to fetch users")
	}
	return c.JSON(http.StatusOK, users)
}

func (h *UserHandler) UpdateUser(c echo.Context) error {
	id := c.Param("id")
	var user models.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request body")
	}

	if err := h.usecase.UpdateUser(id, &user); err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to update user")
	}

	return c.JSON(http.StatusOK, "User updated successfully")
}

func (h *UserHandler) DeleteUser(c echo.Context) error {
	id := c.Param("id")

	if err := h.usecase.DeleteUser(id); err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to delete user")
	}

	return c.JSON(http.StatusOK, "User deleted successfully")
}
