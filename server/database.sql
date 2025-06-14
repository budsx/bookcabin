CREATE TABLE `aircrafts` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `code` varchar(10),
  `created_at` datetime,
  `updated_at` datetime
);

CREATE TABLE `cabins` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `aircraft_id` int,
  `deck` varchar(10),
  `first_row` int,
  `last_row` int,
  `created_at` datetime,
  `updated_at` datetime
);

CREATE TABLE `seat_columns` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `cabin_id` int,
  `column_code` varchar(5),
  `created_at` datetime,
  `updated_at` datetime
);

CREATE TABLE `seat_rows` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `cabin_id` int,
  `row_number` int,
  `created_at` datetime,
  `updated_at` datetime
);

CREATE TABLE `seats` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `seat_row_id` int,
  `code` varchar(5),
  `storefront_slot_code` varchar(20),
  `refund_indicator` varchar(5),
  `free_of_charge` boolean,
  `created_at` datetime,
  `updated_at` datetime
);

CREATE TABLE `seat_characteristics` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `seat_id` int,
  `characteristic` varchar(50),
  `created_at` datetime,
  `updated_at` datetime
);

CREATE TABLE `raw_seat_characteristics` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `seat_id` int,
  `raw_characteristic` varchar(50),
  `created_at` datetime,
  `updated_at` datetime
);

CREATE TABLE `passengers` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `passenger_index` int,
  `passenger_name_number` varchar(10),
  `first_name` varchar(100),
  `last_name` varchar(100),
  `date_of_birth` date,
  `gender` varchar(10),
  `type` varchar(10),
  `street1` varchar(255),
  `street2` varchar(255),
  `postcode` varchar(20),
  `state` varchar(100),
  `city` varchar(100),
  `country` varchar(5),
  `address_type` varchar(20),
  `issuing_country` varchar(5),
  `country_of_birth` varchar(5),
  `document_type` varchar(5),
  `nationality` varchar(5),
  `created_at` datetime,
  `updated_at` datetime
);

CREATE TABLE `passenger_emails` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `passenger_id` int,
  `email` varchar(255),
  `created_at` datetime,
  `updated_at` datetime
);

CREATE TABLE `passenger_phones` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `passenger_id` int,
  `phone_number` varchar(50),
  `created_at` datetime,
  `updated_at` datetime
);

CREATE TABLE `frequent_flyers` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `passenger_id` int,
  `airline` varchar(10),
  `number` varchar(50),
  `tier_number` int,
  `created_at` datetime,
  `updated_at` datetime
);

CREATE TABLE `special_preferences` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `passenger_id` int,
  `meal_preference` varchar(100),
  `seat_preference` varchar(100),
  `created_at` datetime,
  `updated_at` datetime
);

CREATE TABLE `special_requests` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `preference_id` int,
  `request` varchar(255),
  `created_at` datetime,
  `updated_at` datetime
);

CREATE TABLE `special_service_request_remarks` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `preference_id` int,
  `remark` varchar(255),
  `created_at` datetime,
  `updated_at` datetime
);

CREATE TABLE `bookings` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `booking_reference` varchar(50),
  `origin` varchar(5),
  `destination` varchar(5),
  `departure` datetime,
  `arrival` datetime,
  `equipment` varchar(10),
  `fare_basis` varchar(50),
  `booking_class` varchar(5),
  `cabin_class` varchar(20),
  `duration` int,
  `layover_duration` int,
  `segment_ref` varchar(50),
  `subject_to_government_approval` boolean,
  `created_at` datetime,
  `updated_at` datetime
);

CREATE TABLE `booking_flights` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `booking_id` int,
  `flight_number` int,
  `operating_flight_number` int,
  `airline_code` varchar(10),
  `operating_airline_code` varchar(10),
  `departure_terminal` varchar(50),
  `arrival_terminal` varchar(50),
  `created_at` datetime,
  `updated_at` datetime
);

CREATE TABLE `booking_flight_stop_airports` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `flight_id` int,
  `airport_code` varchar(10),
  `created_at` datetime,
  `updated_at` datetime
);

CREATE TABLE `booking_seats` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `booking_id` int,
  `passenger_id` int,
  `seat_id` int,
  `available` boolean,
  `entitled` boolean,
  `fee_waived` boolean,
  `entitled_rule_id` varchar(50),
  `fee_waived_rule_id` varchar(50),
  `price_amount` decimal(10,2),
  `price_currency` varchar(5),
  `tax_amount` decimal(10,2),
  `tax_currency` varchar(5),
  `total_amount` decimal(10,2),
  `total_currency` varchar(5),
  `originally_selected` boolean,
  `created_at` datetime,
  `updated_at` datetime
);

ALTER TABLE `cabins` ADD FOREIGN KEY (`aircraft_id`) REFERENCES `aircrafts` (`id`);

ALTER TABLE `seat_columns` ADD FOREIGN KEY (`cabin_id`) REFERENCES `cabins` (`id`);

ALTER TABLE `seat_rows` ADD FOREIGN KEY (`cabin_id`) REFERENCES `cabins` (`id`);

ALTER TABLE `seats` ADD FOREIGN KEY (`seat_row_id`) REFERENCES `seat_rows` (`id`);

ALTER TABLE `seat_characteristics` ADD FOREIGN KEY (`seat_id`) REFERENCES `seats` (`id`);

ALTER TABLE `raw_seat_characteristics` ADD FOREIGN KEY (`seat_id`) REFERENCES `seats` (`id`);

ALTER TABLE `passenger_emails` ADD FOREIGN KEY (`passenger_id`) REFERENCES `passengers` (`id`);

ALTER TABLE `passenger_phones` ADD FOREIGN KEY (`passenger_id`) REFERENCES `passengers` (`id`);

ALTER TABLE `frequent_flyers` ADD FOREIGN KEY (`passenger_id`) REFERENCES `passengers` (`id`);

ALTER TABLE `special_preferences` ADD FOREIGN KEY (`passenger_id`) REFERENCES `passengers` (`id`);

ALTER TABLE `special_requests` ADD FOREIGN KEY (`preference_id`) REFERENCES `special_preferences` (`id`);

ALTER TABLE `special_service_request_remarks` ADD FOREIGN KEY (`preference_id`) REFERENCES `special_preferences` (`id`);

ALTER TABLE `booking_flights` ADD FOREIGN KEY (`booking_id`) REFERENCES `bookings` (`id`);

ALTER TABLE `booking_flight_stop_airports` ADD FOREIGN KEY (`flight_id`) REFERENCES `booking_flights` (`id`);

ALTER TABLE `booking_seats` ADD FOREIGN KEY (`booking_id`) REFERENCES `bookings` (`id`);

ALTER TABLE `booking_seats` ADD FOREIGN KEY (`passenger_id`) REFERENCES `passengers` (`id`);

ALTER TABLE `booking_seats` ADD FOREIGN KEY (`seat_id`) REFERENCES `seats` (`id`);
