-- Drop the trigger for updating the timestamp on invoices table
DROP TRIGGER IF EXISTS update_invoice_timestamp ON invoices;

-- Drop the trigger function
DROP FUNCTION IF EXISTS update_invoice_updated_at_column;

-- Drop the invoice_package_items table if it exists
DROP TABLE IF EXISTS invoice_package_items;

-- Drop the invoices table if it exists
DROP TABLE IF EXISTS invoices;

-- Drop UUID extension (optional, only if not used elsewhere in your DB)
DROP EXTENSION IF EXISTS "uuid-ossp";
