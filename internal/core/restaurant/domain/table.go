package domain

import "uuid"

type Table struct {
	ID uuid.UUID `db:"id"`
	TenantID uuid.UUID `db:"tenant_id"`
	Label string `db:"label"`
	Status TableStatus `db:"status"`
}

type TableStatus string

const (
	Free TableStatus = "free"
	Occupied TableStatus = "occupied"
	AwaitingPayment TableStatus = "awating_payment"
)