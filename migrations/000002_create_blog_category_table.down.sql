-- Drop the blog_categories table if it exists
DROP TABLE IF EXISTS blog_categories;

-- Drop the trigger for updating the timestamp
DROP TRIGGER IF EXISTS update_blog_category_timestamp ON blog_categories;

-- Drop the trigger function
DROP FUNCTION IF EXISTS update_blog_category_updated_at_column;

-- Drop UUID extension (optional, only if not used elsewhere)
DROP EXTENSION IF EXISTS "uuid-ossp";

-- Enable UUID extension if not already enabled
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";