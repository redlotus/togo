-- the narrower the column for the better 
-- so smallint is chosen for indicator
CREATE EXTENSION pgcrypto;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE OR REPLACE FUNCTION update_timestamp()   
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at := now();
    RETURN NEW;   
END;
$$ language 'plpgsql';

CREATE TABLE IF NOT EXISTS board (
    id UUID NOT NULL DEFAULT uuid_generate_v4(),
    created_by UUID NOT NULL, -- user that create this board
    title VARCHAR(255) DEFAULT 'new',
    status SMALLINT DEFAULT 1, -- -1: DELETED, 0: LOCKED, 1: ACTIVE
    order_number SMALLINT DEFAULT 0,
    visibility SMALLINT DEFAULT 1, -- -1: UNLISTED, 0: PRIVATE, 1: PUBLIC - for sharing mode
    settings JSONB DEFAULT '{}'::JSONB, -- color or else
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS task (
    id UUID NOT NULL DEFAULT uuid_generate_v4(),
    created_by UUID NOT NULL,
    created_by UUID NOT NULL, -- user that create this task
    title VARCHAR(255),
    description TEXT,
    status SMALLINT DEFAULT 1, -- -1: DELETED, 0: LOCKED, 1: ACTIVE
    priority_level SMALLINT DEFAULT 0, -- 0: NORMAL, 1: IMPORTANT, 2: URGENT
    is_done BOOLEAN DEFAULT FALSE,
    order_number SMALLINT DEFAULT 0,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS tag (
    id SERIAL PRIMARY KEY, -- not gonna be exposed so no need for uuid
    name VARCHAR(255),
    type SMALLINT DEFAULT 1, -- 0: BOARD, 1: TASK
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS board_tag (
    board_id UUID REFERENCES board(id),
    tag_id INTEGER REFERENCES tag(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    PRIMARY KEY (board_id, tag_id)
);

CREATE TABLE IF NOT EXISTS task_tag (
    task_id UUID REFERENCES task(id),
    tag_id INTEGER REFERENCES tag(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    PRIMARY KEY (board_id, tag_id)
);

-- functions and triggers
DROP TRIGGER IF EXISTS update_board_updated_at ON board;
CREATE TRIGGER update_board_updated_at BEFORE UPDATE ON board FOR EACH ROW EXECUTE PROCEDURE  update_timestamp();

DROP TRIGGER IF EXISTS update_board_updated_at ON board;
CREATE TRIGGER update_board_updated_at BEFORE UPDATE ON board FOR EACH ROW EXECUTE PROCEDURE  update_timestamp();