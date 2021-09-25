CREATE EXTENSION IF NOT EXISTS "pgcrypto";
CREATE TABLE users (
                             id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
                             email character varying(50) NOT NULL,
                             first_name character varying(36),
                             last_name character varying(36)
)
