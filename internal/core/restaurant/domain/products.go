package domain

import "uuid"

type Products struct {
	ID uuid.UUID `db:"id"`
	TenantID uuid.UUID `db:"tenant_id"`
	CategoryID uuid.UUID `db:"category_id"`
	Name string `db:"name"`
	PriceCents uint64 `db:"price_cents"`
	Active bool `db:"active"`
}