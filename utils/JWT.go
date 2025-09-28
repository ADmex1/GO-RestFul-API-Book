package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func JWT(id int, username string, role string, email string, device string) (string, error) {
	secret := []byte("Sekret Token")
	claims := jwt.MapClaims{
		"id":       id,
		"username": username,
		"role":     role,
		"email":    email,
		"exp":      time.Now().Add(time.Hour).Unix(),
		"iat":      time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(secret)

	if err != nil {
		return " ", err
	}
	// fmt.Println("Generate JWT Token: ")
	return signedToken, nil
}
