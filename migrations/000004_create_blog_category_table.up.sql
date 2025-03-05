-- Create the blog_categories table
CREATE TABLE blog_categories (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),  -- Unique category ID
    name VARCHAR(50) NOT NULL,                       -- Category name
    slug VARCHAR(255) UNIQUE NOT NULL,              -- Unique slug for SEO
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,  -- Created timestamp
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP   -- Last updated timestamp
);

-- Trigger function to update updated_at on row update
CREATE OR REPLACE FUNCTION update_blog_category_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Create trigger to automatically update updated_at column
CREATE TRIGGER update_blog_category_timestamp
BEFORE UPDATE ON blog_categories
FOR EACH ROW
EXECUTE FUNCTION update_blog_category_updated_at_column();