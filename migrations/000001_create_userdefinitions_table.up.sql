CREATE TABLE user_definitions (
                                 user_id UUID PRIMARY KEY,
                                 nickname TEXT NOT NULL,
                                 birthdate DATE NOT NULL,
                                 preferred_sport JSONB NOT NULL,
                                 other_sports JSONB NOT NULL
);
