package user

import "github.com/golang-jwt/jwt/v5"

type UserClaims struct {
	Id       uint
	Username string
	jwt.RegisteredClaims
}
