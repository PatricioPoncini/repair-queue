CREATE TABLE IF NOT EXISTS appointment (
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `reason` VARCHAR(255) NOT NULL,
    `model` VARCHAR(255) NOT NULL,
    `make` VARCHAR(255) NOT NULL,
    `licencePlate` VARCHAR(255) NOT NULL,
    `manufactureYear` INT NOT NULL,
    `status` VARCHAR(255) NOT NULL,
    `ownerPhoneNumber` VARCHAR(255) NOT NULL,
    `createdAt` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY(id)
)