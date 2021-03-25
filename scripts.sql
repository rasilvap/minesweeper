CREATE TABLE games
  (
     game_id     serial PRIMARY KEY,
     state       TEXT NOT NULL,
     columns     SMALLINT NOT NULL,
     rows        SMALLINT NOT NULL,
     mine_amount SMALLINT NOT NULL,
     flag_amount SMALLINT NOT NULL,
     board       JSON NOT NULL
  );