package repository

import (
	"expenseTracker/internal/account/model"
	userModel "expenseTracker/internal/user/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AccountRepository interface {
	Create(account *model.Account) error
	FindAll() ([]model.Account, []userModel.User, error)
	FindByID(id uuid.UUID) (*model.Account, *userModel.User, error)
	FindByUserID(userID uuid.UUID) ([]model.Account, *userModel.User, error)
	Update(account *model.Account) error
	Delete(id uuid.UUID) error
}

type accountRepository struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) AccountRepository {
	return &accountRepository{db}
}

func (r *accountRepository) Create(account *model.Account) error {
	return r.db.Create(account).Error
}

func (r *accountRepository) FindAll() ([]model.Account, []userModel.User, error) {
	var accounts []model.Account
	var users []userModel.User

	// Obtener todas las cuentas
	err := r.db.Find(&accounts).Error
	if err != nil {
		return nil, nil, err
	}

	// Obtener todos los usuarios únicos de las cuentas
	var userIDs []uuid.UUID
	userIDMap := make(map[uuid.UUID]bool)

	for _, account := range accounts {
		if !userIDMap[account.UserID] {
			userIDs = append(userIDs, account.UserID)
			userIDMap[account.UserID] = true
		}
	}

	if len(userIDs) > 0 {
		err = r.db.Where("id IN ?", userIDs).Find(&users).Error
		if err != nil {
			return nil, nil, err
		}
	}

	return accounts, users, nil
}

func (r *accountRepository) FindByID(id uuid.UUID) (*model.Account, *userModel.User, error) {
	var account model.Account
	var user userModel.User

	// Buscar la cuenta
	err := r.db.First(&account, "id = ?", id).Error
	if err != nil {
		return nil, nil, err
	}

	// Buscar el usuario asociado
	err = r.db.First(&user, "id = ?", account.UserID).Error
	if err != nil {
		return nil, nil, err
	}

	return &account, &user, nil
}

func (r *accountRepository) FindByUserID(userID uuid.UUID) ([]model.Account, *userModel.User, error) {
	var accounts []model.Account
	var user userModel.User

	// Obtener las cuentas del usuario
	err := r.db.Where("user_id = ?", userID).Find(&accounts).Error
	if err != nil {
		return nil, nil, err
	}

	// Obtener el usuario
	err = r.db.First(&user, "id = ?", userID).Error
	if err != nil {
		return nil, nil, err
	}

	return accounts, &user, nil
}

func (r *accountRepository) Update(account *model.Account) error {
	return r.db.Save(account).Error
}

func (r *accountRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&model.Account{}, "id = ?", id).Error
}
