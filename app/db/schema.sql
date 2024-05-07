CREATE TABLE IF NOT EXISTS "schema_migrations" (version varchar(128) primary key);
CREATE TABLE authors (
  id INTEGER PRIMARY KEY,
  name text NOT NULL,
  bio text
);
-- Dbmate schema migrations
INSERT INTO "schema_migrations" (version) VALUES
  ('20240507030540');
