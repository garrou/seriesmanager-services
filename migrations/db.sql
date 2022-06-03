CREATE DATABASE seriesmanager;

CREATE TABLE users (
    id VARCHAR(50) PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL
);

CREATE TABLE series (
    id INTEGER NOT NULL,
    title VARCHAR(100) NOT NULL,
    poster VARCHAR(150),
    episode_length INTEGER NOT NULL,
    fk_user VARCHAR(50) REFERENCES users(id) ON DELETE CASCADE,
    sid SERIAL UNIQUE,
    PRIMARY KEY (id, fk_user)
);

CREATE TABLE seasons (
    id SERIAL PRIMARY KEY,
    number INTEGER NOT NUll,
    episodes INTEGER NOT NULL,
    image VARCHAR(150),
    started_at DATE NOT NULL,
    finished_at DATE NOT NULL,
    fk_series INTEGER REFERENCES series (sid) ON DELETE CASCADE
);