CREATE TABLE `candidate` (
  `id` serial,
  `fullname` varchar(255),
  `dob` timestamp NOT NULL,
  `latitude` decimal,
  `longitude` decimal,
  `email` varchar(255) UNIQUE NOT NULL,
  `mobile_phone` varchar(18) UNIQUE NOT NULL,
  `password` varchar(255) NOT NULL,
  `gender` ENUM ('male', 'female') NOT NULL,
  `city_id` integer NOT NULL,
  `province_id` integer NOT NULL,
  `last_education` varchar(3),
  `last_experience` integer,
  `login_date` timestamp,
  `created_at` timestamp NOT NULL DEFAULT (now()),
  `updated_at` timestamp NOT NULL,
  `deleted_at` timestamp
);

CREATE TABLE `education` (
  `id` serial,
  `candidate_id` integer,
  `institution_name` text,
  `major` text,
  `start_year` Date,
  `end_year` Date,
  `until_now` boolean,
  `gpa` float,
  `flag` boolean,
  `role` ENUM ('sd', 'smp', 'sma', 's1', 's2', 's3'),
  `created_at` timestamp NOT NULL DEFAULT (now()),
  `updated_at` timestamp NOT NULL,
  `deleted_at` timestamp
);

CREATE TABLE `experience` (
  `id` serial,
  `candidate_id` integer NOT NULL,
  `company_name` varchar(255) NOT NULL,
  `company_address` text NOT NULL,
  `position` varchar(255) NOT NULL,
  `job_desc` text NOT NULL,
  `start_year` Date NOT NULL,
  `end_year` Date NOT NULL,
  `until_now` boolean NOT NULL DEFAULT true,
  `flag` boolean,
  `created_at` timestamp NOT NULL DEFAULT (now()),
  `updated_at` timestamp NOT NULL,
  `deleted_at` timestamp
);

CREATE TABLE `city` (
  `id` serial,
  `name` varchar(255),
  `province_id` integer,
  `created_at` timestamp NOT NULL DEFAULT (now()),
  `updated_at` timestamp NOT NULL,
  `deleted_at` timestamp
);

CREATE TABLE `province` (
  `id` serial,
  `name` varchar(255),
  `created_at` timestamp NOT NULL DEFAULT (now()),
  `updated_at` timestamp NOT NULL,
  `deleted_at` timestamp
);

ALTER TABLE `education` ADD FOREIGN KEY (`candidate_id`) REFERENCES `candidate` (`id`);

ALTER TABLE `experience` ADD FOREIGN KEY (`candidate_id`) REFERENCES `candidate` (`id`);

ALTER TABLE `candidate` ADD FOREIGN KEY (`city_id`) REFERENCES `city` (`id`);

ALTER TABLE `candidate` ADD FOREIGN KEY (`province_id`) REFERENCES `province` (`id`);

ALTER TABLE `city` ADD FOREIGN KEY (`province_id`) REFERENCES `province` (`id`);
