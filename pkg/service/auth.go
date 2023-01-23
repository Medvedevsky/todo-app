package service

import (
	"crypto/sha256"
	"fmt"
	"time"

	"github.com/Medvedevsky/todo-app"
	"github.com/Medvedevsky/todo-app/pkg/repository"
	"github.com/golang-jwt/jwt"
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

type AutharizationService struct {
	rep repository.Autharization
}

func NewAutharizationService(rep repository.Autharization) *AutharizationService {
	return &AutharizationService{
		rep: rep,
	}
}

func (s AutharizationService) CreateUser(user todo.User) (int, error) {
	user.Password = GetHash(user.Password)
	return s.rep.CreateUser(user)
}

func (s AutharizationService) GenerateToken(login, password string) (string, error) {
	tokenTTL := 12 * time.Hour
	signingKey := "njkbnvlkg$1#dasmnvbmnz106"

	user, err := s.rep.GetUser(login, GetHash(password))

	if err != nil {
		return "", nil
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})

	return token.SignedString([]byte(signingKey))
}

func (s AutharizationService) ParseToken(token string) (int, error) {
	return 0, nil
}

func GetHash(password string) string {
	h := sha256.New()
	h.Write([]byte(password))
	return fmt.Sprintf("%x", h.Sum(nil))
}
