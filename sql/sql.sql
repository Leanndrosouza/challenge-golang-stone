CREATE DATABASE IF NOT EXISTS digital_bank;

USE digital_bank;

DROP TABLE IF EXISTS accounts;

CREATE TABLE accounts(
    id int auto_increment primary key,
    name varchar(50) not null,
    cpf varchar(11) not null unique,
    secret varchar(100) not null,
    balance int not null,
    created_at timestamp default current_timestamp()
) ENGINE=INNODB