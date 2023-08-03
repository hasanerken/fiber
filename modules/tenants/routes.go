package tenants

import (
	"xhantos/config"
	"xhantos/middlewares"
)

func SetupTenantRoutes(app *config.Config) {
	repo := NewTenantRepository(app.DB)
	service := NewTenantService(repo)
	handler := NewTenantHandler(service)

	app.Server.Post("/tenants", handler.CreateTenant)
	app.Server.Put("/tenants", handler.UpdateTenant)
	app.Server.Get("/tenants/:id", handler.GetTenantByID)
	app.Server.Get("/tenants/", middlewares.Authorize(*app.Authorizer, "user"), handler.GetTenantsByQuery)
	app.Server.Delete("/tenants/:id", handler.DeleteTenant)
}
