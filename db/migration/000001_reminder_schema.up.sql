CREATE TABLE IF NOT EXISTS reminder (
  id serial PRIMARY KEY,
  chat_id INT NOT NULL,
  chat_message VARCHAR NOT NULL,
  recurrent_day VARCHAR(20) DEFAULT NULL,
  target_date date,
  target_time time NOT NULL
);