CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TABLE tenants (
  id           UUID        PRIMARY KEY DEFAULT gen_random_uuid(),
  name         TEXT        NOT NULL UNIQUE,
  email        TEXT        NOT NULL,
  phone        TEXT,
  status       TEXT        NOT NULL DEFAULT 'active',
  created_at   TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at   TIMESTAMPTZ NOT NULL DEFAULT now()
);


CREATE FUNCTION set_timestamp()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = now();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trg_tenants_timestamp
BEFORE UPDATE ON tenants
FOR EACH ROW EXECUTE FUNCTION set_timestamp();
