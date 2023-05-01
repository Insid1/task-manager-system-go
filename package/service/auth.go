package service

import (
	"crypto/sha1"
	"fmt"
	todo "go-task-manager-system"
	"go-task-manager-system/package/repository"
)

const SALT = "thisistestsalttohashpassword"

type AuthService struct {
	repo repository.Authorization
}

func newAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user todo.User) (uint64, error) {
	user.Password = s.generatePassword(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) generatePassword(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(SALT)))
}
