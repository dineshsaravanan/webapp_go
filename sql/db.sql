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

--changeset dinesh:2
CREATE TABLE oauthkeys (
    network VARCHAR(255) NOT NULL,
    client_id VARCHAR(255) NOT NULL,
    client_secret VARCHAR(255) NOT NULL,
    email_address VARCHAR(255) NOT NULL,
    PRIMARY KEY (network)
);

insert into oauthkeys VALUES ("TEST NETWORK", "TEST CLIENT ID", "TEST SECRET", "TEST EMAIL");

insert into oauthkeys VALUES ("google", "google app id", "google secret", "google account mail id");
insert into oauthkeys VALUES ("facebook", "facebook app id", "facebook secret", "facebook account mail id");

