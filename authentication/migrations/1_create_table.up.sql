CREATE TABLE users (
  id              UUID NOT NULL PRIMARY KEY,
  firstname       VARCHAR(255) NOT NULL,
  lastname        VARCHAR(255) NOT NULL,
  othernames      VARCHAR(255) DEFAULT NULL,
  username        VARCHAR(255) NOT NULL UNIQUE,
  email           VARCHAR(255) NOT NULL UNIQUE,
  date_of_birth   DATE NOT NULL,
  password        TEXT NOT NULL,
  phone           VARCHAR(255) NOT NULL,
  address         VARCHAR(255) DEFAULT NULL,
  city            VARCHAR(255) NOT NULL,
  country         VARCHAR(255) NOT NULL,
  token           VARCHAR(4096) NOT NULL,
  refresh_token   VARCHAR(4096) NOT NULL,
  role            TEXT NOT NULL DEFAULT 'user',
  created_at      TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at      TIMESTAMP NOT NULL DEFAULT NOW()
);