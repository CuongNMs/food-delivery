package tokenprovider

import (
	"errors"
	"food-delivery/common"
	"time"
)

type Provider interface {
	Generate(data TokenPayload, expiry int) (*Token, error)
	Validate(token string) (*TokenPayload, error)
}

var (
	ErrNotFound = common.NewCustomError(
		errors.New("Token not found"),
		"Token not found",
		"ErrNotFound",
	)
	ErrEncodingToken = common.NewCustomError(
		errors.New("Error encoding the token"),
		"Error encoding the token",
		"ErrEncoding",
	)
	ErrInvalidToken = common.NewCustomError(
		errors.New("Invalid token"),
		"Invalid token",
		"ErrTokenInvalid",
	)
)

type Token struct {
	Token   string    `json:"token"`
	Created time.Time `json:"created"`
	Expiry  int       `json:"expiry"`
}

type TokenPayload struct {
	UserId int    `json:"user_id"`
	Role   string `json:"role"`
}
