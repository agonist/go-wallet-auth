package auth

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTClaim struct {
	Address string `json:"username"`
	jwt.RegisteredClaims
}

func GenerateJWT(address string) (token string, err error) {

	var claims = JWTClaim{
		address,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * 7 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	resToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := os.Getenv("JWT_SECRET")
	fmt.Println(resToken)
	fmt.Println(secret)
	signedToken, err := resToken.SignedString([]byte(secret))
	fmt.Println(signedToken)
	return signedToken, nil
}

func ValidateToken(signedToken string) (addres string, err error) {
	token, err := jwt.ParseWithClaims(signedToken, &JWTClaim{}, func(t *jwt.Token) (interface{}, error) { return []byte(os.Getenv("JWT_SECRET")), nil })
	if err != nil {
		return "", err
	}
	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		return "", errors.New("error parsing claims")
	}
	if claims.ExpiresAt.Unix() < time.Now().Local().Unix() {
		return "", errors.New("token expired")
	}

	return claims.Address, nil
}
