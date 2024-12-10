create table products(
    `id` int unsigned not null AUTO_INCREMENT,
    `name` varchar(15) not null,
    `stockLeft` int unsigned not null,
    `hotel_id` int unsigned not null,
    `category` varchar(20) not null,
    `price` int unsigned not null,
    PRIMARY KEY(`id`),
    FOREIGN KEY(`hotel_id`) References hotels(`id`)
)ENGINE= InnoDB