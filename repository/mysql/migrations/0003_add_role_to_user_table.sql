-- +migrate Up

-- TODO find a better solution instead of keeping the order!!!
ALTER TABLE `users` ADD COLUMN `role` ENUM('user', 'admin') NOT NULL;

-- +migrate Down
ALTER TABLE `users` DROP COLUMN `role`;
