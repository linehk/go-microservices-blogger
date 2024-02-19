CREATE TABLE comment (
  id SERIAL PRIMARY KEY,
  uuid VARCHAR UNIQUE NOT NULL,
  blog_uuid VARCHAR UNIQUE NOT NULL,
  post_uuid VARCHAR UNIQUE NOT NULL,
  status VARCHAR,
  published TIMESTAMP,
  updated TIMESTAMP,
  selfLink VARCHAR,
  content VARCHAR
);