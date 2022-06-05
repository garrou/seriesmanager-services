CREATE DATABASE seriesmanager;

CREATE TABLE users (
    id VARCHAR(50) PRIMARY KEY,
    username VARCHAR(50) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    joined_at TIMESTAMP NOT NULL
);

CREATE TABLE series (
    id VARCHAR(50) UNIQUE NOT NULL,
    aid INTEGER NOT NULL,
    title VARCHAR(100) NOT NULL,
    poster VARCHAR(150),
    episode_length INTEGER NOT NULL,
    added_at TIMESTAMP NOT NULL,
    user_id VARCHAR(50) REFERENCES users(id) ON DELETE CASCADE,
    PRIMARY KEY (aid, fk_user)
);

CREATE TABLE seasons (
    id VARCHAR(50) UNIQUE NOT NULL,
    number INTEGER NOT NUll,
    episodes INTEGER NOT NULL,
    image VARCHAR(150),
    started_at DATE NOT NULL,
    finished_at DATE NOT NULL,
    fk_series VARCHAR(50) REFERENCES series (id) ON DELETE CASCADE
);