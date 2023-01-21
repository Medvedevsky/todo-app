package service

import "github.com/Medvedevsky/todo-app/pkg/repository"

type AutharizationService struct {
	rep repository.Autharization
}

func NewAutharizationService(rep repository.Autharization) *AutharizationService {
	return &AutharizationService{
		rep: rep,
	}
}
