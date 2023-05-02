package service

import (
	"crypto/sha1"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	todo "go-task-manager-system"
	"go-task-manager-system/package/repository"
	"os"
	"time"
)

const (
	TOKEN_TTL = 12 * time.Hour
)

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

	return fmt.Sprintf("%x", hash.Sum([]byte(os.Getenv("HASH_SALT"))))
}

func (s *AuthService) GenerateToken(username, pass string) (string, error) {
	user, err := s.repo.GetUser(username, s.generatePassword(pass))

	if err != nil {
		return "", err
	}

	tokenClaims := jwt.MapClaims{
		"exp":    time.Now().Add(TOKEN_TTL).Unix(),
		"iat":    time.Now().Unix(),
		"userId": user.ID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims)

	return token.SignedString([]byte(os.Getenv("TOKEN_SIGNIN_KEY")))
}
