package handler

import (
	"apiGo/account/dto"
	"apiGo/account/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AccountHandler struct {
	Service service.AccountService
}

func NewAccountHandler(s service.AccountService) *AccountHandler {
	return &AccountHandler{Service: s}
}

// @Summary Create account
// @Tags accounts
// @Accept json
// @Produce json
// @Param account body dto.CreateAccountDTO true "Account to create"
// @Success 201 {object} dto.AccountResponseDTO
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/accounts [post]
func (h *AccountHandler) Create(c *gin.Context) {
	var input dto.CreateAccountDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	account, err := h.Service.Create(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, account)
}

// @Summary Get all accounts
// @Tags accounts
// @Produce json
// @Success 200 {array} dto.AccountResponseDTO
// @Failure 500 {object} map[string]string
// @Router /api/accounts [get]
func (h *AccountHandler) GetAll(c *gin.Context) {
	accounts, err := h.Service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, accounts)
}

// @Summary Get account by ID (includes user information)
// @Tags accounts
// @Produce json
// @Param id path string true "Account ID"
// @Success 200 {object} dto.AccountResponseDTO
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /api/accounts/{id} [get]
func (h *AccountHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	account, err := h.Service.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Account not found"})
		return
	}
	c.JSON(http.StatusOK, account)
}

// @Summary Get accounts by user ID
// @Tags accounts
// @Produce json
// @Param userId path string true "User ID"
// @Success 200 {array} dto.AccountResponseDTO
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/accounts/user/{userId} [get]
func (h *AccountHandler) GetByUserID(c *gin.Context) {
	userID := c.Param("userId")
	accounts, err := h.Service.GetByUserID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, accounts)
}

// @Summary Update account
// @Tags accounts
// @Accept json
// @Produce json
// @Param id path string true "Account ID"
// @Param account body dto.UpdateAccountDTO true "Account to update"
// @Success 200 {object} dto.AccountResponseDTO
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/accounts/{id} [put]
func (h *AccountHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var input dto.UpdateAccountDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	account, err := h.Service.Update(id, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, account)
}

// @Summary Delete account
// @Tags accounts
// @Produce json
// @Param id path string true "Account ID"
// @Success 204
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/accounts/{id} [delete]
func (h *AccountHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	if err := h.Service.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
