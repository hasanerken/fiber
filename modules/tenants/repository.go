package tenants

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"xhantos/sqlboiler/models"
)

type TenantInterface interface {
	Create(ctx context.Context, tenant *Tenant) (*Tenant, error)
	Update(ctx context.Context, tenant *Tenant) (int, error)
	Delete(ctx context.Context, tenantID int) error
	FindOne(ctx context.Context, tenantID int) (*Tenant, error)
	FindMany(ctx context.Context) ([]*Tenant, error)
}

type TenantRepository struct {
	db *sqlx.DB
}

// NewTenantRepository creates a new instance of TenantRepository.
func NewTenantRepository(db *sqlx.DB) *TenantRepository {
	return &TenantRepository{
		db: db,
	}
}

func (r *TenantRepository) Create(ctx context.Context, t *Tenant) (*Tenant, error) {
	// Convert tenant to *models.Tenant if needed and implement Create
	modelTenant := &models.Tenant{
		Alias:  t.Alias,
		APIKey: t.APIKey,
		Status: models.TenantStatus(t.Status),
	}
	err := modelTenant.Insert(ctx, r.db, boil.Infer())
	if err != nil {
		return nil, err
	}

	createdTenant := mapTenantFromModel(modelTenant)

	return &createdTenant, nil
}

func (r *TenantRepository) Update(ctx context.Context, t *Tenant) (int, error) {
	// Convert tenant to *models.Tenant if needed and implement Update
	// Convert tenant to *models.Tenant if needed and implement Create
	modelTenant := &models.Tenant{
		Alias:  t.Alias,
		APIKey: t.APIKey,
		Status: models.TenantStatus(t.Status),
	}
	tenantId, err := modelTenant.Update(ctx, r.db, boil.Infer())
	if err != nil {
		return 0, err
	}

	return int(tenantId), nil
}

func (r *TenantRepository) Delete(ctx context.Context, tenantID int) error {
	// Implement Delete
	return nil
}

func (r *TenantRepository) FindOne(ctx context.Context, tenantID int) (*Tenant, error) {
	// Implement FindOne
	return nil, nil
}

func (r *TenantRepository) FindMany(ctx context.Context) ([]*Tenant, error) {
	// Implement FindMany
	return nil, nil
}
