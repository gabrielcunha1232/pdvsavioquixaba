package repository

import (
	"context"
	"uuid"

	"github.com/cunhazadanodale/pdv/internal/core/users/domain"
)

type UserRepository interface {
	GetByID(ctx context.Context, id uuid.UUID) (*domain.Users, error)
	Create(ctx context.Context, user domain.Users) error
	Me(ctx context.Context, id uuid.UUID) (*domain.Users, error)
	LoginWaiter(ctx context.Context, pin string, password string) error
	LoginCEO(ctx context.Context, email string) error
	UpdateUser(ctx context.Context, id uuid.UUID, user domain.Users)
	Desactive(ctx context.Context, id uuid.UUID) error
}