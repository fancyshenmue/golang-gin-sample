CREATE TABLE `sre_app`.`members` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `uid` VARCHAR(11) CHARACTER SET 'utf8mb4' COLLATE 'utf8mb4_general_ci' NOT NULL,
  `gid` VARCHAR(11) CHARACTER SET 'utf8mb4' COLLATE 'utf8mb4_general_ci' NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `id_UNIQUE` (`id` ASC),
  INDEX `uid_INDEX` (`uid` ASC),
  INDEX `gid_INDEX` (`gid` ASC));
