ALTER TABLE reminder ADD CONSTRAINT unique_reminder UNIQUE (chat_id, chat_message, unique_date, unique_time)