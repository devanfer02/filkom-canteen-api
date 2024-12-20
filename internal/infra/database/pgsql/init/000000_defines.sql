CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE OR REPLACE FUNCTION generate_ulid() RETURNS uuid
    AS $$
        SELECT (lpad(to_hex(floor(extract(epoch FROM clock_timestamp()) * 1000)::bigint), 12, '0') || encode(gen_random_bytes(10), 'hex'))::uuid;
    $$ LANGUAGE SQL;

CREATE TYPE MENU_STATUS_ENUM AS ENUM('Ada', 'Habis');

CREATE TYPE PAYMENT_METHOD_ENUM AS ENUM('COD', 'QRIS');

-- CREATE ADMIN DATA AND ROLES

INSERT INTO roles (role_name) VALUES ('Admin'), ('Owner');