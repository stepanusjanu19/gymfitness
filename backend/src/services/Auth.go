package services

import (
	"backend/src/config"
	"backend/src/model"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func SignUp(user *model.User) (string, error)  {
	var existingUser model.User

	if err := config.DB.Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
		return "", fmt.Errorf("Email its already in use")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	user.Password = string(hashedPassword)

	if err := config.DB.Create(user).Error; err != nil {
		return "", err
	}

	return user.FullName, nil
}

func Login(email, password string) (*model.User, error)  {
	var user model.User

	if err := config.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, fmt.Errorf("User not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, fmt.Errorf("Invalid password")
	}

	return &user, nil
}