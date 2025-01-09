ALTER TABLE 
    posts 
ADD COLUMN tags VARCHAR(100) ARRAY;

ALTER TABLE
    posts
ADD COLUMN updated_at timestamp(0) with time zone NOT NULL DEFAULT NOW()