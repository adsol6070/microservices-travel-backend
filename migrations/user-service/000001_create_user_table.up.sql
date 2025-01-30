-- Enable UUID extension if not already enabled
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Create the users table
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),  -- Unique user ID
    email VARCHAR(255) UNIQUE NOT NULL,              -- User email (unique)
    name VARCHAR(255) NOT NULL,                      -- User full name
    password TEXT NOT NULL,                          -- Hashed password (secured)
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,  -- Timestamp when the user was created
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP   -- Timestamp when the user was last updated
);

-- Trigger function to update updated_at on row update
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Create trigger to automatically update updated_at column
CREATE TRIGGER update_user_timestamp
BEFORE UPDATE ON users
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();