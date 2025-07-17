package service

import (
	"github.com/sahilrana7582/multi-tenant-hotel/auth-service/repo"
	"github.com/sahilrana7582/multi-tenant-hotel/pkg/auth"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Login(email, password string) (string, error)
}

type authService struct {
	repo repo.AuthRepo
}

func NewAuthService(repo repo.AuthRepo) AuthService {
	return &authService{repo: repo}
}

func (s *authService) Login(email, password string) (string, error) {
	dbAuthResp, err := s.repo.Login(email)
	if err != nil {
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(dbAuthResp.Password), []byte(password)); err != nil {
		return "", err
	}

	token, err := auth.GenerateJWT(dbAuthResp.UserId, dbAuthResp.TenantId)
	if err != nil {
		return "", err
	}

	return token, nil
}
