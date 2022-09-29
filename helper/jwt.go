package helper

import (
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
)

type TokenRequest struct {
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
}

type UserTokenClaims struct {
	jwt.StandardClaims
	User interface{} `json:"user"`
}

func GenerateToken(req TokenRequest) (jwtToken string, err error) {
	sesExp, _ := strconv.Atoi(os.Getenv("SESSION_EXP"))
	expirationTime := time.Now().Add(time.Duration(sesExp) * time.Minute)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": req,
		"nbf": time.Now().Unix(),
		"exp": expirationTime.Unix(),
	})

	tokenStr, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}
