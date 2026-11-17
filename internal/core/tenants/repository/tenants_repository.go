package repository

import (
	"context"
	"uuid"

	"github.com/cunhazadanodale/pdv/internal/core/tenants/domain"
)


type TenantsRepository interface {
	MyTenant(ctx context.Context, id uuid.UUID) (*domain.Tenants, error)
	ListAllTenants(ctx context.Context, id uuid.UUID) ([]domain.Tenants, error)
	Create(ctx context.Context, tenant domain.Tenants) error
	Update(ctx context.Context, tenant domain.Tenants) (*domain.Tenants, error)
	Desactive(ctx context.Context, id uuid.UUID) (*domain.Tenants, error)
}