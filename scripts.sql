CREATE TABLE games
  (
     game_id     INT PRIMARY KEY NOT NULL,
     state       TEXT NOT NULL,
     columns     SMALLINT NOT NULL,
     rows        SMALLINT NOT NULL,
     mine_amount SMALLINT NOT NULL,
     flag_amount SMALLINT NOT NULL,
     board       JSON NOT NULL
  );