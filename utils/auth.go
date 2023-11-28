package utils

import (
	"os"

	"github.com/kerimcanbalkan/Portfolio-API/config"
	"github.com/kerimcanbalkan/Portfolio-API/models"
	"golang.org/x/crypto/bcrypt"
)

func LoginCheck(username string, password string) (string, error) {
	var err error
	u := models.Admin{}

	err = config.DB.Model(models.Admin{}).Where("username = ?", username).Take(&u).Error
	print(os.Getenv("TOKEN_LIFESPAN"))
	if err != nil {
		return "Wrong Username", err
	}

	err = VerifyPassword(password, u.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "Wrong Password", err
	}

	token, err := GenerateToken(u.ID)
	if err != nil {
		return "Authentication Failed", err
	}

	return token, nil
}

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
