CREATE TABLE `users`
(
    `id`                 INT unsigned NOT NULL AUTO_INCREMENT,
    `first_name`         VARCHAR(255) NOT NULL DEFAULT '',
    `last_name`          VARCHAR(255) NOT NULL DEFAULT '',
    `profile_image`      VARCHAR(255) NOT NULL DEFAULT '',
    `email`              VARCHAR(120) NOT NULL,
    `is_email_verified`  TINYINT(1)   NULL     DEFAULT 0,
    `bio`                TEXT         NOT NULL,
    `phone_number`       VARCHAR(20)  NOT NULL DEFAULT '',
    `is_phone_verified`  TINYINT(1)   NULL     DEFAULT 0,
    `country`            VARCHAR(255) NOT NULL DEFAULT '',
    `state`              VARCHAR(255) NOT NULL DEFAULT '',
    `area`               VARCHAR(255) NOT NULL DEFAULT '',
    `city`               VARCHAR(255) NOT NULL DEFAULT '',
    `address`            VARCHAR(255) NOT NULL DEFAULT '',
    `post_code`          VARCHAR(50)  NOT NULL DEFAULT '',
    `birth_date`         TIMESTAMP    NULL,
    `invited_by_user_id` INT unsigned NULL,
    `tos_accepted`       TINYINT(1)   NULL     DEFAULT 0,
    `is_active`          TINYINT(1)   NULL     DEFAULT 0,
    `created_at`         TIMESTAMP             DEFAULT CURRENT_TIMESTAMP,
    `updated_at`         TIMESTAMP             DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at`         TIMESTAMP    NULL,
    PRIMARY KEY (`id`),
    UNIQUE INDEX `users_email_idx` (`email` ASC),
    INDEX `fk_invited_by_user_id_idx` (`invited_by_user_id` ASC),
    CONSTRAINT `fk_invited_by_user_id`
        FOREIGN KEY (`invited_by_user_id`)
            REFERENCES `users` (`id`)
            ON DELETE SET NULL
            ON UPDATE CASCADE
) ENGINE = InnoDB;
