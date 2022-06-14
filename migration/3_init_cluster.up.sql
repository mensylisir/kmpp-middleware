CREATE TABLE IF NOT EXISTS `rdev_cluster`
(
    `created_at` datetime     DEFAULT NULL,
    `updated_at` datetime     DEFAULT NULL,
    `id`         varchar(255) NOT NULL,
    `name`       varchar(255) NOT NULL,
    `api_server` varchar(255) NOT NULL,
    `version`    varchar(255) NOT NULL,
    `token`      mediumtext DEFAULT NULL,
    `type`       varchar(255) DEFAULT NULL,
    `status`     varchar(255) DEFAULT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `name` (`name`)
);

CREATE TABLE IF NOT EXISTS rdev_user_cluster
(
    user_id    varchar(64) not null,
    cluster_id varchar(64) not null,
    UNIQUE KEY `user_instance_key` (`user_id`, `cluster_id`)
);