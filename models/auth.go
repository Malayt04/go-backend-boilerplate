package models

import (
	"context"
	"net/mail"

	"golang.org/x/crypto/bcrypt"
)

type AuthCredentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthRepository interface {
	RegisterUser(ctx context.Context, registeredData *AuthCredentials)(*User, error)
	GetUser(ctx context.Context, query interface{}, arge interface{})(*User, error)
}

type AuthService interface{
	Login(ctx context.Context, loginCredentials *AuthCredentials)(*User, error)
	Register(ctx context.Context, registerCredentials *AuthCredentials)(*User, error)
}

func MatchHash(password string, hash string)bool{
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func ValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}