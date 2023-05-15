-- Set Environment Variables
\set userdb `echo "$APP_DB_USER"`
\set passdb `echo "$APP_DB_PASS"`
\set dbname `echo "$APP_DB_NAME"`

-- Create Database
CREATE DATABASE :"dbname";

-- Create User and Grant Privileges
CREATE USER :"userdb" WITH ENCRYPTED PASSWORD :'passdb';
GRANT ALL PRIVILEGES ON DATABASE :"dbname" TO :"userdb";
\connect :"dbname" :"userdb"

-- Create Extension and Install into Database
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Create Address Table
CREATE TABLE "addresses" (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    street VARCHAR(30) NOT NULL,
    city VARCHAR(30) NOT NULL,
    state VARCHAR(5) NOT NULL,
    zip_code VARCHAR(5) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Create User Table
CREATE TABLE "users" (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    first_name VARCHAR(25) NOT NULL,
    last_name VARCHAR(25) NOT NULL,
    email VARCHAR(50) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    address_id UUID REFERENCES "addresses" (id)
);

-- Create a Function to Update Timestamp at "updated_at" Column
CREATE OR REPLACE FUNCTION update_timestamp_updated_at()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = now();
    RETURN NEW;
END;
$$ language 'plpgsql';

-- Create a Trigger to Update the Given Tables Timestamp when their Rows are Updated
CREATE TRIGGER update_user_task_updated_at
    BEFORE UPDATE ON "users"
    FOR EACH ROW
EXECUTE PROCEDURE update_timestamp_updated_at();

CREATE TRIGGER update_address_task_updated_at
    BEFORE UPDATE ON "addresses"
    FOR EACH ROW
EXECUTE PROCEDURE update_timestamp_updated_at();

-- Insert Address Data
INSERT INTO "addresses" (street, city, state, zip_code) VALUES
    ('99/1 Street 1', 'Bangkok', 'AK', '10105');

-- Insert User Data
INSERT INTO "users" (first_name, last_name, email, address_id) VALUES
    ('John', 'Doe', 'john@mail.com', (SELECT id FROM "addresses" WHERE street = '99/1 Street 1'));
