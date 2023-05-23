CREATE TABLE tasks (
  id              UUID NOT NULL PRIMARY KEY,
  uid             UUID NOT NULL,
  title           TEXT NOT NULL,
  description     TEXT NOT NULL,
  status          VARCHAR(255) NOT NULL DEFAULT 'pending',
  category        VARCHAR(255) NOT NULL DEFAULT 'general',
  pinned          BOOLEAN NOT NULL DEFAULT FALSE,
  pinned_at        TIMESTAMP NOT NULL DEFAULT NOW(),
  pinned_position INTEGER NOT NULL DEFAULT 0,
  archived        BOOLEAN NOT NULL DEFAULT FALSE,
  archived_at     TIMESTAMP NOT NULL DEFAULT NOW(),
  completed       BOOLEAN NOT NULL DEFAULT FALSE,
  completed_at    TIMESTAMP NOT NULL DEFAULT NOW(),
  color           VARCHAR(4092),
  created_at      TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at      TIMESTAMP NOT NULL DEFAULT NOW()
);