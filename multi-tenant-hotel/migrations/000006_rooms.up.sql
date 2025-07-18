CREATE TABLE room_types (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  name TEXT NOT NULL UNIQUE,
  description TEXT,
  created_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE rooms (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

  tenant_id UUID NOT NULL REFERENCES tenants(id) ON DELETE CASCADE,
  department_id UUID NOT NULL REFERENCES departments(id) ON DELETE CASCADE,
  room_type_id UUID NOT NULL REFERENCES room_types(id),

  room_number TEXT NOT NULL,
  floor INTEGER NOT NULL,
  price_per_night NUMERIC(10, 2) NOT NULL,

  status TEXT NOT NULL DEFAULT 'available' CHECK (status IN ('available', 'booked', 'maintenance', 'inactive')),
  is_active BOOLEAN DEFAULT TRUE,

  description TEXT,
  created_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP DEFAULT NOW(),

  UNIQUE (department_id, room_number)
);





ALTER TABLE rooms ENABLE ROW LEVEL SECURITY;

CREATE POLICY create_room_policy ON rooms
  FOR INSERT
  WITH CHECK (
    EXISTS (
      SELECT 1
      FROM permissions p
      JOIN roles r ON r.id = p.role_id
      JOIN user_roles ur ON ur.role_id = r.id
      WHERE
        ur.user_id = current_setting('app.current_user_id')::UUID
        AND (
          p.department_id = rooms.department_id OR
          p.department_id IS NULL
        )
        AND (
          p.action = 'create' OR
          p.action = '*'
        )
    )
);
