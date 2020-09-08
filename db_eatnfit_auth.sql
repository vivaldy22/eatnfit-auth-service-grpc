CREATE DATABASE db_eatnfit_auth;
USE db_eatnfit_auth;

CREATE TABLE tb_user (
     user_id VARCHAR(36) PRIMARY KEY NOT NULL,
     user_email VARCHAR(100) NOT NULL,
     user_password VARCHAR(255) NOT NULL,
     user_f_name VARCHAR(100) NOT NULL,
     user_l_name VARCHAR(100) NOT NULL,
     user_gender INT NOT NULL,
     user_photo VARCHAR(255) NOT NULL,
     user_balance INT NULL DEFAULT 0,
     user_level INT NOT NULL,
     user_status INT NOT NULL
);

CREATE TABLE tb_gender (
    gender_id INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
    gender_name VARCHAR(50) NOT NULL,
    gender_status INT NOT NULL
);

CREATE TABLE tb_level (
    level_id INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
    level_name VARCHAR(50) NOT NULL,
    level_status INT NOT NULL
);

# CREATE TABLE tb_level_access (
#     level_access_id INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
#     level_id INT NOT NULL,
#     access_menu_id INT NOT NULL,
#     level_access_status INT NOT NULL
# );
#
# CREATE TABLE tb_access_menu (
#     access_menu_id INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
#     access_menu_name VARCHAR(50) NOT NULL,
#     access_menu_status INT NOT NULL
# );

INSERT INTO tb_gender VALUES (NULL, 'Laki-laki', 1),
                             (NULL, 'Perempuan', 1);

INSERT INTO tb_level VALUES (NULL, 'Admin', 1),
                            (NULL, 'User', 1),
                            (NULL, 'Driver', 1);
