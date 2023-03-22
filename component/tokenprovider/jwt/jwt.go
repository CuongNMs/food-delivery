package jwt

import (
	"food-delivery/component/tokenprovider"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type jwtProvider struct {
	secret string
}

func NewJwtTokenProvider(secret string) *jwtProvider {
	return &jwtProvider{secret: secret}
}

type myClaims struct {
	Payload tokenprovider.TokenPayload `json:"payload"`
	jwt.StandardClaims
}

func (j *jwtProvider) Generate(data tokenprovider.TokenPayload, expiry int) (*tokenprovider.Token, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, myClaims{
		data,
		jwt.StandardClaims{
			ExpiresAt: time.Now(),
		}
	})
}
