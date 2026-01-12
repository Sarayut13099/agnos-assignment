package implement

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func (s *StaffService) Login(ctx context.Context, username, password, hospital_code string) (string, error) {
	staff, err := s.repo.GetByUsername(ctx, username)
	if err != nil {
		return "", errors.New("user not found")
	}

	if !s.encrypt.CheckPasswordHash(password, strings.TrimSpace(staff.PasswordHash)) {
		return "", errors.New("wrong password")
	}

	claims := jwt.MapClaims{
		"hospital_code": staff.HCode,
		"exp":           time.Now().Add(time.Hour * time.Duration(s.jwtTTL)).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString(s.jwtSecret)

	return t, err
}
