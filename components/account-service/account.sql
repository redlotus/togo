CREATE EXTENSION pgcrypto;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE OR REPLACE FUNCTION update_timestamp()   
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at := now();
    RETURN NEW;   
END;
$$ language 'plpgsql';

CREATE TABLE IF NOT EXISTS account (
    id UUID NOT NULL DEFAULT uuid_generate_v4(),
    name VARCHAR(255) DEFAULT 'foo',
    email VARCHAR(255) IS NOT NULL,
    password TEXT,
    status SMALLINT DEFAULT 1, -- -1: DELETED, 0: UNACTIVE, 1: ACTIVE - default should be 0 before validation
    settings JSON,
    roles text[] DEFAULT '{user}',
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY (id)
);

DROP TRIGGER IF EXISTS update_account_updated_at ON account;
CREATE TRIGGER update_account_updated_at BEFORE UPDATE ON account FOR EACH ROW EXECUTE PROCEDURE  update_timestamp();

-- reserve for later
-- CREATE TABLE IF NOT EXISTS social_account (
--     id UUID UNIQUE NOT NULL,
--     account_id UUID REFERENCES account(id),
--     status INTEGER DEFAULT 1, -- 1: ACTIVE, 0: UNACTIVE
--     token TEXT, -- social token
--     type VARCHAR(10), -- Type social, example: fb, gg
--     extras JSON,
--     updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
--     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
-- );
-- DROP TRIGGER IF EXISTS update_social_account_updated_at ON social_account;
-- CREATE TRIGGER update_social_account_updated_at BEFORE UPDATE ON social_account FOR EACH ROW EXECUTE PROCEDURE  update_timestamp();

-- CREATE TABLE IF NOT EXISTS tag (
--     id SERIAL PRIMARY KEY,
--     name TEXT,
--     description TEXT,
--     type VARCHAR(10),
--     updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
--     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
-- );

-- CREATE TABLE IF NOT EXISTS account_tag (
--     id SERIAL PRIMARY KEY,
--     account_id UUID REFERENCES account(id),
--     device_id UUID,
--     tag_id INTEGER REFERENCES tag(id),
--     expire_at TIMESTAMP,
--     updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
--     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
-- );

-- CREATE TABLE IF NOT EXISTS avatar_image (
--     id SERIAL PRIMARY KEY NOT NULL,
--     location VARCHAR(255) NOT NULL,
--     status INTEGER DEFAULT 1, -- 0: INACTIVE 1: ACTIVE
--     updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
--     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
-- );

-- CREATE TRIGGER update_avatar_image_updated_at BEFORE UPDATE ON avatar_image FOR EACH ROW EXECUTE FUNCTION  update_timestamp();

CREATE TABLE IF NOT EXISTS session (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    account_id UUID,
    platform TEXT, -- for tracking, eg. web, mobile
    type_login VARCHAR(10) DEFAULT 'basic',
    status INTEGER, -- -1: REVOKED 0: INACTIVE 1: ACTIVE 2: DEACTIVE
    expire_at TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TRIGGER update_glxp_session_updated_at BEFORE UPDATE ON glxp_session FOR EACH ROW EXECUTE FUNCTION  update_timestamp();

-- functions and triggers
CREATE OR REPLACE FUNCTION encrypt_password()
    RETURNS TRIGGER 
AS $$
BEGIN
	NEW.password := crypt(OLD.password, gen_salt('bf'));
RETURN NEW;
END;
$$ language 'plpgsql';

DROP TRIGGER IF EXISTS encrypt_created_account_password ON account;
CREATE TRIGGER encrypt_created_account_password BEFORE INSERT ON account FOR EACH ROW EXECUTE FUNCTION  encrypt_password();
DROP TRIGGER IF EXISTS encrypt_updated_account_password ON account;
CREATE TRIGGER encrypt_updated_account_password BEFORE UPDATE ON account FOR EACH ROW EXECUTE FUNCTION  encrypt_password();

CREATE OR REPLACE FUNCTION check_password (em varchar(255), pwd varchar(255))
  RETURNS boolean
  AS $$
BEGIN
  IF EXISTS (
    SELECT
      TRUE
    FROM
      account
    WHERE
      email = em
      AND password = crypt(pwd, password)) THEN
  RETURN TRUE;
END IF;
  RETURN FALSE;
END;
$$
LANGUAGE 'plpgsql';