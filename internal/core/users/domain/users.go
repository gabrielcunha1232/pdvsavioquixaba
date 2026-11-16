package domain

import "uuid"

type Users struct{
	ID uuid.UUID `db:"users"`
	TenantID uuid.UUID `db:"tenant_id"`
	Name string `db:"name"`
	PasswordHash string `db:"password_hash"`
	PinHash string `db:"pin_hash"`
	Role UserRole `db:"user_role"`
	Email string `db:"email"`
} 

type UserRole string

const(
	OwnerRole UserRole = "owner"
	ManagerRole UserRole = "manager"
	CashierRole UserRole = "cashier"
	WaiterRole UserRole = "waiter"
)

