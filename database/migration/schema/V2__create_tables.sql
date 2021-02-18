CREATE TABLE IF NOT EXISTS `result_movement_direction_discrimination_constant_width`
(
    `id`              INT PRIMARY KEY AUTO_INCREMENT,
    `experiment_id`   INT       NOT NULL,
    `result_url`      TEXT      NOT NULL,
    `subject_id`      INT       NOT NULL,
    `mean`            DOUBLE    NOT NULL,
    `sd`              DOUBLE    NOT NULL,
    `lower_asymptote` DOUBLE    NOT NULL,
    `lapse_rate`      DOUBLE    NOT NULL,
    `created_at`      TIMESTAMP NOT NULL DEFAULT (now())
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `experiment`
(
    `id`         INT PRIMARY KEY AUTO_INCREMENT,
    `created_at` TIMESTAMP NOT NULL DEFAULT (now())
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `experiment_active`
(
    `id`            INT PRIMARY KEY AUTO_INCREMENT,
    `experiment_id` INT UNIQUE NOT NULL,
    `created_at`    TIMESTAMP  NOT NULL DEFAULT (now())
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `experiment_inactive`
(
    `id`            INT PRIMARY KEY AUTO_INCREMENT,
    `experiment_id` INT UNIQUE NOT NULL,
    `created_at`    TIMESTAMP  NOT NULL DEFAULT (now())
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `experiment_movement_direction_discrimination_constant_width`
(
    `id`                             INT PRIMARY KEY AUTO_INCREMENT,
    `experiment_id`                  INT UNIQUE         NOT NULL,
    `questplus_parameter_normcdf_id` INT                NOT NULL,
    `name`                           varchar(64) UNIQUE NOT NULL,
    `description`                    varchar(1024)      NOT NULL,
    `azimuth`                        INT UNSIGNED       NOT NULL COMMENT '0[10^-1 deg] is the front, 900[10^-1 deg] is the right side',
    `altitude`                       INT UNSIGNED       NOT NULL COMMENT '0[10^-1 deg] is the front, 900[10^-1 deg] is the zenith, -900[10^-1 deg] is the nadir',
    `coordinate_variable_id`         INT                NOT NULL,
    `width`                          INT UNSIGNED       NOT NULL COMMENT '[10^-1 deg]',
    `velocity_range_lower`           INT UNSIGNED       NOT NULL COMMENT '[10^-1 deg/sec]',
    `velocity_range_upper`           INT UNSIGNED       NOT NULL COMMENT '[10^-1 deg/sec]',
    `velocity_range_step`            INT UNSIGNED       NOT NULL COMMENT '[10^-1 deg/sec]',
    `num_trials`                     INT UNSIGNED       NOT NULL,
    `created_at`                     TIMESTAMP          NOT NULL DEFAULT (now())
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `questplus_parameter_normcdf`
(
    `id`                           INT PRIMARY KEY AUTO_INCREMENT,
    `questplus_parameter_json_url` TEXT      NOT NULL,
    `created_at`                   TIMESTAMP NOT NULL DEFAULT (now())
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `m_coordinate_variable`
(
    `id`         INT PRIMARY KEY,
    `type`       VARCHAR(64) NOT NULL COMMENT 'azimuth or altitude',
    `created_at` TIMESTAMP   NOT NULL DEFAULT (now())
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `subject`
(
    `id`                        INT PRIMARY KEY,
    `sex`                       CHAR(1)      NOT NULL DEFAULT '0' COMMENT 'ISO5218',
    `age`                       INT UNSIGNED NOT NULL,
    `deaf_and_hearing_impaired` BOOL         NOT NULL,
    `created_at`                TIMESTAMP    NOT NULL DEFAULT (now())
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

ALTER TABLE `result_movement_direction_discrimination_constant_width`
    ADD FOREIGN KEY (`experiment_id`) REFERENCES `experiment` (`id`);

ALTER TABLE `result_movement_direction_discrimination_constant_width`
    ADD FOREIGN KEY (`subject_id`) REFERENCES `subject` (`id`);

ALTER TABLE `experiment_active`
    ADD FOREIGN KEY (`experiment_id`) REFERENCES `experiment` (`id`);

ALTER TABLE `experiment_inactive`
    ADD FOREIGN KEY (`experiment_id`) REFERENCES `experiment` (`id`);

ALTER TABLE `experiment_movement_direction_discrimination_constant_width`
    ADD FOREIGN KEY (`experiment_id`) REFERENCES `experiment` (`id`);

ALTER TABLE `experiment_movement_direction_discrimination_constant_width`
    ADD FOREIGN KEY (`questplus_parameter_normcdf_id`) REFERENCES `questplus_parameter_normcdf` (`id`);

ALTER TABLE `experiment_movement_direction_discrimination_constant_width`
    ADD FOREIGN KEY (`coordinate_variable_id`) REFERENCES `m_coordinate_variable` (`id`);
