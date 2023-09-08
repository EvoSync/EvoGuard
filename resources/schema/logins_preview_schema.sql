
-- logins is the structure which is used to store every
-- login we have overviewed, a login inside this program
-- is classed as a token creation event.
CREATE TABLE `logins` (
      `email` TEXT NOT NULL,
      `token` TEXT NOT NULL,
      `username` TEXT NOT NULL,
      `remoteAddress` TEXT NOT NULL
);

INSERT INTO `logins` (`email`, `token`, `username`, `remoteAddress`) VALUES (?, ?, ?, ?)