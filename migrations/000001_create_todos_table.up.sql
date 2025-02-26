CREATE TABLE todos (
                       id INTEGER PRIMARY KEY AUTOINCREMENT,
                       content TEXT NULL,
                       createdAtUtc INT NOT NULL,
                       updatedAtUtc INT NOT NULL,
                       done INT NOT NULL,
                       isDeleted INT NOT NULL
);
