-- +goose Up
-- Create skeleton_items table
CREATE TABLE skeleton_items (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    description TEXT,
    user_id UUID NOT NULL,
    active BOOLEAN DEFAULT true,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ
);

-- Create indexes for better query performance
CREATE INDEX idx_skeleton_items_user_id ON skeleton_items(user_id);
CREATE INDEX idx_skeleton_items_created_at ON skeleton_items(created_at DESC);
CREATE INDEX idx_skeleton_items_active ON skeleton_items(active);

-- Add comment to table for documentation
COMMENT ON TABLE skeleton_items IS 'Example items table for skeleton plugin demonstration';

-- +goose Down
-- Drop table and all related objects
DROP TABLE IF EXISTS skeleton_items CASCADE;
