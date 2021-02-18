CREATE DATABASE go_app;
USE go_app;
CREATE TABLE roles (
  id		BIGINT PRIMARY KEY,
  name		VARCHAR(255),
  password	VARCHAR(255),
  reg_date	DATETIME
);
INSERT INTO roles VALUES(1, 'admin', 'admin', now());
