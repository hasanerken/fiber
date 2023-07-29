package tenants

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"xhantos/common"
	"xhantos/sqlboiler/models"
)

type TenantRepositoryInterface interface {
	Create(ctx context.Context, tenant *Tenant) (*Tenant, error)
	Update(ctx context.Context, tenant *Tenant) error
	Delete(ctx context.Context, tenantID int) error
	FindOne(ctx context.Context, tenantID int) (*Tenant, error)
	FindMany(ctx context.Context, queryParams common.QueryParams) ([]*Tenant, error)
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

	return createdTenant, nil
}

func (r *TenantRepository) Update(ctx context.Context, t *Tenant) error {
	modelTenant := &models.Tenant{
		ID:     t.ID,
		Alias:  t.Alias,
		APIKey: t.APIKey,
		Status: models.TenantStatus(t.Status),
	}
	_, err := modelTenant.Update(ctx, r.db, boil.Infer())
	if err != nil {
		return err
	}

	return nil
}

func (r *TenantRepository) Delete(ctx context.Context, tenantID int) error {
	modelTenant := &models.Tenant{
		ID: tenantID,
	}
	_, err := modelTenant.Delete(ctx, r.db)
	if err != nil {
		return err
	}
	return nil
}

func (r *TenantRepository) FindOne(ctx context.Context, tenantID int) (*Tenant, error) {
	t, err := models.FindTenant(ctx, r.db, tenantID)
	if err != nil {
		return nil, err
	}

	tenant := mapTenantFromModel(t)
	return tenant, nil
}

func (r *TenantRepository) FindMany(ctx context.Context, queryParams common.QueryParams) ([]*Tenant, error) {
	// Implement FindMany
	modifiers := common.BuildQueryModifiers(queryParams)
	tenants, err := models.Tenants(modifiers...).All(ctx, r.db)
	if err != nil {
		return nil, err
	}

	var returnedTenants []*Tenant
	for _, tenant := range tenants {
		returnedTenants = append(returnedTenants, mapTenantFromModel(tenant))
	}

	return returnedTenants, nil
}
