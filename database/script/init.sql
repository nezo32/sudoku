CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE IF NOT EXISTS users (
    id uuid DEFAULT uuid_generate_v4(),
    username VARCHAR(20) NOT NULL UNIQUE,
    password TEXT NOT NULL,
    PRIMARY KEY (id),
    CONSTRAINT user_username CHECK(username != '')
);
CREATE TABLE IF NOT EXISTS games (
    id uuid DEFAULT uuid_generate_v4(),
    player_id uuid,
    board CHAR [] [] NOT NULL,
    begin_date date,
    PRIMARY KEY (id)
);
ALTER TABLE games
ADD FOREIGN KEY (player_id) REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE;