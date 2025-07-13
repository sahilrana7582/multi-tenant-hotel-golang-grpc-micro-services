package service

import (
	"context"

	"github.com/sahilrana7582/multi-tenant-hotel/user-service/models"
	"github.com/sahilrana7582/multi-tenant-hotel/user-service/repo"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	CreateUser(ctx context.Context, user *models.NewUser) (*models.User, error)
	// GetUser(ctx context.Context, id string) (*models.User, error)
	// UpdateUser(ctx context.Context, user *models.User) error
	// DeleteUser(ctx context.Context, id string) error
}

type userService struct {
	repo repo.UserRepo
}

func NewUserService(repo repo.UserRepo) UserService {
	return &userService{repo: repo}
}

func (s *userService) CreateUser(ctx context.Context, user *models.NewUser) (*models.User, error) {

	hashedPaassword, err := hashPassword(user.Password)
	if err != nil {
		return nil, err
	}

	user.Password = hashedPaassword
	return s.repo.Create(ctx, user)
}

func hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
