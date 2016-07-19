CREATE EXTENSION IF NOT EXISTS postgis;
CREATE TABLE games (
  id text NOT NULL DEFAULT encode( gen_random_bytes( 32 ), 'hex' ) primary key,
  name text NOT NULL,
  createdAt timestamp NOT NULL DEFAULT now()
);
CREATE UNIQUE INDEX gamename_idx ON games (name);

CREATE TABLE gameStages (
  id text NOT NULL DEFAULT encode( gen_random_bytes( 32 ), 'hex' ) primary key,
  gameId text NOT NULL REFERENCES games (id) ON DELETE CASCADE,
  number integer,
  name text NOT NULL,
  url text NOT NULL,
  createdAt timestamp NOT NULL DEFAULT now()
);
CREATE INDEX gamestages_gameid_idx ON gameStages (gameId);
CREATE INDEX gamestages_name_idx ON gameStages (name);
CREATE INDEX gamestages_url_idx ON gameStages (url);

CREATE TABLE gameProgresses (
  id text NOT NULL DEFAULT encode( gen_random_bytes( 32 ), 'hex' ) primary key,
  gameId text NOT NULL REFERENCES games (id) ON DELETE CASCADE,
  userId text NOT NULL REFERENCES users (id) ON DELETE CASCADE,
  event text NOT NULL,
  createdAt timestamp NOT NULL DEFAULT now()
);
CREATE INDEX progresses_game_idx ON gameProgresses (gameId);
CREATE INDEX progresses_user_idx ON gameProgresses (userId);
CREATE INDEX progresses_event_idx ON gameProgresses (event);

