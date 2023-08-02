package repository

import (
	"gorgom/internal/entity"
	"sort"

	"github.com/google/uuid"
)

func (r *repository) CreateUser(email string, password string) (*entity.User, error) {
	tx := r.DB.Begin()

	user := entity.User{Email: email}
	result := tx.Create(&user)
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
		Preload("Groups.DefinedRoles").
		Preload("Groups.Members").
		Preload("Roles").
		Take(&user, userID)

	if result.Error != nil {
		return nil, result.Error
	}

	sort.Slice(user.Roles, func(i, j int) bool { return user.Roles[i].GroupID == user.Groups[j].ID })
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
