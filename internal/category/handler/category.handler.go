package handler

import (
	"expenseTracker/internal/category/dto"
	"expenseTracker/internal/category/service"
	"expenseTracker/pkg/errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CategoryHandlder struct {
	Service service.CategoryService
}

func NewCategoryHandlder(s service.CategoryService) *CategoryHandlder {
	return &CategoryHandlder{Service: s}
}

func (h *CategoryHandlder) Create(c *gin.Context) {
	var input dto.CreateCategoryDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		er := errors.NewApiError(
			errors.CodeCategoryBadRequest,
			errors.StatusBadRequest,
			errors.MsgCategoryBadRequest,
			[]string{err.Error()},
			http.StatusBadRequest,
		)
		errors.WriteError(c.Writer, er.(*errors.ApiError))
		return
	}

	category, err := h.Service.Create(input)
	if err != nil {
		if userErr, ok := err.(*errors.UserNotFoundError); ok {
			er := errors.NewApiError(
				errors.CodeCategoryUserNotFound,
				errors.StatusNotFound,
				errors.MsgCategoryUserNotFound,
				[]string{userErr.Error()},
				http.StatusNotFound,
			)
			errors.WriteError(c.Writer, er.(*errors.ApiError))
			return
		}
		er := errors.NewApiError(
			errors.CodeCategoryCreateError,
			errors.StatusInternalServerError,
			errors.MsgCategoryCreateError,
			[]string{err.Error()},
			http.StatusInternalServerError,
		)
		errors.WriteError(c.Writer, er.(*errors.ApiError))
		return
	}

	c.JSON(http.StatusCreated, category)
}

func (h *CategoryHandlder) GetAll(c *gin.Context) {
	categories, err := h.Service.GetAll()
	if err != nil {
		er := errors.NewApiError(
			errors.CodeCategoryCreateError,
			errors.StatusInternalServerError,
			errors.MsgCategoryCreateError,
			[]string{err.Error()},
			http.StatusInternalServerError,
		)
		errors.WriteError(c.Writer, er.(*errors.ApiError))
		return
	}
	c.JSON(http.StatusOK, categories)
}

func (h *CategoryHandlder) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		er := errors.NewApiError(
			errors.CodeCategoryBadRequest,
			errors.StatusBadRequest,
			errors.MsgCategoryBadRequest,
			[]string{"ID inválido"},
			http.StatusBadRequest,
		)
		errors.WriteError(c.Writer, er.(*errors.ApiError))
		return
	}
	category, err := h.Service.GetByID(id)
	if err != nil {
		er := errors.NewApiError(
			errors.CodeCategoryUserNotFound,
			errors.StatusNotFound,
			errors.MsgCategoryUserNotFound,
			[]string{err.Error()},
			http.StatusNotFound,
		)
		errors.WriteError(c.Writer, er.(*errors.ApiError))
		return
	}
	c.JSON(http.StatusOK, category)
}

func (h *CategoryHandlder) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		er := errors.NewApiError(
			errors.CodeCategoryBadRequest,
			errors.StatusBadRequest,
			errors.MsgCategoryBadRequest,
			[]string{"ID inválido"},
			http.StatusBadRequest,
		)
		errors.WriteError(c.Writer, er.(*errors.ApiError))
		return
	}
	var input dto.UpdateCategoryDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		er := errors.NewApiError(
			errors.CodeCategoryBadRequest,
			errors.StatusBadRequest,
			errors.MsgCategoryBadRequest,
			[]string{err.Error()},
			http.StatusBadRequest,
		)
		errors.WriteError(c.Writer, er.(*errors.ApiError))
		return
	}
	category, err := h.Service.Update(id, input)
	if err != nil {
		er := errors.NewApiError(
			errors.CodeCategoryCreateError,
			errors.StatusInternalServerError,
			errors.MsgCategoryCreateError,
			[]string{err.Error()},
			http.StatusInternalServerError,
		)
		errors.WriteError(c.Writer, er.(*errors.ApiError))
		return
	}
	c.JSON(http.StatusOK, category)
}

func (h *CategoryHandlder) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		er := errors.NewApiError(
			errors.CodeCategoryBadRequest,
			errors.StatusBadRequest,
			errors.MsgCategoryBadRequest,
			[]string{"ID inválido"},
			http.StatusBadRequest,
		)
		errors.WriteError(c.Writer, er.(*errors.ApiError))
		return
	}
	if err := h.Service.Delete(id); err != nil {
		er := errors.NewApiError(
			errors.CodeCategoryCreateError,
			errors.StatusInternalServerError,
			errors.MsgCategoryCreateError,
			[]string{err.Error()},
			http.StatusInternalServerError,
		)
		errors.WriteError(c.Writer, er.(*errors.ApiError))
		return
	}
	c.Status(http.StatusNoContent)
}

func (h *CategoryHandlder) CategoryByUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		er := errors.NewApiError(
			errors.CodeCategoryBadRequest,
			errors.StatusBadRequest,
			errors.MsgCategoryBadRequest,
			[]string{"ID inválido"},
			http.StatusBadRequest,
		)
		errors.WriteError(c.Writer, er.(*errors.ApiError))
		return
	}
	categories, err := h.Service.CategoryByUser(id)
	if err != nil {
		er := errors.NewApiError(
			errors.CodeCategoryUserNotFound,
			errors.StatusNotFound,
			errors.MsgCategoryUserNotFound,
			[]string{err.Error()},
			http.StatusNotFound,
		)
		errors.WriteError(c.Writer, er.(*errors.ApiError))
		return
	}
	c.JSON(http.StatusOK, categories)
}
