package domain

import (
	"uuid"
	"time")

type Tenants struct {
	ID uuid.UUID `db:"id"`
	Name string `db:"name"`
	Slug string `db:"slug"`
	CreatedAt time.Time `db:"created_at"`
	Active bool `db:"active"`
}