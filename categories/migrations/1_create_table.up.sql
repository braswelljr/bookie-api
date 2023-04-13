CREATE TABLE categories (
  id              UUID NOT NULL PRIMARY KEY,
  uid             UUID NOT NULL,
  name            TEXT NOT NULL,
  description     TEXT NOT NULL,
  color           VARCHAR(255) NOT NULL DEFAULT '#00b3ff',
  createdAt       TIMESTAMP NOT NULL DEFAULT NOW(),
  updatedAt       TIMESTAMP NOT NULL DEFAULT NOW()
);