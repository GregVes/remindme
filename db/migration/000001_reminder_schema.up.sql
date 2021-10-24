CREATE TABLE IF NOT EXISTS reminder (
    id serial PRIMARY KEY,
    chat_id INT NOT NULL,
    chat_message VARCHAR NOT NULL,
    is_recurrent BOOLEAN DEFAULT FALSE,
    is_everyday BOOLEAN DEFAULT FALSE,
    weekly_day VARCHAR(10) DEFAULT NULL,
    monthly_day INT DEFAULT NULL,
    annual_date VARCHAR(30) DEFAULT NULL,
    unique_date date,
    unique_time time NOT NULL,
    CONSTRAINT valid_month_date CHECK (monthly_day BETWEEN 1 AND 31)
);