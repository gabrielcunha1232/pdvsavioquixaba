CREATE TABLE IF NOT EXISTS tenants(
	id UUID PRIMARY KEY,
	name TEXT NOT NULL,
	slug TEXT NOT NULL UNIQUE,
	created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS users(
	id UUID PRIMARY KEY,
	tenant_id UUID NOT NULL REFERENCES tenants(id),
	name TEXT NOT NULL,
	role TEXT NOT NULL CHECK(role IN("owner", "manager", "cashier", "waiter")),
	email TEXT,
	password_hash TEXT,
	pin_hash TEXT,
	active BOOLEAN NOT NULL DEFAULT true
);

CREATE TABLE IF NOT EXISTS tables(
	id UUID PRIMARY KEY,
	tenant_id UUID NOT NULL REFERENCES tenants(id),
	label TEXT NOT NULL,
	status TEXT NOT NULL DEFAULT "free" CHECK(status IN("free", "occupied", "awaiting_payment")),
	UNIQUE (tenant_id, label)
);

CREATE TABLE IF NOT EXISTS categories(
	id UUID PRIMARY KEY,
	tenant_id UUID NOT NULL REFERENCES tenants(id),
	name TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS products(
	id UUID PRIMARY KEY,
	tenant_id UUID NOT NULL REFERENCES tenants(id),
	category_id UUID NOT NULL REFERENCES category(id),
	name TEXT NOT NULL,
	price_cents INTEGER NOT NULL CHECK (price_cents >= 0),
	active BOOLEAN NOT NULL DEFAULT true
);

CREATE TABLE IF NOT EXISTS tabs(
	id UUID PRIMARY KEY,
	tenant_id UUID NOT NULL REFERENCES tenants(id),
	table_id UUID NOT NULL REFERENCES tables(id),
	waiter_id UUID NOT NULL REFERENCES users(id),
	status TEXT NOT NULL DEFAULT 'open' CHECK(status IN("open", "closed")),
	opened_at TIMESTAMPTZ NOT NULL DEFAULT now(),
	closed_at TIMESTAMPTZ,
	closed_by UUID REFERENCES users(id),
	reopened_by UUID REFERENCES users(id)
);

CREATE UNIQUE INDEX open_open_tab_per_table ON tabs(table_id) WHERE status = 'open';

CREATE TABLE IF NOT EXISTS tab_items(
	id UUID PRIMARY KEY,
	tab_id NOT NULL REFERENCES tabs(id),
	product_id NOT NULL REFERENCES products(id),
	quantity INTEGER NOT NULL CHECK (quantity > 0),
	unit_price_cents INTEGER NOT NULL,
	note TEXT,
	status TEXT NOT NULL DEFAULT 'sent' CHECK (status IN("sent", "preparing", "ready", "delivered", "canceled")),
	created_by UUID NOT NULL REFERENCES users(id)
	created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);