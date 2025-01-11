package helpers

import (
	"backend/src/model"
	"backend/lib/utils/formatstring"
	"crypto/rand"
	"math/big"
	"time"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecretKey = []byte("rKDBhsW3PCOgy3Ui")

func JWTGenerator(user model.User) (string, error)  {
	claims := jwt.MapClaims{
		"user_id": user.UserId,
		"role": user.Role,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(jwtSecretKey)
	if err != nil{
		return "", err
	}
	return signedToken, nil
}

func JWTValidate(token string) (jwt.MapClaims, error)  {
	
	parseToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, formatstring.FormatStringError("ErrInvalidSigningMethod")
		}
		return jwtSecretKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, formatstring.FormatStringError("ErrInvalidTokenSign")
		} else if err.Error() == "Token is expired" {
			return nil, formatstring.FormatStringError("ErrTokenExpired")
		}
		return nil, err
	}

	if claims, ok := parseToken.Claims.(jwt.MapClaims); ok && parseToken.Valid {
		return claims, nil
	}

	return nil, formatstring.FormatStringError("ErrInvalidToken")
}

func HashPassword(password string) (string, error){
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword), err
}

func VerifyPassword(password, hashPassword string) bool  {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
	return err == nil
}

func OtpGenerate() (int, error) {
	otp, err := rand.Int(rand.Reader, big.NewInt(900000))
	if err != nil {
		return 0, err
	}
	return int(otp.Int64()) + 100000, nil
}

