package helpers

import (
	"backend/src/model"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecretKey = []byte("rKDBhsW3PCOgy3Ui")

var ErrInvalidToken = errors.New("invalid token")
var ErrInvalidSigningMethod = errors.New("invalid signing method")

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
			return nil, ErrInvalidSigningMethod
		}
		return jwtSecretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := parseToken.Claims.(jwt.MapClaims); ok && parseToken.Valid {
		return claims, nil
	}

	return nil, ErrInvalidToken
}