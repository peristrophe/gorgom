//go:generate mockgen -source=$GOFILE -package=mock -destination=../mock/${GOPACKAGE}_${GOFILE}
package repository

import (
	"gorgom/internal/entity"

	"github.com/google/uuid"
)

func (r *repository) CreateUser(email string, password string) (*entity.User, error) {
	tx := r.DB.Begin()

	var initialRole entity.Role
	result := tx.First(&initialRole)
	if result.Error != nil {
		tx.Rollback()
		return nil, result.Error
	}

	user := entity.User{Email: email, Role: initialRole}
	result = tx.Create(&user)
	if result.Error != nil {
		tx.Rollback()
		return nil, result.Error
	}

	err := user.SetPassword(password)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	result = tx.Model(&user).Updates(entity.User{Password: user.Password})
	if result.Error != nil {
		tx.Rollback()
		return nil, result.Error
	}

	tx.Commit()
	return &user, nil
}

func (r *repository) GetUserByID(userID uuid.UUID) (*entity.User, error) {
	var user entity.User
	result := r.DB.
		Preload("Groups").
		Preload("Role").
		Take(&user, userID)

	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (r *repository) GetUserByEmail(email string) (*entity.User, error) {
	var user entity.User
	result := r.DB.Where("email = ?", email).First(&user)

	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
