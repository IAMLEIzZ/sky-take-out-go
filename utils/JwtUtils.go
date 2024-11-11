package utils

import (
	"fmt"
	"sky-take-out-go/internal/api/request"
	"github.com/golang-jwt/jwt/v5"
)

// Defined a Jwtutils

// Create JwtToken
func CreateJwt(claim *request.JwtClaimDTO_Admin, JwtAdminSecretKey []byte) (string, error) {
	// create jwt
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	tokenString, err := token.SignedString(JwtAdminSecretKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// Parse JwtToken
func ParseToken(tokenString string) (*request.JwtClaimDTO_Admin, error) {
	token, err := jwt.ParseWithClaims(tokenString, &request.JwtClaimDTO_Admin{}, func(token *jwt.Token) (interface{}, error) {
		return request.JwtAdminSecretKey, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*request.JwtClaimDTO_Admin); ok && token.Valid {
		return claims, nil
	} else {
		return nil, fmt.Errorf("invild token")
	}
}
