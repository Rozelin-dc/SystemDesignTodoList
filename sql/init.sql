DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `user_id` VARCHAR(32) NOT NULL,
  `user_name` VARCHAR(32) NOT NULL,
  `password` VARCHAR(86) DEFAULT NULL,
  PRIMARY KEY (`user_id`),
  UNIQUE KEY `user_name` (`user_name`)
);

-- status 0: 未完, 1: 完了
DROP TABLE IF EXISTS `tasks`;
CREATE TABLE `tasks` (
  `task_id` VARCHAR(32) NOT NULL,
  `creator_id` VARCHAR(32) NOT NULL,
  `task_name` VARCHAR(32) NOT NULL,
  `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `status` SMALLINT DEFAULT 0,
  `time_limit` TIMESTAMP DEFAULT NULL,
  PRIMARY KEY (`task_id`)
);

DROP TABLE IF EXISTS `sessions`;
CREATE TABLE `sessions` (
  `session_id` VARCHAR(32) NOT NULL,
  `user_id` VARCHAR(32) NOT NULL,
  PRIMARY KEY (`session_id`)
);
