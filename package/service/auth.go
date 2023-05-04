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

func NewAuthService(repo repository.Authorization) *AuthService {
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

type CustomClaims struct {
	UserId uint64 `json:"userId"`
	jwt.RegisteredClaims
}

func (s *AuthService) GenerateToken(username, pass string) (string, error) {
	user, err := s.repo.GetUser(username, s.generatePassword(pass))

	if err != nil {
		return "", err
	}

	claims := CustomClaims{
		UserId: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TOKEN_TTL)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(os.Getenv("TOKEN_SIGNIN_KEY")))
}

func (s *AuthService) ParseToken(accessToken string) (uint64, error) {
	token, err := jwt.ParseWithClaims(accessToken, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("TOKEN_SIGNIN_KEY")), nil // todo check with diff value
	})
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims.UserId, err
	} else {
		return 0, err
	}
}
