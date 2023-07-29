package tenants

import (
	"context"
)

type TenantService struct {
	repo TenantInterface
}

// NewTenantService creates a new instance of TenantService.
func NewTenantService(repo TenantInterface) *TenantService {
	return &TenantService{
		repo: repo,
	}
}

func (s *TenantService) CreateTenant(ctx context.Context, tenant *Tenant) (*Tenant, error) {
	return s.repo.Create(ctx, tenant)
}

func (s *TenantService) UpdateTenant(ctx context.Context, tenant *Tenant) (int, error) {
	// Implement UpdateTenant
	return s.repo.Update(ctx, tenant)
}

func (s *TenantService) DeleteTenant(ctx context.Context, tenantID int) error {
	// Implement DeleteTenant
	return nil
}

func (s *TenantService) GetTenantByID(ctx context.Context, tenantID int) (*Tenant, error) {
	// Implement GetTenantByID
	return nil, nil
}

func (s *TenantService) GetAllTenants(ctx context.Context) ([]*Tenant, error) {
	// Implement GetAllTenants
	return nil, nil
}
