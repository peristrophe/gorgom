//go:generate mockgen -source=$GOFILE -package=mock -destination=../mock/${GOPACKAGE}_${GOFILE}
package repository

import (
	"gorgom/internal/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	CreateUser(string, string) (*entity.User, error)
	Authentication(string, string) (*string, error)
	BoardByID(uuid.UUID) *entity.Board
	BoardsByGroupID(uuid.UUID) []*entity.Board
}

type repository struct {
	DB *gorm.DB
}

func NewRepository() *repository {
	db := GetDBConn()
	repo := repository{DB: db}
	return &repo
}

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

	//result = tx.Update("Password", &user)
	result = tx.Model(&user).Updates(entity.User{Password: user.Password})
	if result.Error != nil {
		tx.Rollback()
		return nil, result.Error
	}

	tx.Commit()
	return &user, nil
}

func (r *repository) Authentication(email string, password string) (*string, error) {
	var user entity.User
	result := r.DB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	err := user.Authentication(password)
	if err != nil {
		return nil, err
	}
	token := "TOKEN"
	return &token, nil
}

func (r *repository) BoardByID(boardId uuid.UUID) *entity.Board {
	var board entity.Board
	r.DB.
		Preload("Boxes.Cards.Comments").
		Preload("Boxes.Cards.Tags").
		Take(&board, boardId)

	return &board
}

func (r *repository) BoardsByGroupID(groupId uuid.UUID) []*entity.Board {
	return nil
}
