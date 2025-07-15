-- Drop the trigger
DROP TRIGGER IF EXISTS trg_tenants_timestamp ON tenants;

-- Drop the trigger function
DROP FUNCTION IF EXISTS set_timestamp;

-- Drop the tenants table
DROP TABLE IF EXISTS tenants;



