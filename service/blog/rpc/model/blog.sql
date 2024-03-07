CREATE TABLE blog (
  id SERIAL PRIMARY KEY,
  uuid VARCHAR UNIQUE NOT NULL,
  app_user_uuid VARCHAR,
  name VARCHAR,
  description VARCHAR,
  published TIMESTAMP,
  updated TIMESTAMP,
  url VARCHAR,
  self_link VARCHAR,
  custom_meta_data VARCHAR
);

CREATE TABLE blog_user_info (
  id SERIAL PRIMARY KEY,
  uuid VARCHAR UNIQUE NOT NULL,
  user_uuid VARCHAR,
  blog_uuid VARCHAR,
  photos_album_key VARCHAR,
  has_admin_access BOOLEAN
);

CREATE TABLE page_views (
  id SERIAL PRIMARY KEY,
  uuid VARCHAR UNIQUE NOT NULL,
  blog_uuid VARCHAR UNIQUE NOT NULL,
  count INT
);