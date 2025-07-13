package repo

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sahilrana7582/hotel-mgmt/services/tenant-service/models"
	"github.com/sahilrana7582/multi-tenant-hotel/pkg/errs"
)

var (
	ErrTenantNotFound  = errors.New("tenent not found")
	ErrDuplicateTenant = errors.New("tenant with this name already exists")
)

type Tenant struct {
	ID        uuid.UUID
	Name      string
	Email     string
	Phone     *string
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type TenantRepo interface {
	Create(ctx context.Context, t *models.CreateTenantInput) (*models.Tenant, error)
	// GetByID(ctx context.Context, id uuid.UUID) (*Tenant, error)
	// List(ctx context.Context) ([]*Tenant, error)
	// Update(ctx context.Context, t *Tenant) (*Tenant, error)
}

type DbTenantRepo struct {
	db *pgxpool.Pool
}

func NewTenantRepo(conn *pgxpool.Pool) TenantRepo {

	if conn == nil {
		panic("❌ pgxpool.Pool is nil: DB connection not initialized")
	}

	return &DbTenantRepo{db: conn}
}

func (r *DbTenantRepo) Create(ctx context.Context, input *models.CreateTenantInput) (*models.Tenant, error) {
	query := `
		INSERT INTO tenants (name, email, phone)
		VALUES ($1, $2, $3)
		RETURNING id, name, email, phone, status, created_at, updated_at
	`

	t := &models.Tenant{}

	err := r.db.QueryRow(ctx, query, input.Name, input.Email, input.Phone).
		Scan(&t.ID, &t.Name, &t.Email, &t.Phone, &t.Status, &t.CreatedAt, &t.UpdatedAt)

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return nil, errs.ErrDuplicateTenant
		}
		return nil, errs.ErrInternal
	}

	return t, nil
}
