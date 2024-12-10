create table hotels(
    `id` int unsigned not null AUTO_INCREMENT,
    `name` varchar(30) not null , 
    `city` varchar(50) not null,
    `address` varchar(200) not null,
    `state` varchar(50) not null,
    PRIMARY KEY(`id`)
)ENGINE=InnoDB;