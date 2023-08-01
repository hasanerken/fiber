package auth

import (
	"xhantos/config"
)

func SetupAuthRoutes(app *config.Config) {
	repo := NewAuthRepository(app.DB)
	service := NewAuthService(repo)
	handler := NewAuthHandler(service, app.Authorizer, app.DB)

	app.Server.Post("/login", handler.Login)
	app.Server.Post("/register", handler.Register)
	app.Server.Post("/forgot-password", handler.ForgotPassword)
}
