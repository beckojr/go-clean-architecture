DROP TABLE IF EXISTS books ;
CREATE TABLE books (
    id BIGSERIAL NOT NULL,
    title VARCHAR(50) NOT NULL,
    description TEXT NOT NULL,
    author VARCHAR(50) NOT NULL,
    edition VARCHAR(50) NOT NULL,
    expires_at TIMESTAMP DEFAULT NULL,
    created_at TIMESTAMP DEFAULT NULL,
    modified_at TIMESTAMP DEFAULT NULL,
    deleted_at TIMESTAMP DEFAULT NULL,
    PRIMARY KEY (id)
);