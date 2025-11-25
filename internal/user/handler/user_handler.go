package handler

import (
	"apiGo/internal/user/dto"
	service2 "apiGo/internal/user/service"
	errors2 "apiGo/pkg/errors"
	"apiGo/users/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	Service service2.UserService
}

func NewUserHandler(s service2.UserService) *UserHandler {
	return &UserHandler{Service: s}
}

// @Summary Create user
// @Tags user
// @Accept json
// @Produce json
// @Param user body dto.CreateUserDTO true "User to create"
// @Success 201 {object} dto.UserDTO
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/user [post]
func (h *UserHandler) Create(c *gin.Context) {
	var input dto.CreateUserDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		er := errors2.NewApiError(
			errors2.CodeUserBadRequest,
			errors2.StatusBadRequest,
			errors2.MsgUserBadRequest,
			[]string{err.Error()},
			http.StatusBadRequest,
		)
		errors2.WriteError(c.Writer, er.(*errors2.ApiError))
		return
	}
	user, err := h.Service.Create(input)
	if err != nil {
		if err == service.ErrEmailAlreadyExists {
			er := errors2.NewApiError(
				errors2.CodeUserEmailExists,
				errors2.StatusConflict,
				errors2.MsgUserEmailExists,
				[]string{"El email ya está en uso: " + input.Email},
				http.StatusConflict,
			)
			errors2.WriteError(c.Writer, er.(*errors2.ApiError))
			return
		}
		if apiErr, ok := err.(*errors2.ApiError); ok {
			errors2.WriteError(c.Writer, apiErr)
			return
		}
		er := errors2.NewApiError(
			errors2.CodeUserCreateError,
			errors2.StatusInternalServerError,
			errors2.MsgUserCreateError,
			[]string{err.Error()},
			http.StatusInternalServerError,
		)
		errors2.WriteError(c.Writer, er.(*errors2.ApiError))
		return
	}
	c.JSON(http.StatusCreated, user)
}

// @Summary Get all user
// @Tags user
// @Produce json
// @Success 200 {array} dto.UserDTO
// @Failure 500 {object} map[string]string
// @Router /api/user [get]
func (h *UserHandler) GetAll(c *gin.Context) {
	users, err := h.Service.GetAll()
	if err != nil {
		er := errors2.NewApiError(
			errors2.CodeUserCreateError,
			errors2.StatusInternalServerError,
			errors2.MsgUserGetError,
			[]string{err.Error()},
			http.StatusInternalServerError,
		)
		errors2.WriteError(c.Writer, er)
		return
	}
	c.JSON(http.StatusOK, users)
}

// @Summary Get user by ID
// @Tags user
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} dto.UserDTO
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /api/user/{id} [get]
func (h *UserHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	user, err := h.Service.GetByID(id)
	if err != nil {
		er := errors2.NewApiError(
			errors2.CodeUserNotFound,
			errors2.StatusNotFound,
			errors2.MsgUserNotFound,
			[]string{err.Error()},
			http.StatusNotFound,
		)
		errors2.WriteError(c.Writer, er)
		return
	}
	c.JSON(http.StatusOK, user)
}

// @Summary Update user
// @Tags user
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param user body dto.UpdateUserDTO true "User to update"
// @Success 200 {object} dto.UserDTO
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/user/{id} [put]
func (h *UserHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var input dto.UpdateUserDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		er := errors2.NewApiError(
			errors2.CodeUserBadRequest,
			errors2.StatusBadRequest,
			errors2.MsgUserBadRequest,
			[]string{err.Error()},
			http.StatusBadRequest,
		)
		errors2.WriteError(c.Writer, er)
		return
	}
	user, err := h.Service.Update(id, input)
	if err != nil {
		er := errors2.NewApiError(
			errors2.CodeUserCreateError,
			errors2.StatusInternalServerError,
			errors2.MsgUserUpdateError,
			[]string{err.Error()},
			http.StatusInternalServerError,
		)
		errors2.WriteError(c.Writer, er)
		return
	}
	c.JSON(http.StatusOK, user)
}

// @Summary Delete user
// @Tags user
// @Produce json
// @Param id path string true "User ID"
// @Success 204
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/user/{id} [delete]
func (h *UserHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	if err := h.Service.Delete(id); err != nil {
		er := errors2.NewApiError(
			errors2.CodeUserCreateError,
			errors2.StatusInternalServerError,
			errors2.MsgUserDeleteError,
			[]string{err.Error()},
			http.StatusInternalServerError,
		)
		errors2.WriteError(c.Writer, er)
		return
	}
	c.Status(http.StatusNoContent)
}
