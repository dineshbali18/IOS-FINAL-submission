create table user_carts(
    `user_id` int unsigned not null,
    `product_id` int unsigned not null,
    `quantity` int unsigned not null,
    FOREIGN KEY(`user_id`) REFERENCES users(`id`),
    FOREIGN KEY(`product_id`) REFERENCES products(`id`)
)ENGINE=InnoDB;