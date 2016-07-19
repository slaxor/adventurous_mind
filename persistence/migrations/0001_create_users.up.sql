CREATE EXTENSION IF NOT EXISTS pgcrypto;
CREATE TABLE users (
  id char(32) NOT NULL DEFAULT encode(gen_random_bytes(16),'hex') primary key,
  createdAt timestamp NOT NULL DEFAULT now()
);
CREATE INDEX users_createdat_idx ON users (createdAt);

CREATE TABLE userEvents (
  id char(32) NOT NULL DEFAULT encode(gen_random_bytes(16),'hex') primary key,
  userId char(32) NOT NULL,
  name text NOT NULL,
  oldValue text,
  newValue text NOT NULL,
  createdAt timestamp NOT NULL DEFAULT now()
);
CREATE INDEX user_events_userid_idx ON userEvents (userId);
CREATE INDEX user_events_name_idx ON userEvents (name);
CREATE INDEX user_events_createdat_idx ON userEvents (createdAt);


--  CREATE OR REPLACE FUNCTION auth(_name text, pass text, OUT _id char(32))
--    RETURNS char(32)
--    LANGUAGE plpgsql
--    AS $$
--    BEGIN
--      SELECT id INTO _id FROM users INNER JOIN userEvents AS ue ON
--        (ue.name = 'changedUsername' OR ue.name = 'changedEmail') AND
--        ue.newValue=lower(_name)
--          OR users.email=lower(_name)) AND userEvents.userId = _id AND
--          userEvents.newValue = crypt(pass, pwhash)
--          ORDER ue.createdAt DESC LIMIT 1;
--    END;
--  $$;

CREATE OR REPLACE FUNCTION create_user(username text, email text, pass text)
  RETURNS void

  LANGUAGE plpgsql
  AS $$
  DECLARE u char(32);
  BEGIN
    INSERT INTO users DEFAULT VALUES RETURNING id INTO u;
    INSERT INTO userEvents (userId, name, newValue) VALUES (u, 'changedUsername', lower(username));
    INSERT INTO userEvents (userId, name, newValue) VALUES (u, 'changedEmail', lower(email));
    INSERT INTO userEvents (userId, name, newValue) VALUES (u, 'changedPassword', crypt(pass, gen_salt('bf')));
  END;
$$;
SELECT create_user('slaxor', 'sascha.teske@gmail.com', 'ohmohS7chooWaiLa');
SELECT create_user('mazeon', 'sk@saschakurz.net', 'Xeequaelaiho2yef');

--  INSERT INTO users (name,email,pwhash) VALUES ('slaxor', 'sascha.teske@gmail.com', crypt('ohmohS7chooWaiLa', gen_salt('bf')));
--  INSERT INTO users (name,email,pwhash) VALUES ('mazeon', 'sk@saschakurz.net ', crypt('Xeequaelaiho2yef', gen_salt('bf')));
--- UPDATE users SET pwhash = crypt('new password', gen_salt('md5'));
--- SELECT pwhash = crypt('password', pwhash) FROM users WHERE name = 'slaxor';
--- SELECT pwhash = crypt('password', pwhash) FROM users;

