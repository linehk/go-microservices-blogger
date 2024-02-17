CREATE TABLE app_user (
  id SERIAL PRIMARY KEY,
  uuid VARCHAR UNIQUE NOT NULL,
  created TIMESTAMP,
  url VARCHAR,
  self_link VARCHAR,
  display_name VARCHAR,
  about VARCHAR,
  language VARCHAR,
  country VARCHAR,
  variant VARCHAR
);