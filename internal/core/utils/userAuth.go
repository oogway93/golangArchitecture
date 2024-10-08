package utils

import (
	"errors"
	"log/slog"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type tokenClaims struct {
	jwt.StandardClaims
	Login string `json:"login"`
	Id    uint   `json:"id"`
}

func HashPassword(password string) string {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		slog.Warn("Error in HASHING password", "error", err)
	}
	return string(hashPassword)
}

func CheckHashPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		slog.Warn("Error in UNHASHING password", "error", err.Error)
		return false
	}
	return true
}

func ParseToken(accessToken string) (string, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(os.Getenv("SECRET")), nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return "", errors.New("token claims are not of type *tokenClaims")
	}

	return claims.Login, nil
}

// FIXME: сделать тут так, чтобы можно было получать из jwt токена id юзера, ибо почему то только логин отображается в data.
func GetUserID(c *gin.Context) (string, error) {
	id, ok := c.Get("currentUserLogin")
	if !ok {
		return "", errors.New("cannot get user's id")
	}

	idInt, ok := id.(string)
	if !ok {
		return "", errors.New("id is of invalid type")
	}

	return idInt, nil
}
