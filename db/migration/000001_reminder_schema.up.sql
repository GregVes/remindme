CREATE TABLE IF NOT EXISTS reminder (
  id serial PRIMARY KEY,
  chat_id INT NOT NULL,
  chat_message VARCHAR NOT NULL,
  is_reccurent BOOLEAN DEFAULT FALSE,
  target_recurrent_date date DEFAULT NULL,
  target_recurrent_day VARCHAR(10) DEFAULT NULL,
  target_date date,
  target_time time NOT NULL
)