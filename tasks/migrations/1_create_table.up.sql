CREATE TABLE tasks (
  id              UUID NOT NULL PRIMARY KEY,
  uid             UUID NOT NULL,
  title           TEXT NOT NULL,
  description     TEXT NOT NULL,
  status          VARCHAR(255) NOT NULL DEFAULT 'pending',
  category        VARCHAR(255) NOT NULL DEFAULT 'general',
  pinned          BOOLEAN NOT NULL DEFAULT FALSE,
  pinnedAt       TIMESTAMP NOT NULL DEFAULT NOW(),
  pinnedPosition INTEGER NOT NULL DEFAULT 0,
  archived        BOOLEAN NOT NULL DEFAULT FALSE,
  archivedAt     TIMESTAMP NOT NULL DEFAULT NOW(),
  completed       BOOLEAN NOT NULL DEFAULT FALSE,
  completedAt    TIMESTAMP NOT NULL DEFAULT NOW(),
  color           VARCHAR(4092),
  createdAt      TIMESTAMP NOT NULL DEFAULT NOW(),
  updatedAt      TIMESTAMP NOT NULL DEFAULT NOW()
);