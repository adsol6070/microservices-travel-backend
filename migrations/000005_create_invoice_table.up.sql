-- Enable UUID extension if not already enabled
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Create the invoices table
CREATE TABLE invoices (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL,
    booking_id UUID NOT NULL,
    invoice_number VARCHAR(255) NOT NULL UNIQUE,
    currency VARCHAR(10) NOT NULL,
    amount DOUBLE PRECISION NOT NULL CHECK (amount > 0),
    tax_amount DOUBLE PRECISION DEFAULT 0,
    discount_amount DOUBLE PRECISION DEFAULT 0,
    final_amount DOUBLE PRECISION NOT NULL CHECK (final_amount > 0),
    status VARCHAR(20) NOT NULL CHECK (status IN ('paid', 'unpaid', 'cancelled', 'refunded')),
    payment_method VARCHAR(50),
    paid_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ DEFAULT now(),
    updated_at TIMESTAMPTZ DEFAULT now(),
    deleted_at TIMESTAMPTZ,

    CONSTRAINT fk_invoice_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- Create the invoice_package_items table
CREATE TABLE invoice_package_items (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    invoice_id UUID NOT NULL,
    package_name VARCHAR(255) NOT NULL,
    description TEXT,
    quantity INTEGER NOT NULL CHECK (quantity >= 1),
    unit_price DOUBLE PRECISION NOT NULL CHECK (unit_price > 0),
    total_price DOUBLE PRECISION NOT NULL CHECK (total_price > 0),

    CONSTRAINT fk_invoice FOREIGN KEY (invoice_id) REFERENCES invoices(id) ON DELETE CASCADE
);

-- Create index for deleted_at on invoices table
CREATE INDEX idx_invoices_deleted_at ON invoices(deleted_at);

-- Trigger function to update updated_at on invoices table
CREATE OR REPLACE FUNCTION update_invoice_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Create trigger for updating updated_at on invoices table
CREATE TRIGGER update_invoice_timestamp
BEFORE UPDATE ON invoices
FOR EACH ROW
EXECUTE FUNCTION update_invoice_updated_at_column();
