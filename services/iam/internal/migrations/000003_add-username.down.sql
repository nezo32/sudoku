BEGIN;
UPDATE users
SET first_name = ''
WHERE first_name IS NULL;
UPDATE users
SET last_name = ''
WHERE last_name IS NULL;
ALTER TABLE users
ALTER COLUMN first_name
SET NOT NULL;
ALTER TABLE users
ALTER COLUMN last_name
SET NOT NULL;
ALTER TABLE users DROP COLUMN IF EXISTS username;
COMMIT;