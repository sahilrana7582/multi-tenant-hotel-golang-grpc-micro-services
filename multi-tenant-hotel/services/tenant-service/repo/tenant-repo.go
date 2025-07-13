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
	GetByID(ctx context.Context, id string) (*models.Tenant, error)
	List(ctx context.Context) ([]*models.Tenant, error)
	Update(ctx context.Context, t *models.UpdateTenantInput, id string) (*models.Tenant, error)
	Delete(ctx context.Context, id string) error
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

func (r *DbTenantRepo) GetByID(ctx context.Context, id string) (*models.Tenant, error) {

	query := `
		SELECT id, name, email, phone, status, created_at, updated_at
		FROM tenants
		WHERE id = $1
	`

	t := &models.Tenant{}

	err := r.db.QueryRow(ctx, query, id).Scan(&t.ID, &t.Name, &t.Email, &t.Phone, &t.Status, &t.CreatedAt, &t.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return t, nil
}

func (r *DbTenantRepo) List(ctx context.Context) ([]*models.Tenant, error) {
	query := `
		SELECT id, name, email, phone, status, created_at, updated_at
		FROM tenants
	`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	tenants := []*models.Tenant{}

	for rows.Next() {
		t := &models.Tenant{}
		err := rows.Scan(&t.ID, &t.Name, &t.Email, &t.Phone, &t.Status, &t.CreatedAt, &t.UpdatedAt)
		if err != nil {
			return nil, err
		}
		tenants = append(tenants, t)
	}
	return tenants, nil
}

func (r *DbTenantRepo) Update(ctx context.Context, input *models.UpdateTenantInput, id string) (*models.Tenant, error) {

	query := `
		UPDATE tenants
		SET 
			name  = COALESCE(NULLIF($1, ''), name),
			email = COALESCE(NULLIF($2, ''), email),
			phone = COALESCE(NULLIF($3, ''), phone)
		WHERE id = $4
		RETURNING id, name, email, phone, status, created_at, updated_at
	`

	t := &models.Tenant{}

	err := r.db.QueryRow(ctx, query, input.Name, input.Email, input.Phone, id).Scan(&t.ID, &t.Name, &t.Email, &t.Phone, &t.Status, &t.CreatedAt, &t.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return t, nil
}

func (r *DbTenantRepo) Delete(ctx context.Context, id string) error {
	query := `
		DELETE FROM tenants
		WHERE id = $1
	`
	_, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}
