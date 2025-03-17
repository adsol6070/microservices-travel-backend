-- Enable UUID extension if not already enabled
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Create the blogs table
CREATE TABLE blogs (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),  
    title VARCHAR(150) NOT NULL,                    
    slug VARCHAR(255) UNIQUE NOT NULL,              
    content TEXT NOT NULL,                   
    excerpt VARCHAR(300),
    meta_title VARCHAR(150),
    meta_description VARCHAR(300),
    author_id UUID NOT NULL,                         
    category VARCHAR(100) NOT NULL,                       
    tags TEXT[],                                    
    thumbnail TEXT,                         
    status VARCHAR(20) NOT NULL DEFAULT 'draft',        
    published_at TIMESTAMPTZ,      
    scheduled_at TIMESTAMPTZ,
    is_published BOOLEAN DEFAULT FALSE,               
    created_at TIMESTAMPTZ DEFAULT now(),  
    updated_at TIMESTAMPTZ DEFAULT now(),
    deleted_at TIMESTAMPTZ NULL,  -- Nullable for soft delete
    
    CONSTRAINT fk_author FOREIGN KEY (author_id) REFERENCES users(id) ON DELETE CASCADE
);

-- Create an index for deleted_at to speed up queries for soft deletes
CREATE INDEX idx_blogs_deleted_at ON blogs(deleted_at);

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
