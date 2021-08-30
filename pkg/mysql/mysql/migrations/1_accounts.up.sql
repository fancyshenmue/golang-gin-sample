CREATE TABLE `sre_app`.`accounts` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `uid` VARCHAR(11) CHARACTER SET 'utf8mb4' COLLATE 'utf8mb4_general_ci' NOT NULL,
  `firstname` VARCHAR(45) CHARACTER SET 'utf8mb4' COLLATE 'utf8mb4_general_ci' NOT NULL,
  `lastname` VARCHAR(45) CHARACTER SET 'utf8mb4' COLLATE 'utf8mb4_general_ci' NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `id_UNIQUE` (`id` ASC),
  UNIQUE INDEX `uid_UNIQUE` (`uid` ASC),
  INDEX `firstname_INDEX` (`firstname` ASC),
  INDEX `lastname_INDEX` (`lastname` ASC));
