package tenants

import (
	"context"
	"xhantos/common"
)

type TenantService struct {
	repo TenantRepositoryInterface
}

// NewTenantService creates a new instance of TenantService.
func NewTenantService(repo TenantRepositoryInterface) *TenantService {
	return &TenantService{
		repo: repo,
	}
}

func (s *TenantService) CreateTenant(ctx context.Context, tenant *Tenant) (*Tenant, error) {
	return s.repo.Create(ctx, tenant)
}

func (s *TenantService) UpdateTenant(ctx context.Context, tenant *Tenant) error {
	// Implement UpdateTenant
	return s.repo.Update(ctx, tenant)
}

func (s *TenantService) DeleteTenant(ctx context.Context, tenantID int) error {
	// Implement DeleteTenant
	return s.repo.Delete(ctx, tenantID)
}

func (s *TenantService) GetTenantByID(ctx context.Context, tenantID int) (*Tenant, error) {
	// Implement GetTenantByID
	return s.repo.FindOne(ctx, tenantID)
}

func (s *TenantService) GetTenantsByQuery(ctx context.Context, queryParams common.QueryParams) ([]*Tenant, error) {
	// Implement GetAllTenants
	return s.repo.FindMany(ctx, queryParams)
}
