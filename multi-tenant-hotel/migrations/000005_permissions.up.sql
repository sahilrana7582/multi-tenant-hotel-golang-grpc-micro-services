CREATE TYPE permission_action AS ENUM (
  'view',
  'create',
  'update',
  'delete',
  '*'
);


CREATE TABLE permissions (
  id            UUID               PRIMARY KEY DEFAULT uuid_generate_v4(),
  tenant_id     UUID               NOT NULL REFERENCES tenants(id) ON DELETE CASCADE,
  role_id       UUID               NOT NULL REFERENCES roles(id) ON DELETE CASCADE,
  department_id UUID               REFERENCES departments(id) ON DELETE CASCADE,
  action        permission_action  NOT NULL,
  UNIQUE (tenant_id, role_id, department_id, action)
);
