CREATE TABLE user (
   uuid UUID PRIMARY KEY,
   email VARCHAR(255) NOT NULL,
   username VARCHAR(255) NOT NULL,
   phone_number VARCHAR(255) NOT NULL,
   password VARCHAR(255) NOT NULL
);