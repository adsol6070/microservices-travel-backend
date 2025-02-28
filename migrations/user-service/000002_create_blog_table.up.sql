-- Enable UUID extension if not already enabled
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Create the blogs table
CREATE TABLE blogs (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),  -- Unique blog ID
    title VARCHAR(255) NOT NULL,                     -- Blog title
    slug VARCHAR(255) UNIQUE NOT NULL,               -- Unique slug for SEO
    content TEXT NOT NULL,                           -- Blog content
    author_id UUID NOT NULL,                         -- Reference to the user (author)
    tags TEXT[],                                     -- Tags as an array
    category VARCHAR(255),                           -- Blog category
    thumbnail VARCHAR(255),                          -- URL of the blog thumbnail
    published_at TIMESTAMP,                          -- Timestamp when the blog was published
    status VARCHAR(50) NOT NULL,                     -- Status (draft, published, etc.)
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,  -- Created timestamp
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,  -- Last updated timestamp
    FOREIGN KEY (author_id) REFERENCES users(id) ON DELETE CASCADE
);

-- Trigger function to update updated_at on row update
CREATE OR REPLACE FUNCTION update_blog_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Create trigger to automatically update updated_at column
CREATE TRIGGER update_blog_timestamp
BEFORE UPDATE ON blogs
FOR EACH ROW
EXECUTE FUNCTION update_blog_updated_at_column();