package domain

import (
	"time"
	"uuid"
)

type TabsItems struct {
	ID uuid.UUID `db:"id"`
	TabID uuid.UUID `db:"tabs_id"`
	ProductID uuid.UUID `db:"product_id"`
	Quantity uint64 `db:"quantity"`
	UnitPriceCents uint64 `db:"unit_price_cents"`
	Note string `db:"note"`
	Status TabsItemsStatus `db:"status"`
	CreatedBy uuid.UUID `db:"created_by"`
	CreatedAt time.Time `db:"created_at"`
}

type TabsItemsStatus string

const (
	SentStatus TabsItemsStatus = "sent"
	PreparingStatus TabsItemsStatus = "preparing"
	ReadyStatus TabsItemsStatus = "ready"
	DeliveredStatus TabsItemsStatus = "delivered"
	CanceledStatus TabsItemsStatus = "canceled"
)