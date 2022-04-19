CREATE TABLE users (
    id VARCHAR(50) PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL
);

CREATE TABLE series (
    id NUMERIC PRIMARY KEY,
    fk_user VARCHAR NOT NULL REFERENCES users(id)
);
/*
CREATE TABLE seasons (
    id SERIAL PRIMARY KEY,
    number NUMERIC NOT NULL,
    start_at DATE NOT NULL,
    finish_at DATE,
    ended BOOLEAN NOT NULL,
    fk_series NUMERIC NOT NULL REFERENCES series(id)
);

CREATE TABLE episodes (
    id NUMERIC PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    view_at DATE NOT NULL,
    fk_series NUMERIC NOT NULL REFERENCES series(id),
    fk_season INTEGER NOT NULL REFERENCES seasons(id)
);

CREATE TABLE favorites (
    fk_series NUMERIC NOT NULL REFERENCES series(id),
    fk_user VARCHAR NOT NULL REFERENCES users(id),
    PRIMARY KEY (fk_series, fk_user)
);*/