package repo

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sahilrana7582/multi-tenant-hotel/auth-service/models"
)

type AuthRepo interface {
	Login(email string) (models.DBAuthResp, error)
	
}

type authRepo struct {
	db *pgxpool.Pool
}

func NewAuthRepo(db *pgxpool.Pool) AuthRepo {
	return &authRepo{db: db}
}

func (r *authRepo) Login(email string) (models.DBAuthResp, error) {
	var dbAuthResp models.DBAuthResp
	query := `
		SELECT tenant_id, id, password_hash
		FROM users
		WHERE email = $1
	`
	err := r.db.QueryRow(context.Background(), query, email).Scan(&dbAuthResp.TenantId, &dbAuthResp.UserId, &dbAuthResp.Password)
	if err != nil {
		return struct {
			TenantId string
			UserId   string
			Password string
		}{}, err
	}

	return dbAuthResp, nil
}
