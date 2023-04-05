-- +goose Up
CREATE TABLE `users` (
  `id` varchar(36) NOT NULL,
  `username` varchar(255) DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  `created_at` int NOT NULL,
  `updated_at` int NOT NULL,
  PRIMARY KEY (`id`)
);

CREATE TABLE `items` (
  `id` varchar(36) NOT NULL,
  `name` varchar(255) DEFAULT NULL,
  `image` varchar(255) DEFAULT NULL,
  `rating` int DEFAULT NULL,
  `reputation_value` int DEFAULT NULL,
  `price` int DEFAULT NULL,
  `availibility` int DEFAULT NULL,
  `category` enum('photo','sketch','cartoon','animation') DEFAULT NULL,
  `creator_id` varchar(36) DEFAULT NULL,
  `created_at` int DEFAULT NULL,
  `updated_at` int DEFAULT NULL,
  `reputation_badge` varchar(20) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `creator_id` (`creator_id`),
  CONSTRAINT `items_ibfk_1` FOREIGN KEY (`creator_id`) REFERENCES `users` (`id`)
);

CREATE TABLE `transactions` (
  `id` varchar(36) NOT NULL,
  `item_id` varchar(36) DEFAULT NULL,
  `buyer_id` varchar(36) DEFAULT NULL,
  `created_at` int DEFAULT NULL,
  `updated_at` int DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `item_id` (`item_id`),
  KEY `buyer_id` (`buyer_id`),
  CONSTRAINT `transactions_ibfk_1` FOREIGN KEY (`item_id`) REFERENCES `items` (`id`),
  CONSTRAINT `transactions_ibfk_2` FOREIGN KEY (`buyer_id`) REFERENCES `users` (`id`)
);



INSERT INTO `users` (`id`, `username`, `email`, `created_at`, `updated_at`) VALUES
('9bbbd26e-4738-4539-8ab4-e91f085231f5', 'test', 'test@gmail.com', 1680619396, 1680619396);

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
