CREATE TABLE games
  (
     game_id     serial PRIMARY KEY,
     state       int NOT NULL,
     columns     SMALLINT NOT NULL,
     rows        SMALLINT NOT NULL,
     mine_amount SMALLINT NOT NULL,
     flag_amount SMALLINT NOT NULL,
     board       JSONB NOT NULL
  );