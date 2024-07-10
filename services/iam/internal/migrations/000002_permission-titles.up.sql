BEGIN;
CREATE TYPE permission_titles AS ENUM (
    'iam.viewer',
    'iam.admin',
    'iam.editor',
    'iam.self.editor',
    'iam.self.viewer',
    'sudoku.solver',
    'sudoku.viewer',
    'sudoku.admin'
);
ALTER TABLE permissions
ALTER COLUMN title TYPE permission_titles USING title::permission_titles;
INSERT INTO permissions (title, description)
VALUES (
        'iam.viewer',
        'Позволяет просматривать список пользователей'
    ),
    (
        'iam.admin',
        'Позволяет редактировать, добавлять и изменять других пользователей'
    ),
    (
        'iam.editor',
        'Позволяет редактировать других пользователей'
    ),
    (
        'iam.self.editor',
        'Позволяет изменять личные данные пользователя'
    ),
    (
        'iam.self.viewer',
        'Позволяет просматривать личные данные пользователя'
    ),
    ('sudoku.solver', 'Позволяет решать sudoku'),
    (
        'sudoku.viewer',
        'Позволяет просматривать список решенных другими судоку'
    ),
    (
        'sudoku.admin',
        'Позволяет удалять sudoku пользователя'
    );
COMMIT;