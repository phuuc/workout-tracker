package models

import (
	"regexp"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/finnpn/workout-tracker/pkg/log"
)

const (
	emailRegexp = `/^\S+@\S+\.\S+$/`
)

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

type JWTOut struct {
	Token   string    `json:"token"`
	Expired time.Time `json:"expired"`
}

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Register struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterResponse struct {
	Message string `json:"message"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func EmailRegexp(email string) error {
	_, err := regexp.MatchString(emailRegexp, email)
	if err != nil {
		log.Error("email doesnt match regexp ... with err=%v", err)
		return err
	}
	return nil
}

func PasswordRegexp(pw string) bool {
	return len(pw) > 7
}
