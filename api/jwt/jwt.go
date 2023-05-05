package jwt

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTClaim struct {
	Address string `json:"address"`
	jwt.RegisteredClaims
}

const JWT_EXPIRATION = 24 * 7 * time.Hour

func GenerateJWT(address string) (token string, err error) {

	var claims = JWTClaim{
		address,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(JWT_EXPIRATION)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	resToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := os.Getenv("JWT_SECRET")
	signedToken, err := resToken.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
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
