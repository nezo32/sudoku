BEGIN;
CREATE TYPE roles_titles AS ENUM (
    'iam.viewer',
    'iam.admin',
    'iam.editor',
    'sudoku.solver',
    'sudoku.viewer',
    'sudoku.admin'
);
CREATE TABLE roles (
    id UUID NOT NULL PRIMARY KEY DEFAULT gen_random_uuid(),
    title roles_titles NOT NULL
);
CREATE TABLE users (
    id UUID NOT NULL PRIMARY KEY DEFAULT gen_random_uuid(),
    username VARCHAR(64) NOT NULL UNIQUE,
    first_name VARCHAR(255),
    last_name VARCHAR(255),
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);
CREATE TABLE user_roles (
    id UUID NOT NULL PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    role_id UUID NOT NULL REFERENCES roles(id) ON DELETE CASCADE
);
INSERT INTO roles (title)
VALUES ('iam.viewer'),
    ('iam.admin'),
    ('iam.editor'),
    ('sudoku.solver'),
    ('sudoku.viewer'),
    ('sudoku.admin');
COMMIT;