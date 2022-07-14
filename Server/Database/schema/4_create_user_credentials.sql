CREATE TABLE IF NOT EXISTS user_identity.user_credentials (
    user_email VARCHAR(255) NOT NULL,
    username VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    PRIMARY KEY (user_email)
);