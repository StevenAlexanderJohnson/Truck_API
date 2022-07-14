CREATE TABLE IF NOT EXISTS user_identity.roles(
    user_email varchar(255),
    role_id INT,
    role_name VARCHAR(255) NOT NULL,
    FOREIGN KEY (user_email) REFERENCES user_identity.user_credentials(user_email) ON DELETE CASCADE 
);