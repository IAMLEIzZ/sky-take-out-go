package request

import "github.com/golang-jwt/jwt/v5"

var JwtAdminSecretKey = []byte("itcast") // 确保密钥为 []byte 类型


type JwtClaimDTO_Admin struct {
	EmpId uint64	`json:"empId"`
	jwt.RegisteredClaims
}