package users

import (
	"Dema-backend/engine"
	"Dema-backend/graph/model"
	"Dema-backend/models"
	"Dema-backend/utils"
	"errors"
)

var (
	ErrWrongCreds  = errors.New("invalid email or password")
	ErrWrongToken  = errors.New("invalid token")
)

func ValidateRegistration(input model.RegisterUserInput) (*model.AuthPayload, error) {
	if !(input.Password == input.ConfirmPassword) {
		return nil, errors.New("passwords do not match")
	}

	_, err := engine.FetchUserByEmail(input.Email)
	if err == nil {
		return nil, errors.New("email already registered")
	}

	password, err := utils.HashString(input.Password)
	if err != nil {
		return nil, errors.New("internal encryption error")
	}

	newUser := models.User{
		FullName: input.Name,
		Email:    input.Email,
		Password: password,
		Role:     models.Role(input.Role),
	}

	err = utils.DB.Create(&newUser).Error
	if err != nil {
		return nil, err
	}

	return engine.GenerateToken(newUser.ID)
}

func Login(input model.LoginInput) (*model.AuthPayload, error) {
	user, err := engine.FetchUserByEmail(input.Email)
	if err != nil {
		return nil, engine.ErrWrongCreds
	}

	if !utils.CompareHashedString(user.Password, input.Password) {
		return nil, engine.ErrWrongCreds
	}
	return engine.GenerateToken(user.ID)
}

func FetchUser(token string) (*model.User, error){
	userId, err := utils.ValidateJWTForAuthId(token)
	if err != nil {
		return nil, ErrWrongToken
	}

	user, err := engine.FetchUserByID(userId)
	if err != nil {
		return nil, err
	}
	return user, nil
}

