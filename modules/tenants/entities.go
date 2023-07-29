package tenants

import (
	"xhantos/sqlboiler/models"
)

type Tenant struct {
	ID     int    `json:"id,omitempty"`
	Alias  string `json:"alias"`
	APIKey string `json:"api_key"`
	Status string `json:"status"`
}

// mapTenantFromModel converts the SQLBoiler-generated *models.Tenant to the Tenant interface.
func mapTenantFromModel(modelTenant *models.Tenant) Tenant {
	return Tenant{
		ID:     modelTenant.ID,
		Alias:  modelTenant.Alias,
		APIKey: modelTenant.APIKey,
		Status: string(modelTenant.Status),
	}
}

// mapTenantToModel converts the Tenant interface to the SQLBoiler-generated *models.Tenant.
func mapTenantToModel(tenant Tenant) *models.Tenant {
	return &models.Tenant{
		ID:     tenant.ID,
		Alias:  tenant.Alias,
		APIKey: tenant.APIKey,
		Status: models.TenantStatus(tenant.Status),
	}
}
