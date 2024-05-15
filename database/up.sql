DROP TABLE IF EXISTS users;

CREATE TABLE users (

    id VARCHAR(255) PRIMARY KEY,
    password VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL, 
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

DROP table IF EXISTS posts;

CREATE TABLE posts (
    id VARCHAR(255) PRIMARY KEY,
    post_content varchar(200) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    user_id VARCHAR(255) NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id)
)

