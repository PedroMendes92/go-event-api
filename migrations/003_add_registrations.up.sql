CREATE TABLE IF NOT EXISTS registrations (
    id INT PRIMARY KEY AUTO_INCREMENT,
    event_id INT REFERENCES events(id),
    user_id INT REFERENCES users(id)
);