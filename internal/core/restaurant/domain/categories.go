package domain

import "uuid"

type Category struct {
	ID uuid.UUID `db:"id"`
	TenantID uuid.UUID `db:"tenant_id"`
	Name string `db:"name"`
}