CREATE TABLE IF NOT EXISTS rdev_templates
(
    `created_at` datetime     DEFAULT NULL,
    `updated_at` datetime     DEFAULT NULL,
    `id`    varchar(64) not null primary key,
    `name`  varchar(64) not null,
    `icon`  longblob not null,
    `base_template` longtext  not null,
    `advance_template` longtext not null
);