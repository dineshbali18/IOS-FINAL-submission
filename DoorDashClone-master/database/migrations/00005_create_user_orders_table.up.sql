create table `user_orders` ( 
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `user_id` INT UNSIGNED NOT NULL,
    `phone_number` VARCHAR(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'Represents phone number of the user',
    `is_delivered` BOOLEAN NOT NULL DEFAULT 0,
    `order_total` INT UNSIGNED NOT NULL DEFAULT 0,
    `drive_thru_code` INT UNSIGNED DEFAULT 1 COMMENT 'Represents drive thru code', 
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Timestamp at which the row was created',
    `order_status` VARCHAR(40) NOT NULL DEFAULT 'pending' COMMENT 'Represents the status of the order (e.g., cancelled, pending)',
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Timestamp at which the row has been updated',
    PRIMARY KEY (`id`),
    FOREIGN KEY (`user_id`) REFERENCES users(`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT ='Table to store orders data';

  
-- Add primary key and unique index to deviceID
ALTER TABLE `user_orders`
ADD KEY `created_at`(created_at),
ADD KEY `phone_number`(phone_number);
