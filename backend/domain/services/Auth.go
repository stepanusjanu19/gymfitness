package services

import (
	"backend/domain/requests"
	"backend/lib/config"
	"backend/lib/helpers"
	"backend/lib/pkg/mails"
	"backend/lib/utils/formatstring"
	"backend/src/model"
	"html"
	"log"
	"os"
	"strings"

	"github.com/jinzhu/gorm"
)

func getDB() *gorm.DB {
	db := config.ConnectDB()
	if db == nil {
		log.Fatalf("Failed to initialize database connection")
		os.Exit(1)
	}
	return db.Debug()
}

func Login(email, password string) (model.User, error) {
	var user model.User
	db := getDB()

	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		return user, formatstring.FormatStringError("userfound")
	}

	if !user.IsActive {
		return user, formatstring.FormatStringError("userinactive")
	}

	if !helpers.VerifyPassword(password, user.Password) {
		return user, formatstring.FormatStringError("hashedPassword")
	}

	return user, nil
}

func SignUp(request *requests.Register) (model.User, error) {
	var existingUser model.User
	db := getDB()

	if err := db.Where("email = ?", request.Email).First(&existingUser).Error; err == nil {
		return existingUser, formatstring.FormatStringError("emailalready")
	}

	hashedPassword, err := helpers.HashPassword(request.Password)
	if err != nil {
		return model.User{}, formatstring.FormatStringError("userfound")
	}

	newUser := model.User{
		FullName: html.EscapeString(strings.TrimSpace(request.FullName)),
		Email:    html.EscapeString(strings.TrimSpace(request.Email)),
		Password: hashedPassword,
		IsActive: true,
		Role:     model.Guest,
	}

	if newUser.FullName != "" {
		parts := strings.Fields(newUser.FullName)
		if len(parts) > 0 {
			newUser.FirstName = parts[0]
			if len(parts) > 1 {
				newUser.LastName = strings.Join(parts[:1], " ")
			}
		}
	}

	transactionData := db.Begin()

	if transactionData.Error != nil {
		return model.User{}, formatstring.FormatStringErrorWithDetails("transactionStartError", transactionData.Error)
	}

	defer func() {
		if r := recover(); r != nil {
			transactionData.Rollback()
			panic(r)
		} else if transactionData.Error != nil {
			transactionData.Rollback()
		}
	}()

	if err := transactionData.Create(&newUser).Error; err != nil {
		transactionData.Rollback()
		return model.User{}, formatstring.FormatStringErrorWithDetails("userCreationError", err)
	}

	otpGenerate, err := helpers.OtpGenerate()
	if err != nil {
		return newUser, formatstring.FormatStringErrorWithDetails("otpGenerationError", err)
	}

	err = mails.SendMailerOTPbyAPI(newUser.Email, newUser.FullName, otpGenerate)
	if err != nil {
		return newUser, formatstring.FormatStringErrorWithDetails("otpEmailError", err)
	}

	verificationUser := model.VerificationUser{
		UserId: newUser.UserId,
		Otp:    otpGenerate,
		IsUsed: false,
	}

	if err := transactionData.Create(&verificationUser).Error; err != nil {
		transactionData.Rollback()
		return newUser, formatstring.FormatStringErrorWithDetails("otpSaveError", err)
	}

	if err := transactionData.Commit().Error; err != nil {
		transactionData.Rollback()
		return model.User{}, formatstring.FormatStringErrorWithDetails("transactionCommitError", err)
	}

	return newUser, nil
}
