USE UserDb;

CREATE TABLE IF NOT EXISTS users(
    `id` varchar(36) NOT NULL,
    `name` varchar(100) NOT NULL,
    `type` integer NOT NULL,
    `active` boolean NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARACTER SET = UTF8MB4
  COLLATE = utf8mb4_unicode_520_ci;

CREATE TABLE IF NOT EXISTS `address`(
    `id` varchar(36) NOT NULL,
    `user_id` varchar(36) NOT NULL,
    `address_complement` varchar(100) NOT NULL,
    `address_number` varchar(10) NOT NULL,
    `address_street` varchar(50) NOT NULL,
    `address_neighborhood` varchar(50) NOT NULL,
    `address_state` varchar(50) NOT NULL,
    `address_country` varchar(50) NOT NULL,
    `address_code` varchar(30) NOT NULL,
    `address_latitude` varchar(30) NOT NULL,
    `address_longitude` varchar(30) NOT NULL,

    CONSTRAINT `fk_address_user_id` FOREIGN KEY (`user_id`) REFERENCES users (`id`) ON DELETE CASCADE
) ENGINE = InnoDB
  DEFAULT CHARACTER SET = UTF8MB4
  COLLATE = utf8mb4_unicode_520_ci;

CREATE TABLE IF NOT EXISTS `user_type_domain`(
    `id` integer NOT NULL,
    `name` varchar(20) NOT NULL,

    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARACTER SET = UTF8MB4
  COLLATE = utf8mb4_unicode_520_ci;


CREATE TABLE IF NOT EXISTS `user_type`(
    `user_type_id` integer NOT NULL,
    `user_id` varchar(36) NOT NULL,

    CONSTRAINT `fk_type_user_id` FOREIGN KEY (`user_id`) REFERENCES users (`id`) ON DELETE CASCADE,
    CONSTRAINT `fk_type_user_type_id` FOREIGN KEY (`user_type_id`) REFERENCES `user_type_domain` (`id`) ON DELETE CASCADE
) ENGINE = InnoDB
  DEFAULT CHARACTER SET = UTF8MB4
  COLLATE = utf8mb4_unicode_520_ci;
