CREATE DATABASE problem3;

USE `problem3`;

CREATE TABLE users (
    id INT NOT NULL AUTO INCREMENT PRIMARY KEY
    name VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
