package domain

import (
	"time"
	"uuid"
)

type Tabs struct {
	ID uuid.UUID `db:"id"` 
	TenantID uuid.UUID `db:"tenant_id"`
	TableID uuid.UUID `db:"table_id"`
	WaiterID uuid.UUID `db:"waiter_id"`
	Status TabsStatus `db:"status"`
	OpenedAt time.Time `db:"opened_at"`
	ClosedAt time.Time `db:"closed_at"`
	ClosedBy uuid.UUID `db:"closed_by"`
	ReopenedBy uuid.UUID `db:"reopened_by"`
}


type TabsStatus string

const (
	OpenTabs TabsStatus = "open"
	CloseTabs TabsStatus = "closed"
)