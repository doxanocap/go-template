CREATE TABLE exchangers_keys (
    id SERIAL PRIMARY KEY,
    city VARCHAR(100),
    name VARCHAR(100)
);

CREATE TABLE exchangers_currencies (
    exchanger_id INT,
    upload_time BIGINT,
    usd_buy FLOAT,
    usd_sell FLOAT,
    eur_buy FLOAT,
    eur_sell FLOAT,
    rub_buy FLOAT,
    rub_sell FLOAT
);

CREATE TABLE exchangers_info (
    exchanger_id INT,
    address TEXT,
    wholesale TEXT,
    updated_time BIGINT,
    phone_numbers TEXT
);

CREATE INDEX exchangers_currencies_exchanger_id_idx
    ON exchangers_currencies (exchanger_id);

CREATE INDEX exchangers_info_exchanger_id_idx
    ON exchangers_info (exchanger_id);

CREATE TABLE chat_list (
    id SERIAL PRIMARY KEY,
    chat_uid VARCHAR(36),
    message_qty INT,
    started_at BIGINT,
    blocked BOOLEAN
);

CREATE INDEX chat_list_id
    ON chat_list (chat_uid);

CREATE TABLE chat_messages (
    chat_id INT REFERENCES chat_list(id),
    UNIQUE(chat_id),
    sender_id INT,
    sent_at BIGINT,
    message TEXT
);

CREATE TABLE chat_users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) NOT NULL,
    username VARCHAR(255) NOT NULL,
    phone_number TEXT,
    online BOOLEAN
);

CREATE TABLE users (
   id SERIAL PRIMARY KEY,
   email VARCHAR(255) NOT NULL,
   username VARCHAR(255) NOT NULL,
   phone_number VARCHAR(255) NOT NULL, 
   is_activated Boolean DEFAULT FALSE,
   password TEXT NOT NULL
);

CREATE TABLE user_props (
    user_id INT REFERENCES users(id),
    UNIQUE(user_id),
    user_uid VARCHAR(38) NOT NULL,
    city VARCHAR(128) NOT  NULL
);

CREATE TABLE tokens (
    token_id INT REFERENCES users(id),
    UNIQUE(token_id),
    refreshToken TEXT NOT NULL
);