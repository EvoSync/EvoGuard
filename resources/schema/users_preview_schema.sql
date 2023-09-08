
-- users is the structure which is contained within a database and
-- holds pretty much everything we know about a users account.
CREATE TABLE `users` (
        `id`            INTEGER PRIMARY KEY AUTOINCREMENT,
        `username`      TEXT NOT NULL,
        `password`      BLOB NOT NULL,
        `email`         TEXT NOT NULL,
        `salt`          SALT NOT NULL,
        `accountLevel`  INTEGER NOT NULL,
        `parent`        INTEGER NOT NULL
);