BEGIN;
ALTER TABLE permissions
ALTER COLUMN title TYPE VARCHAR(255);
DELETE FROM user_permissions;
DELETE FROM permissions;
DROP TYPE IF EXISTS permission_titles;
COMMIT;