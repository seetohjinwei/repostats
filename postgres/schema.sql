/*
  Schema to set up the database.

  psql -U USERNAME -d DATABASE -a -f FILE.sql
  psql -U deploy -d repostats -a -f FILE.sql
  psql -U postgres -d repostats -a -f postgres/schema.sql
*/

DROP TABLE IF EXISTS Users, Repositories, TypeData CASCADE;

CREATE TABLE Users (
  username TEXT,
  last_updated TIMESTAMP,
  PRIMARY KEY (username)
);

CREATE TABLE Repositories (
  username TEXT REFERENCES Users (username),
  repo TEXT,
  last_updated TIMESTAMP,
  default_branch TEXT NOT NULL,
  PRIMARY KEY (username, repo)
);

CREATE TABLE TypeData (
  username TEXT,
  repo TEXT,
  language TEXT,
  file_count INT NOT NULL,
  bytes INT NOT NULL,
  PRIMARY KEY (username, repo, language),
  FOREIGN KEY (username, repo) REFERENCES Repositories (username, repo)
);

/**
  * Adds a repository, with the provided username, repo, default_branch.
  *
  * Creates the user, if the user does not exist.
*/
CREATE OR REPLACE PROCEDURE add_repo(username TEXT, repo TEXT, default_branch TEXT)
AS $$
DECLARE
  user_count INT;
BEGIN
  SELECT COUNT(*) INTO user_count FROM Users U WHERE U.username = $1;

  IF user_count < 1 THEN
    -- create user, if not exists
    INSERT INTO Users VALUES ($1);
  END IF;

  INSERT INTO Repositories VALUES ($1, $2, NULL, $3);
END; $$ LANGUAGE plpgsql;
