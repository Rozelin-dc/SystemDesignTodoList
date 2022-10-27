DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `user_id` VARCHAR(37) NOT NULL,
  `user_name` VARCHAR(30) NOT NULL,
  `password` VARCHAR(90) DEFAULT NULL,
  PRIMARY KEY (`user_id`),
  UNIQUE KEY `user_name` (`user_name`)
);

-- status 0: 未完, 1: 完了
DROP TABLE IF EXISTS `tasks`;
CREATE TABLE `tasks` (
  `task_id` VARCHAR(37) NOT NULL,
  `creator_id` VARCHAR(37) NOT NULL,
  `task_name` VARCHAR(60) NOT NULL,
  `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `status` SMALLINT DEFAULT 0,
  `time_limit` VARCHAR(10) DEFAULT NULL,
  PRIMARY KEY (`task_id`)
);

DROP TABLE IF EXISTS `sessions`;
CREATE TABLE `sessions` (
  `session_id` VARCHAR(37) NOT NULL,
  `user_id` VARCHAR(37) NOT NULL,
  PRIMARY KEY (`session_id`)
);
