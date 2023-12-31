CREATE TABLE `candidate` (
  `id` integer PRIMARY KEY,
  `fullname` varchar(255),
  `dob` Date,
  `latitude` decimal,
  `longitude` decimal,
  `email` varchar(255),
  `mobile_phone` varchar(18),
  `password` varchar(255),
  `gender` ENUM ('male', 'female'),
  `city_id` integer,
  `province_id` integer,
  `last_education` varchar(3),
  `last_experience` integer,
  `login_date` timestamp,
  `created_at` timestamp,
  `updated_at` timestamp,
  `deleted_at` timestamp
);

CREATE TABLE `education` (
  `id` integer PRIMARY KEY,
  `candidate_id` integer,
  `institution_name` text,
  `major` text,
  `start_year` Date,
  `end_year` Date,
  `until_now` boolean,
  `gpa` float,
  `flag` boolean,
  `role` ENUM ('sd', 'smp', 'sma', 's1', 's2', 's3'),
  `created_at` timestamp,
  `updated_at` timestamp,
  `deleted_at` timestamp
);

CREATE TABLE `experience` (
  `id` integer PRIMARY KEY,
  `candidate_id` integer,
  `company_name` varchar(255),
  `company_address` text,
  `position` varchar(255),
  `job_desc` text,
  `start_year` Date,
  `end_year` Date,
  `until_now` boolean,
  `flag` boolean,
  `created_at` timestamp,
  `updated_at` timestamp,
  `deleted_at` timestamp
);

CREATE TABLE `city` (
  `id` integer PRIMARY KEY,
  `name` integer,
  `province_id` varchar(255),
  `created_at` timestamp,
  `updated_at` timestamp,
  `deleted_at` timestamp
);

CREATE TABLE `province` (
  `id` integer PRIMARY KEY,
  `name` integer,
  `created_at` timestamp,
  `updated_at` timestamp,
  `deleted_at` timestamp
);

ALTER TABLE `education` ADD FOREIGN KEY (`candidate_id`) REFERENCES `candidate` (`id`);

ALTER TABLE `experience` ADD FOREIGN KEY (`candidate_id`) REFERENCES `candidate` (`id`);

ALTER TABLE `candidate` ADD FOREIGN KEY (`city_id`) REFERENCES `city` (`id`);

ALTER TABLE `candidate` ADD FOREIGN KEY (`province_id`) REFERENCES `province` (`id`);

ALTER TABLE `province` ADD FOREIGN KEY (`id`) REFERENCES `city` (`province_id`);
