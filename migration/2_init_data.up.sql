INSERT INTO `middleware`.`rdev_user`(`created_at`,
                             `updated_at`,
                             `id`,
                             `name`,
                             `password`,
                             `is_active`,
                             `is_admin`,
                             `type`,
                             `role`)
VALUES (date_add(now(), interval 8 HOUR),
        date_add(now(), interval 8 HOUR),
        '5e81095f-3c0c-4cb2-8033-bde03d60135c',
        'admin',
        'Pb4BAQEBAQHH2XmOtIOlsIViX4E8vnrkYHoWEi6lUFo=',
        1,
        1,
        'LOCAL',
        0);