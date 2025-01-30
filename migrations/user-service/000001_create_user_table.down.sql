-- Drop the users table if it exists
DROP TABLE IF EXISTS users;

-- Drop the trigger for updating the timestamp
DROP TRIGGER IF EXISTS update_user_timestamp ON users;

-- Drop the trigger function
DROP FUNCTION IF EXISTS update_updated_at_column;

-- Drop UUID extension (optional, only if not used elsewhere)
DROP EXTENSION IF EXISTS "uuid-ossp";