package auth

import "context"

type AuthService struct {
	repo RepositoryInterface
}

// NewAuthService creates a new instance of TenantService.
func NewAuthService(repo RepositoryInterface) *AuthService {
	return &AuthService{
		repo: repo,
	}
}

func (a AuthService) setTenant(ctx context.Context, userId, tenant string) error {
	return nil
}
