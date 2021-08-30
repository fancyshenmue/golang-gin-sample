CREATE TABLE `sre_app`.`groups` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `gid` VARCHAR(11) CHARACTER SET 'utf8mb4' COLLATE 'utf8mb4_general_ci' NOT NULL,
  `groupname` VARCHAR(45) CHARACTER SET 'utf8mb4' COLLATE 'utf8mb4_general_ci' NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `id_UNIQUE` (`id` ASC),
  UNIQUE INDEX `gid_UNIQUE` (`gid` ASC),
  UNIQUE INDEX `groupname_UNIQUE` (`groupname` ASC));
