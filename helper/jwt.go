package helper

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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

func GetTokenParse(bearToken string) (interface{}, error) {
	strArr := strings.Split(bearToken, " ")
	if len(strArr) != 2 {
		log.Println(strArr)
	}

	token, err := jwt.Parse(strArr[1], func(token *jwt.Token) (interface{}, error) {
		// ! Don't forget to validate the alg what you expect
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	jwtToken := token.Claims.(jwt.MapClaims)
	sub := jwtToken["sub"]

	return sub, err
}
