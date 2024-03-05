CREATE TABLE post (
  id SERIAL PRIMARY KEY,
  uuid VARCHAR UNIQUE NOT NULL,
  blog_uuid VARCHAR UNIQUE NOT NULL,
  published TIMESTAMP,
  updated TIMESTAMP,
  url VARCHAR UNIQUE NOT NULL,
  self_link VARCHAR,
  title VARCHAR,
  title_link VARCHAR,
  content VARCHAR,
  custom_meta_data VARCHAR,
  status VARCHAR
);

CREATE TABLE location (
  id SERIAL PRIMARY KEY,
  uuid VARCHAR UNIQUE NOT NULL,
  post_uuid VARCHAR UNIQUE NOT NULL,
  name VARCHAR,
  lat FLOAT(8),
  lng FLOAT(8),
  span VARCHAR
);

CREATE TABLE label (
  id SERIAL PRIMARY KEY,
  uuid VARCHAR UNIQUE NOT NULL,
  post_uuid VARCHAR UNIQUE NOT NULL,
  label_value VARCHAR
);

CREATE TABLE image (
  id SERIAL PRIMARY KEY,
  uuid VARCHAR UNIQUE NOT NULL,
  post_uuid VARCHAR UNIQUE NOT NULL,
  author_uuid VARCHAR UNIQUE NOT NULL,
  url VARCHAR
);

CREATE TABLE post_user_info (
  id SERIAL PRIMARY KEY,
  uuid VARCHAR UNIQUE NOT NULL,
  user_uuid VARCHAR UNIQUE NOT NULL,
  blog_uuid VARCHAR UNIQUE NOT NULL,
  post_uuid VARCHAR UNIQUE NOT NULL,
  has_edit_access BOOLEAN
);

CREATE TABLE author (
  id SERIAL PRIMARY KEY,
  uuid VARCHAR UNIQUE NOT NULL,
  post_uuid VARCHAR UNIQUE NOT NULL,
  page_uuid VARCHAR UNIQUE NOT NULL,
  comment_uuid VARCHAR UNIQUE NOT NULL,
  display_name VARCHAR,
  url VARCHAR
);