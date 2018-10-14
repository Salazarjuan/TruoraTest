DROP TABLE step;
DROP TABLE recipe;
DROP TABLE users;

CREATE TABLE users(
  id int NOT NULL AUTO_INCREMENT,
  first varchar(255) NOT NULL,
  last varchar(255) NOT NULL,
  email varchar(255) NOT NULL,
  password varchar(255) NOT NULL,
  PRIMARY KEY (id)
);

CREATE TABLE recipe (
  id int NOT NULL AUTO_INCREMENT,
  name varchar(255) NOT NULL,
  description varchar(255) NOT NULL,
  oven varchar(255) NOT NULL,
  time int NOT NULL,
  noPersons int NOT NULL,
  creatorID int NOT NULL,
  PRIMARY KEY (id),
  FOREIGN KEY fk_users(creatorID) REFERENCES users(id) ON UPDATE RESTRICT ON DELETE RESTRICT
);

CREATE TABLE step(
  id int NOT NULL AUTO_INCREMENT,
  step varchar(255) NOT NULL,
  recipeID int NOT NULL,
  PRIMARY KEY (id),
  FOREIGN KEY fk_recipe (recipeID) REFERENCES recipe (id) ON UPDATE RESTRICT ON DELETE RESTRICT
);

INSERT INTO users (first, last, email, password) VALUES('jose','juan','jose@mail.com','1234');

INSERT INTO users (first, last, email, password) VALUES('juan','jose','juan@mail.com','1234');