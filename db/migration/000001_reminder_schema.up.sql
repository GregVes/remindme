CREATE TABLE IF NOT EXISTS reminder (
  id serial PRIMARY KEY,
  chat_id INT NOT NULL,
  chat_message VARCHAR NOT NULL,
  is_recurrent BOOLEAN DEFAULT FALSE,
  is_everyday BOOLEAN DEFAULT FALSE,
  recurrent_week_day VARCHAR(10) DEFAULT NULL,
  recurrent_month_day INT DEFAULT NULL,
  recurrent_date VARCHAR(30) DEFAULT NULL,
  unique_date date,
  unique_time time NOT NULL,
  CONSTRAINT valid_month_date CHECK (recurrent_month_day BETWEEN 1 AND 31),
  UNIQUE (chat_id, chat_message)
)
