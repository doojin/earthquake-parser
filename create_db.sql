DROP DATABASE IF EXISTS earthquake_tracker;
CREATE DATABASE earthquake_tracker;

USE earthquake_tracker;

CREATE TABLE earthquakes (
  id INT NOT NULL AUTO_INCREMENT,
  start_date DATE NOT NULL,
  end_date DATE NOT NULL,
  magnitude DOUBLE NOT NULL,
  PRIMARY KEY(id)
) CHARACTER SET "utf8";
