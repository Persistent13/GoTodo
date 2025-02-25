CREATE TABLE users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    content TEXT NULL,
    createdAtUtc INT NOT NULL,
    updatedAtUtc INT NOT NULL,
    done INT NOT NULL,
    deleted INT NOT NULL
);
