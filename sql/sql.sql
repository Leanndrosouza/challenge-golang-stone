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
);

DROP TABLE IF EXISTS transfers;

CREATE TABLE transfers(
    id int auto_increment primary key,
    account_origin_id int not null REFERENCES accounts(id),
    account_destination_id int not null REFERENCES accounts(id),
    amount int not null,
    created_at timestamp default current_timestamp(),
    FOREIGN KEY (account_origin_id) REFERENCES accounts(id) ON DELETE CASCADE,
    FOREIGN KEY (account_destination_id) REFERENCES accounts(id) ON DELETE CASCADE
) ENGINE=INNODB