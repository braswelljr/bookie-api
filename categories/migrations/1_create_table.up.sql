CREATE TABLE categories (
  id              UUID NOT NULL PRIMARY KEY,
  uid             UUID NOT NULL,
  name            TEXT NOT NULL,
  description     TEXT NOT NULL,
  color           VARCHAR(4092),
  created_at      TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at      TIMESTAMP NOT NULL DEFAULT NOW()
);