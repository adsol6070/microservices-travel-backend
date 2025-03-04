-- Drop the blogs table if it exists
DROP TABLE IF EXISTS blogs;

-- Drop the trigger for updating the timestamp
DROP TRIGGER IF EXISTS update_blog_timestamp ON blogs;

-- Drop the trigger function
DROP FUNCTION IF EXISTS update_blog_updated_at_column;

-- Drop UUID extension (optional, only if not used elsewhere)
DROP EXTENSION IF EXISTS "uuid-ossp";