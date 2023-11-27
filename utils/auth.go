package utils

import (
	"github.com/kerimcanbalkan/Portfolio-API/config"
	"github.com/kerimcanbalkan/Portfolio-API/models"
	token "github.com/kerimcanbalkan/Portfolio-API/utils/token"
	"golang.org/x/crypto/bcrypt"
)

func LoginCheck(username string, password string) (string, error) {
	var err error
	u := models.Admin{}

	err = config.DB.Model(models.Admin{}).Where("username = ?", username).Take(&u).Error

	if err != nil {
		return "", err
	}

	err = VerifyPassword(password, u.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token, err := token.GenerateToken(u.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
