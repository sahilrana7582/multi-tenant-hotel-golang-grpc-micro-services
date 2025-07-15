CREATE TABLE features (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  key TEXT NOT NULL UNIQUE  -- e.g., 'can_invite_users'
);

CREATE TABLE tenant_features (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  tenant_id UUID NOT NULL REFERENCES tenants(id),
  feature_id UUID NOT NULL REFERENCES features(id),
  enabled BOOLEAN NOT NULL DEFAULT true,
  UNIQUE (tenant_id, feature_id)
);
