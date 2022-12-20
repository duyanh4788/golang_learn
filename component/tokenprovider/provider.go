package tokenprovider

import (
	"errors"
	"golang_01/common"
	"time"
)

type Provider interface {
	Generate(data TokenPayload, expiry int) (*Token, error)
	Validate(token string) (*TokenPayload, error)
}

var (
	ErrNotFound = common.NewCustomError(
		errors.New("token not found"),
		"token not found",
		"ErrNotFound",
	)

	ErrEncodingToken = common.NewCustomError(
		errors.New("err encoding token"),
		"err encoding token",
		"ErrEncodingToken",
	)

	ErrInvalidToken = common.NewCustomError(
		errors.New("err invalid token"),
		"err invalid token",
		"ErrInvalidToken",
	)
)

type TokenPayload struct {
	UserId int    `json:"userId"`
	Role   string `json:"role"`
}

type Token struct {
	Token   string    `json:"token"`
	Created time.Time `json:"created"`
	Expiry  int       `json:"expiry"`
}
