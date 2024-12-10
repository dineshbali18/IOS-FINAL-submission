create table users(
    `id` int unsigned not null AUTO_INCREMENT,
    `phone_number` varchar(13) not null,
    `name` varchar(20) not null,
    `email` varchar(40) not null, 
    `password_hash` varchar(100) not null,
    `isVerified` boolean not null DEFAULT 0,
    `role` varchar(20) not null,
    PRIMARY KEY(`id`)
)ENGINE=InnoDB;