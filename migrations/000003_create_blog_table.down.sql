-- Drop the many-to-many relationship table for blog tags first to prevent FK errors
DROP TABLE IF EXISTS blog_tags;

-- Drop the trigger for updating the timestamp
DROP TRIGGER IF EXISTS update_blog_timestamp ON blogs;

-- Drop the trigger function
DROP FUNCTION IF EXISTS update_blog_updated_at_column;

-- Drop the blogs table
DROP TABLE IF EXISTS blogs;

-- Drop the ENUM type (Only if not used elsewhere)
DROP TYPE IF EXISTS blog_status;

-- Drop UUID extension (optional, only if not used elsewhere)
DROP EXTENSION IF EXISTS "uuid-ossp";