CREATE TYPE gender_type AS ENUM ('M', 'W');

CREATE TABLE user_definitions (
    user_id UUID PRIMARY KEY,
    nickname TEXT NOT NULL,
    birthdate DATE NOT NULL,
    `weight` REAL NOT NULL,
    height REAL NOT NULL,
    gender gender_type NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
