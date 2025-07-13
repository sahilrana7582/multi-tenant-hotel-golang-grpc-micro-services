package repo

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sahilrana7582/multi-tenant-hotel/user-service/models"
)

type UserRepo interface {
	Create(ctx context.Context, user *models.NewUser) (*models.User, error)
	// GetUserByEmail(ctx context.Context, email string) (*models.User, error)
}

type userRepoImpl struct {
	db *pgxpool.Pool
}

func NewUserRepo(db *pgxpool.Pool) UserRepo {
	return &userRepoImpl{
		db: db,
	}
}

func (r *userRepoImpl) Create(ctx context.Context, user *models.NewUser) (*models.User, error) {

	query := `
		INSERT INTO users (tenant_id, name, email, password_hash)
		VALUES ($1, $2, $3, $4)
		RETURNING id, tenant_id, name, email, created_at, updated_at
	`

	var newUser models.User
	err := r.db.QueryRow(ctx, query, user.TenantID, user.Name, user.Email, user.Password).Scan(
		&newUser.ID,
		&newUser.TenantID,
		&newUser.Name,
		&newUser.Email,
		&newUser.CreatedAt,
		&newUser.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &newUser, nil
}
