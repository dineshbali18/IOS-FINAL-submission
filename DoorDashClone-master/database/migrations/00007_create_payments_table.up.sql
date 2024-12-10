create table PAYMENTS(
    `id` int unsigned not null, 
    `user_id` int unsigned not null,
    `amount` int unsigned not null,
    `order_id` int unsigned not null,
    `transaction_id` varchar(100) not null,
    PRIMARY KEY(id),
    FOREIGN KEY(`user_id`) References users(`id`),
    FOREIGN KEY(`order_id`) References user_orders(`id`)
)ENGINE=InnoDB;