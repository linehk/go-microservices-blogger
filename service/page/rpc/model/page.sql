CREATE TABLE page (
  id SERIAL PRIMARY KEY,
  uuid VARCHAR UNIQUE NOT NULL,
  blog_uuid VARCHAR,
  status VARCHAR,
  published TIMESTAMP,
  updated TIMESTAMP,
  url VARCHAR,
  self_link VARCHAR,
  title VARCHAR,
  content VARCHAR
);