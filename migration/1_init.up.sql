CREATE DATABASE IF NOT EXISTS `mw`;

use mw;

CREATE TABLE IF NOT EXISTS `rdev_user` (
    `created_at` datetime DEFAULT NULL,
    `updated_at` datetime DEFAULT NULL,
    `id` varchar(64) NOT NULL,
    `name` varchar(256) NOT NULL,
    `password` varchar(256) DEFAULT NULL,
    `is_active` tinyint(1) DEFAULT '1',
    `is_admin` tinyint(1) DEFAULT '1',
    `type` varchar(256) NOT NULL,
    `role` int NOT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `name` (`name`)
);

CREATE TABLE  IF NOT EXISTS  rdev_instance (
    id varchar(64) not null,
    name varchar(64) not null,
    type varchar(64) not null,
    cluster_id varchar(64) not null,
--     template_id varchar(64) not null,
    namespace varchar(64) not null,
    count         int         not null,
    request_cpu           varchar(64) not null,
    request_memory        varchar(64) not null,
    limit_cpu           varchar(64) not null,
    limit_memory        varchar(64) not null,
    volume        varchar(64) not null,
    status        varchar(64) not null,
    created_at                                  datetime     not null,
    updated_at                                  datetime     not null,
    PRIMARY KEY (`id`)
);

CREATE TABLE  IF NOT EXISTS  rdev_user_instance (
    user_id varchar(64) not null,
    instance_id varchar(64) not null,
    UNIQUE KEY `user_instance_key` (`user_id`, `instance_id`)
);