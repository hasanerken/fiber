package tenants

import "xhantos/config"

func SetupTenantRoutes(app *config.Config) {
	repo := NewTenantRepository(app.DB)
	service := NewTenantService(repo)
	handler := NewTenantHandler(service)

	app.Server.Post("/tenants", handler.CreateTenant)
}
