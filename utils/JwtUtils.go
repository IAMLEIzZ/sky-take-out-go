package utils

import (
	"fmt"
	"sky-take-out-go/model/dto"
	"github.com/golang-jwt/jwt/v5"
)

// Defined a Jwtutils

// Create JwtToken
func CreateJwt(claim *dto.JwtClaimDTO_Admin, JwtAdminSecretKey []byte) (string, error){
	// create jwt
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	tokenString, err := token.SignedString(JwtAdminSecretKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// Parse JwtToken
func ParseToken(tokenString string) (*dto.JwtClaimDTO_Admin, error) {
	token, err := jwt.ParseWithClaims(tokenString, &dto.JwtClaimDTO_Admin{}, func(token *jwt.Token) (interface{}, error) {
        return dto.JwtAdminSecretKey, nil
    })
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*dto.JwtClaimDTO_Admin); ok && token.Valid {
		return claims, nil
	} else {
		return nil, fmt.Errorf("invild token")
	}
}