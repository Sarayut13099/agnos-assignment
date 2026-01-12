package utils

import (
	"golang.org/x/crypto/bcrypt"
)

type Encrypt interface {
	HashPassword(password string) string
	CheckPasswordHash(password, hash string) bool
}

type encrypt struct{}

func NewEncrypt() (e Encrypt) {
	return &encrypt{}
}

func (e *encrypt) HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return ""
	}
	return string(bytes)
}

func (e *encrypt) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
