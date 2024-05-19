package engine

import (
	
	"errors"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm/clause"
	"Dema-backend/graph/model"
	"Dema-backend/models"
	"Dema-backend/utils"
)

var (
	ErrWrongCreds = errors.New("invalid email or password")
	ErrWrongToken = errors.New("invalid token")
)

func FetchUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := utils.DB.Where("email =?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
func FetchUserById(id string) (*models.User, error) {
	var user models.User
	err := utils.DB.Where("id =?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
func FetchUserByAuthToken(jwt string) (*models.User, error) {
	userId, err := utils.ValidateJWTForAuthId(jwt)
	if err != nil {
		return nil, ErrWrongToken
	}
	return FetchUserById(userId)
}

func GenerateToken(authid uuid.UUID) (*model.AuthPayload, error) {
	authToken, err := utils.GenerateJWTForAuthId(authid)
	if err != nil {
		return nil, err
	}
	return &model.AuthPayload{
		Token: authToken,
	}, nil
}

func FetchUserByID(id string) (*model.User, error) {
	var user model.User
	err := utils.DB.Preload(clause.Associations).Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}


