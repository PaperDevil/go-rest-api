package service

import (
	"GoRestAPI/pkg/dao"
	"GoRestAPI/pkg/model"
	"crypto/sha1"
	"fmt"
)

const SECRET = "SOME_SECRE_HERE"

type AuthService struct {
	dao dao.Authorization
}

func NewAuthService(dao dao.Authorization) *AuthService {
	return &AuthService{dao: dao}
}

func (s *AuthService) CreateUser(user model.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.dao.CreateUser(user)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(SECRET)))
}
