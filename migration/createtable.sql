CREATE TABLE IF NOT EXISTS `users` (
  `id` int NOT NULL AUTO_INCREMENT,
  `id_generate` varchar(150),
  `role` varchar(50) NOT NULL,
  `fullname` varchar(150) NOT NULL,
  `phone_number` varchar(50) NOT NULL,
  `username` varchar(50) NOT NULL UNIQUE,
  `password` varchar(50) NOT NULL,
  `parent_id` int,
  `position` set('left', 'center', 'right'),
  `sas_balance` int NOT NULL DEFAULT 0,
  `ro_balance` int NOT NULL DEFAULT 0,
  `money_balance` int NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`)
);

CREATE TABLE IF NOT EXISTS `bank_accounts` (
  `id` int NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `bank_name` varchar(50) NOT NULL,
  `bank_number` varchar(50) NOT NULL,
  `name_on_bank` varchar(50) NOT NULL,
  `user_id` int NOT NULL,
  FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS `transactions` (
  `id` int NOT NULL AUTO_INCREMENT,
  `from_id` int NOT NULL,
  `to_id` int NOT NULL,
  `sas_balance` int,
  `ro_balance` int,
  `money_balance` int,
  PRIMARY KEY (`id`)
);

CREATE TABLE IF NOT EXISTS `withdraw_requests` (
  `id` int NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `user_id` int NOT NULL,
  `bank_acc_id` int NOT NULL,
  `money_balance` int,
  `ro_balance` int,
  `ro_money_balance` int,
  `created_at` DATETIME NOT NULL,
  `updated_at` DATETIME NOT NULL,
  `approved` BOOLEAN NOT NULL,
  FOREIGN KEY (user_id) REFERENCES users(id),
  FOREIGN KEY (bank_acc_id) REFERENCES bank_accounts(id)
);