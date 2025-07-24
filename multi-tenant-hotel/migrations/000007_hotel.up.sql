CREATE TABLE hotel_info (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    tenant_id UUID NOT NULL,

    name TEXT NOT NULL CHECK (char_length(name) <= 150),
    description TEXT,
    star_rating SMALLINT CHECK (star_rating BETWEEN 1 AND 5),


    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now()
);


CREATE TABLE hotel_contacts (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    hotel_id UUID NOT NULL REFERENCES hotel_info(id) ON DELETE CASCADE,

    emails TEXT[] NOT NULL,
    phones TEXT[],
    websites TEXT,

    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now()
);

CREATE TABLE hotel_locations (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    hotel_id UUID NOT NULL REFERENCES hotel_info(id) ON DELETE CASCADE,

    address TEXT,
    city TEXT,
    state TEXT,
    country TEXT,
    zip_code TEXT,
    latitude DECIMAL(9,6),
    longitude DECIMAL(9,6),

    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now()
);

CREATE TABLE hotel_announcements (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    tenant_id UUID NOT NULL REFERENCES tenants(id) ON DELETE CASCADE,
    title TEXT NOT NULL,
    message TEXT NOT NULL,
    type TEXT CHECK (type IN ('greeting', 'breakfast', 'event', 'custom')),
    schedule_at TIMESTAMPTZ,
    repeat_daily BOOLEAN DEFAULT false,
    is_active BOOLEAN DEFAULT true,
    created_at TIMESTAMPTZ DEFAULT now(),
    updated_at TIMESTAMPTZ DEFAULT now()
)

