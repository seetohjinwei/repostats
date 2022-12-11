/*
  Schema to set up the database.

  psql -U USERNAME -d DATABASE -a -f FILE.sql
  psql -U deploy -d repostats -a -f FILE.sql
  psql -U postgres -d repostats -a -f schema.sql
*/

DROP TABLE IF EXISTS Users, Repositories, TypeData CASCADE;

CREATE TABLE Users (
  username TEXT,
  last_updated TIMESTAMP WITH TIME ZONE,
  PRIMARY KEY (username)
);

CREATE TABLE Repositories (
  username TEXT REFERENCES Users (username),
  repo TEXT,
  last_updated TIMESTAMP WITH TIME ZONE,
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
    -- Create user, if not exists.
    INSERT INTO Users VALUES ($1);
  END IF;

  INSERT INTO Repositories VALUES ($1, $2, NULL, $3);
END; $$ LANGUAGE plpgsql;

/**
  * Shape of `TypeData`.
*/
CREATE TYPE TypeDataShape AS (
  language TEXT,
  file_count INT,
  bytes INT
);

/**
  * Upserts TypeData for a repository.
  *
  * Creates the repository, if the repository does not exist.
*/
CREATE OR REPLACE PROCEDURE upsert_typedata(_username TEXT, _repo TEXT, _default_branch TEXT, _type_data TypeDataShape[])
AS $$
DECLARE
  repo_count INT;
  tempTable CONSTANT TEXT := '_temp_upsert_typedata';
  _td TypeDataShape;
BEGIN
  SELECT COUNT(*) INTO repo_count FROM Repositories R WHERE R.username = $1 AND R.repo = $2;

  IF repo_count < 1 THEN
    -- Create repository, if don't exist.
    CALL add_repo($1, $2, $3);
  END IF;

  CREATE TEMPORARY TABLE tempTable (LIKE TypeData INCLUDING ALL)
  ON COMMIT DROP;

  FOREACH _td IN ARRAY $4 LOOP
    INSERT INTO tempTable VALUES ($1, $2, _td.language, _td.file_count, _td.bytes);
  END LOOP;

  INSERT INTO TypeData (SELECT * FROM tempTable)
  ON CONFLICT (username, repo, language) DO UPDATE
		SET file_count = excluded.file_count,
			bytes = excluded.bytes;

  -- Update `last_updated` field.
  UPDATE Repositories
    SET last_updated = NOW()
    WHERE username = $1
      AND repo = $2;
END; $$ LANGUAGE plpgsql;

/*
--- FOR TESTING ---

CALL upsert_typedata('__user', '__repo', 'main', array[('java', 1, 420), ('go', 2, 34)]::TypeDataShape[]);
CALL upsert_typedata('__user', '__repo', 'main', array[('java', 2, 200), ('py', 111, 3)]::TypeDataShape[]);
SELECT * FROM TypeData;

DELETE FROM TypeData WHERE username = '__user' AND repo = '__repo';
DELETE FROM Repositories WHERE repo = '__repo';
DELETE FROM Users WHERE username = '__user';
*/