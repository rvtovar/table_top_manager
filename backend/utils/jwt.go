package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"os"
	"time"
)

func GenerateToken(email string, userId int64) (string, error) {
	salt, err := OpenSalt()
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	})
	return token.SignedString([]byte(salt))
}

func ValidateToken(token string) (int64, error) {
	salt, err := OpenSalt()
	if err != nil {
		return 0, err
	}
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("Unexpected signing method")
		}
		return []byte(salt), nil
	})

	if err != nil {
		return 0, err
	}

	tokenIsValid := parsedToken.Valid
	if !tokenIsValid {
		return 0, errors.New("Token is not valid")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("Invalid Token claims")
	}

	//email, _ := claims["email"].(string)
	userId, _ := claims["userId"].(float64)
	return int64(userId), nil

}
func OpenSalt() (string, error) {
	err := godotenv.Load()
	if err != nil {
		return "", errors.New("Error loading .env file")
	}
	salt := os.Getenv("salt")
	if salt == "" {
		return "", errors.New("Salt not found")
	}

	return salt, nil
}
