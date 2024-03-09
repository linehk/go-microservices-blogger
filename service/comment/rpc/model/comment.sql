CREATE TABLE comment (
  id SERIAL PRIMARY KEY,
  uuid VARCHAR UNIQUE NOT NULL,
  blog_uuid VARCHAR,
  post_uuid VARCHAR,
  status VARCHAR,
  published TIMESTAMP,
  updated TIMESTAMP,
  self_link VARCHAR,
  content VARCHAR
);