package config

import (
	"github.com/golang-jwt/jwt/v5"
)

var JWT_KEY = []byte("afsjkyiodsufdfgnkdfleoeou12434738375200")

type JWTClaim struct {
	Username string
	jwt.RegisteredClaims
}
