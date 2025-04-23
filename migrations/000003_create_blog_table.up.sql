-- Enable UUID extension if not already enabled
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- ENUM for blog status
CREATE TYPE blog_status AS ENUM ('draft', 'published', 'archived', 'deleted');

-- Create the blogs table
CREATE TABLE blogs (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),  
    title VARCHAR(150) NOT NULL,                    
    slug VARCHAR(255) UNIQUE NOT NULL,              
    content TEXT NOT NULL,                   
    excerpt VARCHAR(300),
    meta_title VARCHAR(150),
    meta_description VARCHAR(300),
    author VARCHAR(10) NOT NULL,                         
    category_id UUID NOT NULL REFERENCES blog_categories(id) ON DELETE CASCADE,                       
    thumbnail TEXT,                         
    status blog_status NOT NULL DEFAULT 'draft',        
    published_at TIMESTAMPTZ NULL,      
    scheduled_at TIMESTAMPTZ NULL,
    created_at TIMESTAMPTZ DEFAULT now(),  
    updated_at TIMESTAMPTZ DEFAULT now(),    
);

-- Many-to-Many table for blog tags 
CREATE TABLE blog_tags (
    blog_id UUID REFERENCES blogs(id) ON DELETE CASCADE,
    tag VARCHAR(50) NOT NULL,
    PRIMARY KEY (blog_id, tag)
);

-- Trigger function to update updated_at on row update
CREATE OR REPLACE FUNCTION update_blog_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    IF ROW(NEW.*) IS DISTINCT FROM ROW(OLD.*) THEN
        NEW.updated_at = NOW();
    END IF;    
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Create trigger to automatically update updated_at column
CREATE TRIGGER update_blog_timestamp
BEFORE UPDATE ON blogs
FOR EACH ROW
EXECUTE FUNCTION update_blog_updated_at_column();
