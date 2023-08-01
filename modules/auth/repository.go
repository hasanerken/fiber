package auth

import (
	"context"
	"github.com/jmoiron/sqlx"
	"log"
	"xhantos/sqlboiler/models"
)

type RepositoryInterface interface {
	setTenant(ctx context.Context, userId, tenantId string) error
}

// NewRepository creates a new instance of AuthRepository.
func NewAuthRepository(db *sqlx.DB) *Repository {
	return &Repository{
		db: db,
	}
}

type Repository struct {
	db *sqlx.DB
}

func (r Repository) setTenant(ctx context.Context, userId, tenant string) error {
	query := `
		UPDATE authorizer_users
		SET tenant = :tenant
		WHERE id = :userID
	`

	// Parameters to bind to the query
	params := map[string]interface{}{
		"tenant": tenant,
		"userID": userId,
	}

	// Execute the update query
	_, err := r.db.NamedExec(query, params)
	if err != nil {
		log.Printf("Failed to update the Tenant %v", err)
		user := new(models.AuthorizerUser)
		user.ID = userId
		_, err := user.Delete(ctx, r.db)
		if err != nil {
			log.Printf("User created but Tenant can not be set user-id %v", userId)
		}
		return err
	}
	return nil
}
