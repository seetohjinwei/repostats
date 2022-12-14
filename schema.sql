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
  username TEXT REFERENCES Users (username) ON DELETE CASCADE ON UPDATE CASCADE,
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
  FOREIGN KEY (username, repo) REFERENCES Repositories (username, repo) ON DELETE CASCADE ON UPDATE CASCADE
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
  SELECT COUNT(*) INTO user_count FROM Users U WHERE U.username ILIKE $1;

  IF user_count < 1 THEN
    -- Create user, if not exists.
    INSERT INTO Users VALUES (LOWER($1));
  END IF;

  INSERT INTO Repositories VALUES (LOWER($1), LOWER($2), NULL, LOWER($3));
END; $$ LANGUAGE plpgsql;

/**
  * Updates repositories for a user.
  *
  * Deletes the old repositories and adds the new repositories.
  * Creates the user, if the user does not exist.
*/
CREATE OR REPLACE PROCEDURE update_repos(_username TEXT, _repos Repositories[])
AS $$
DECLARE
  user_count INT;
  _repo Repositories;
BEGIN
  SELECT COUNT(*) INTO user_count FROM Users U WHERE U.username ILIKE $1;

  IF user_count < 1 THEN
    -- Create user, if not exists.
    INSERT INTO Users VALUES (LOWER($1));
  END IF;

  DELETE FROM Repositories WHERE username ILIKE $1;

  FOREACH _repo IN ARRAY $2 LOOP
    INSERT INTO Repositories VALUES (LOWER($1), LOWER(_repo.repo), NULL, LOWER(_repo.default_branch));
  END LOOP;

  -- Update `last_updated` field.
  UPDATE Users
    SET last_updated = NOW()
    WHERE username ILIKE $1;
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
  * Updates TypeData for a repository.
  *
  * Deletes the old TypeData and adds the new TypeData.
  * Creates the repository, if the repository does not exist.
*/
CREATE OR REPLACE PROCEDURE update_typedata(_username TEXT, _repo TEXT, _default_branch TEXT, _type_data TypeDataShape[])
AS $$
DECLARE
  repo_count INT;
  _td TypeDataShape;
BEGIN
  SELECT COUNT(*) INTO repo_count FROM Repositories R WHERE R.username ILIKE $1 AND R.repo ILIKE $2;

  IF repo_count < 1 THEN
    -- Create repository, if don't exist.
    CALL add_repo($1, $2, $3);
  END IF;

  DELETE FROM TypeData WHERE username ILIKE $1 AND repo ILIKE $2;

  FOREACH _td IN ARRAY $4 LOOP
    INSERT INTO TypeData VALUES (LOWER($1), LOWER($2), LOWER(_td.language), _td.file_count, _td.bytes);
  END LOOP;

  -- Update `last_updated` field.
  UPDATE Repositories
    SET last_updated = NOW()
    WHERE username ILIKE $1
      AND repo ILIKE $2;
END; $$ LANGUAGE plpgsql;

/*
--- FOR TESTING ---

CALL update_repos('__user', array[('__user', '__repo1', NULL, 'main'), ('__user', '__repo2', NULL, 'master')]::Repositories[]);
SELECT * FROM Repositories; -- should have [1, 2]
CALL update_repos('__user', array[('__user', '__repo3', NULL, 'main')]::Repositories[]);
SELECT * FROM Repositories; -- should have [3]
CALL update_repos('__user', array[('__user', '__repo3', NULL, 'main'), ('__user', '__repo4', NULL, 'main')]::Repositories[]);
SELECT * FROM Repositories; -- should have [3, 4]

CALL update_typedata('__user', '__repo', 'main', array[('java', 1, 420), ('go', 2, 34)]::TypeDataShape[]);
SELECT * FROM TypeData; -- should have [(java, 1, 420), (go, 2, 34)]
CALL update_typedata('__user', '__repo', 'main', array[('java', 2, 200), ('py', 111, 3)]::TypeDataShape[]);
SELECT * FROM TypeData; -- should have [(java, 2, 200), (py, 111, 3)]

DELETE FROM Users WHERE username = '__user';
*/
