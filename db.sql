CREATE DATABASE db_eatnfit_auth;
USE db_eatnfit_auth;

CREATE TABLE tb_user (
    user_id VARCHAR(36) PRIMARY KEY NOT NULL,
    user_email VARCHAR(100) NOT NULL,
    user_password VARCHAR(255) NOT NULL,
    user_name VARCHAR(100) NOT NULL,
    user_photo VARCHAR(255) NOT NULL,
    user_balance INT NULL DEFAULT 0,
    user_level INT NOT NULL,
    user_status INT NOT NULL
);

CREATE DATABASE db_eatnfit_food;
USE db_eatnfit_food;

CREATE TABLE tb_food (
     food_id VARCHAR(36) PRIMARY KEY NOT NULL,
     food_portion INT NULL DEFAULT 0,
     food_name VARCHAR(100) NOT NULL,
     food_calories INT NULL DEFAULT 0,
     food_fat INT NULL DEFAULT 0,
     food_carbs INT NULL DEFAULT 0,
     food_protein INT NULL DEFAULT 0,
     food_price INT NOT NULL,
     food_desc TEXT NULL,
     food_status INT NOT NULL
);

CREATE TABLE tb_menu (
    menu_id VARCHAR(36) PRIMARY KEY NOT NULL,
    menu_portion INT NULL DEFAULT 0,
    menu_name VARCHAR(100) NOT NULL,
    menu_desc TEXT NULL,
    menu_status INT NOT NULL
);

CREATE TABLE tb_packet (
    packet_id VARCHAR(36) PRIMARY KEY NOT NULL,
    packet_name VARCHAR(100) NOT NULL,
    packet_price INT NOT NULL,
    packet_desc TEXT NOT NULL,
    packet_status INT NOT NULL
);

CREATE TABLE tb_menu_and_food (
    mf_id VARCHAR(36) PRIMARY KEY NOT NULL,
    menu_id VARCHAR(36) NOT NULL,
    food_id VARCHAR(36) NOT NULL,
    mf_status INT NOT NULL
);

CREATE TABLE tb_packet_and_menu (
    pd_id VARCHAR(36) PRIMARY KEY NOT NULL,
    packet_id VARCHAR(36) NOT NULL,
    menu_id VARCHAR(36) NOT NULL,
    pd_status INT NOT NULL
);

CREATE TABLE tb_transaction (
    trans_id VARCHAR(36) PRIMARY KEY NOT NULL,
    trans_date DATE NOT NULL,
    user_id VARCHAR(36) NOT NULL,
    packet_id VARCHAR(36) NOT NULL,
    portion INT NOT NULL,
    start_date DATE NOT NULL,
    end_date DATE NOT NULL,
    start_time TIMESTAMP NOT NULL,
    address TEXT NOT NULL,
    payment_id VARCHAR(36) NOT NULL
);

CREATE TABLE tb_payment (
    payment_id VARCHAR(36) PRIMARY KEY NOT NULL,
    payment_name VARCHAR(100) NOT NULL,
    payment_status INT NOT NULL
);
