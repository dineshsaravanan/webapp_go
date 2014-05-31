--liquibase formatted sql
--changeset dinesh:1
CREATE DATABASE snowball_im;

use snowball_im;

CREATE TABLE users (
    id MEDIUMINT NOT NULL AUTO_INCREMENT,
    name CHAR(30) NOT NULL,
    password CHAR(128),
    PRIMARY KEY (id)
);

CREATE TABLE authorizations (
    network VARCHAR(255) NOT NULL,
    network_id VARCHAR(255) NOT NULL,
    user_id MEDIUMINT NOT NULL,
    PRIMARY KEY (network, network_id),
    FOREIGN KEY (user_id) REFERENCES users(id)
);



