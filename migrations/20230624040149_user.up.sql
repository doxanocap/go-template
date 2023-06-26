CREATE TABLE user (
   id SERIAL PRIMARY KEY,
   uuid uuid UNIQUE,
   email VARCHAR(255) NOT NULL,
   username VARCHAR(255) NOT NULL,
   phone_number VARCHAR(255) NOT NULL,
   password VARCHAR(255) NOT NULL
);


CREATE TABLE user_params (
    token_id INT REFERENCES user(id),
    refresh_token TEXT NOT NULL,
    updated_at TIMESTAMP DEFAULT NOW(),
    created_at TIMESTAMP DEFAULT NOW()
);

